// Package ebiten contains ebiten's implementation of the audio interface
package ebiten

import (
	"io"

	"github.com/BismarckDD/Caocao/c2common/c2interface"
	"github.com/BismarckDD/Caocao/c2common/c2util"
	"github.com/BismarckDD/Caocao/c2core/c2asset"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

const DefaultSampleRate = 44100 // 采样率 44.1K Hz

const logPrefix = "Ebiten Audio Provider"

var _ c2interface.AudioProvider = &AudioProvider{} // Static check to confirm struct conforms to interface

// CreateAudio creates an instance of ebiten's audio provider
func CreateAudio(l c2util.LogLevel, am *c2asset.AssetManager) *AudioProvider {

	audioProvider := &AudioProvider{
		asset: am,
	}

	audioProvider.Logger = c2util.NewLogger()
	audioProvider.Logger.SetLevel(l)
	audioProvider.Logger.SetPrefix(logPrefix)

	audioProvider.audioContext = audio.NewContext(DefaultSampleRate)

	return audioProvider
}

// AudioProvider represents a provider capable of playing audio
type AudioProvider struct {
	asset        *c2asset.AssetManager
	audioContext *audio.Context // The Audio context
	bgmAudio     *audio.Player  // The Audio player
	bgmStream    *wav.Stream
	lastBgm      string
	sfxVolume    float64
	bgmVolume    float64

	*c2util.Logger
}

// PlayBGM loads an audio stream and plays it in the background
func (eap *AudioProvider) PlayBGM(song string) {
	if eap.lastBgm == song {
		return
	}

	eap.lastBgm = song

	if song == "" && eap.bgmAudio != nil && eap.bgmAudio.IsPlaying() {
		eap.bgmAudio.Pause()
		return
	}

	if eap.bgmAudio != nil {
		err := eap.bgmAudio.Close()

		if err != nil {
			eap.Fatal(err.Error())
		}
	}

	audioStream, err := eap.asset.LoadFileStream(song)

	if err != nil {
		panic(err)
	}

	if _, err = audioStream.Seek(0, io.SeekStart); err != nil {
		eap.Fatal(err.Error())
	}

	eap.bgmStream, err = wav.Decode(eap.audioContext, audioStream)

	if err != nil {
		eap.Fatal(err.Error())
	}

	s := audio.NewInfiniteLoop(eap.bgmStream, eap.bgmStream.Length())
	eap.bgmAudio, err = audio.NewPlayer(eap.audioContext, s)

	if err != nil {
		eap.Fatal(err.Error())
	}

	eap.bgmAudio.SetVolume(eap.bgmVolume)

	// Play the infinite-length stream. This never ends.
	err = eap.bgmAudio.Rewind()

	if err != nil {
		panic(err)
	}

	eap.bgmAudio.Play()
}

// LoadSound loads a sound affect so that it canb e played
func (eap *AudioProvider) LoadSound(sfx string, loop, bgm bool) (c2interface.SoundEffect, error) {
	volume := eap.sfxVolume
	if bgm {
		volume = eap.bgmVolume
	}

	result := eap.createSoundEffect(sfx, eap.audioContext, loop)

	result.volumeScale = volume
	result.SetVolume(volume)

	return result, nil
}

// SetVolumes sets the volumes of the audio provider
func (eap *AudioProvider) SetVolumes(bgmVolume, sfxVolume float64) {
	eap.sfxVolume = sfxVolume
	eap.bgmVolume = bgmVolume
}

// createSoundEffect creates a new instance of ebiten's sound effect implementation.
func (eap *AudioProvider) createSoundEffect(sfx string, context *audio.Context,
	loop bool) *SoundEffect {
	result := &SoundEffect{}

	soundFile := "data/global/sfx/"

	if _, exists := eap.asset.Records.Sound.Details[sfx]; exists {
		soundEntry := eap.asset.Records.Sound.Details[sfx]
		soundFile += soundEntry.FileName
	} else {
		soundFile += sfx
	}

	if fileExists, _ := eap.asset.FileExists(soundFile); !fileExists {
		soundFile = "data/global/music/" + sfx
	}

	audioData, err := eap.asset.LoadFileStream(soundFile)

	if err != nil {
		audioData, err = eap.asset.LoadFileStream("data/global/music/" + sfx)
	}

	if err != nil {
		panic(err)
	}

	d, err := wav.Decode(context, audioData)

	if err != nil {
		eap.Fatal(err.Error())
	}

	if d == nil {
		eap.Fatal("Decoded data is nil")
	}

	var player *audio.Player

	if loop {
		s := audio.NewInfiniteLoop(d, d.Length())
		result.panStream = newPanStreamFromReader(s)
		player, err = audio.NewPlayer(context, result.panStream)
	} else {
		result.panStream = newPanStreamFromReader(d)
		player, err = audio.NewPlayer(context, result.panStream)
	}

	if err != nil {
		eap.Fatal(err.Error())
	}

	result.player = player

	return result
}

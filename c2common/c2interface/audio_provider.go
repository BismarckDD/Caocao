package c2interface

type SoundEffect interface {
	Play()
	Stop()
	SetPan(pan float64)
	IsPlaying() bool
	SetVolume(volume float64)
}

// AudioProvider is something that can play music, load audio files managed
// by the asset manager, and set the game engine's volume levels
type AudioProvider interface {
	PlayBGM(song string)
	LoadSound(sfx string, loop bool, bgm bool) (SoundEffect, error)
	SetVolumes(bgmVolume, sfxVolume float64)
}

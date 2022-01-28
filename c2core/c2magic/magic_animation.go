package c2magic

import (
	"bufio"
	"fmt"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"

	"github.com/BismarckDD/Caocao/c2common/c2util"
	"github.com/BismarckDD/Caocao/c2core/c2uires"
	"github.com/hajimehoshi/ebiten/v2"
)

const magicAnimationSrc string = "resources/high_res/magic/animation/"

var DefaultCallAnimation []*MagicAnimation = make([]*MagicAnimation, 256)
var DefaultEffAnimation []*MagicAnimation = make([]*MagicAnimation, 256)

/*
Mcall: only few magics have call animation.
Meff: all magics have eff animation.
*/
var callFileTemplate string = "Mcall%2d_%d.png"
var effFileTemplate string = "Meff%2d_%d.png"

/*
MagicAnimation belongs to some specific magic, e.g. Fire/Water/Wind/Land-Dragon
Not all magics has this animation.
*/
type MagicAnimation struct {
	MagicAnimationId uint8           // decide which source file
	FrameNum         uint8           // how many frames in this magic
	DurationInSec    uint8           // durtion of this animation
	Images           []*ebiten.Image // all images in this animation
}

/*
MagicDescriptionFile Description.

*/
func LoadMagicAnimation(magicAnimationMetaFilePath string) error {

	magicMetaFile, err := os.OpenFile(c2util.GetAppPath()+"/"+magicAnimationMetaFilePath, os.O_RDONLY, 6)
	if err != nil {
		return nil
	}
	rd := bufio.NewReader(magicMetaFile)

	magicDigit, err := rd.Read()
	callAnimationNum, err := rd.ReadByte()

	for {

		magicAnimationId, err := rd.ReadByte()
		if err != nil {
			break
		}

		frameNum, err := rd.ReadByte()
		if err != nil {
			log.Fatal(err.Error())
		}

		durationInSec, err := rd.ReadByte()
		if err != nil {
			log.Fatal(err.Error())
		}
		images := []*ebiten.Image{}
		for i := 0; i < int(frameNum); i++ {
			filename := fmt.Sprintf(callFileTemplate, magicAnimationId, i)
			file, err := os.Create(c2util.GetAppPath() + "/" + magicAnimationSrc + filename)
			if err != nil {
				log.Println()
				continue
			}
			img := c2uires.NewImageFromFile(file)
			images = append(images, img)
		}
		pAnimation := &MagicAnimation{
			MagicAnimationId: magicAnimationId,
			FrameNum:         frameNum,
			DurationInSec:    durationInSec,
			Images:           images,
		}
		DefaultCallAnimation[magicAnimationId] = pAnimation
	}
	return nil
}

func SaveMagicAnimation(magicAnimationMetaFilePath string) error {

	magicMetaFile, err := os.OpenFile(c2util.GetAppPath()+"/"+magicAnimationMetaFilePath, os.O_WRONLY, 0777)
	if err != nil {
		return nil
	}
	wr := bufio.NewWriter(magicMetaFile)
	magicDigit := 20220129
	wr.Write(make([]byte, magicDigit))
	return nil
}

func (animation *MagicAnimation) Play() {
	// TODO: play the animation on target area on the screen.
	return
}

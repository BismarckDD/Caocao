package c2uires

import (
	go_image "image"
	"image/color"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/BismarckDD/Caocao/c2common/c2util"
	"github.com/blizzy78/ebitenui/image"
	"github.com/blizzy78/ebitenui/widget"
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/font"
)

func LoadImage(path string) (*ebiten.Image, go_image.Image, error) {
	return ebitenutil.NewImageFromFile(c2util.GetAppPath() + "/" + path)
}

func loadImageFromFile(path string) (*ebiten.Image, go_image.Image, error) {
	return ebitenutil.NewImageFromFile(c2util.GetAppPath() + "/" + path)
}

func LoadFont(path string) (*truetype.Font, error) {
	fontData, err := ioutil.ReadFile(c2util.GetAppPath() + "/" + path)
	if err != nil {
		return nil, err
	}

	ttfFont, err := truetype.Parse(fontData)
	if err != nil {
		return nil, err
	}
	return ttfFont, nil

	// return truetype.NewFace(ttfFont, &truetype.Options{
	// 	Size:    size,
	// 	DPI:     72,
	// 	Hinting: font.HintingFull,
	// }), nil
}

func GetFontFace(truetypeFont *truetype.Font, size float64) (font.Face, error) {
	return truetype.NewFace(truetypeFont, &truetype.Options{
		Size:    size,
		DPI:     72,
		Hinting: font.HintingFull,
	}), nil
}

func loadGraphicImages(idle string, disabled string) (*widget.ButtonImageImage, error) {
	idleImage, _, err := ebitenutil.NewImageFromFile(c2util.GetAppPath() + "/" + idle)
	if err != nil {
		return nil, err
	}

	var disabledImage *ebiten.Image
	if disabled != "" {
		disabledImage, _, err = ebitenutil.NewImageFromFile(c2util.GetAppPath() + "/" + disabled)
		if err != nil {
			return nil, err
		}
	}

	return &widget.ButtonImageImage{
		Idle:     idleImage,
		Disabled: disabledImage,
	}, nil
}

func LoadImageNineSlice(path string, centerWidth int, centerHeight int) (*image.NineSlice, error) {
	return loadImageNineSlice(path, centerWidth, centerHeight)
}

func loadImageNineSlice(path string, centerWidth int, centerHeight int) (*image.NineSlice, error) {
	rawImage, _, err := ebitenutil.NewImageFromFile(c2util.GetAppPath() + "/" + path)
	if err != nil {
		return nil, err
	}

	width, height := rawImage.Size()
	return image.NewNineSlice(rawImage,
			[3]int{(width - centerWidth) / 2, centerWidth, width - (width-centerWidth)/2 - centerWidth},
			[3]int{(height - centerHeight) / 2, centerHeight, height - (height-centerHeight)/2 - centerHeight}),
		nil
}

func hexToColor(hexString string) color.Color {
	u, err := strconv.ParseUint(hexString, 16, 0)
	if err != nil {
		log.Fatalln("Failed to get hex color")
		panic(err)
	}
	return color.RGBA{
		R: uint8(u & 0xff0000 >> 16),
		G: uint8(u & 0xff00 >> 8),
		B: uint8(u & 0xff),
		A: 255,
	}
}

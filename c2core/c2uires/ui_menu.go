package c2uires

import (
	"image/color"

	"github.com/blizzy78/ebitenui/image"
	"github.com/blizzy78/ebitenui/widget"
	"golang.org/x/image/font"
)

type MenuUI struct {
	Background *image.NineSlice
	Padding    widget.Insets
	Face       font.Face // contain font-style && size
	Color      color.Color
}

func CreateWindowsStyleMenuUI(uiRes *UIResources) (*MenuUI, error) {

	// window-style font size is 12.
	face, err := GetFontFace(uiRes.FontMap["lishu"], 12)
	if err != nil {
		return nil, err
	}

	background, err := loadImageNineSlice("graphics/header.png", 446, 9)
	if err != nil {
		return nil, err
	}

	return &MenuUI{
		Background: background,
		Face:       face,
		Color:      hexToColor(buttonIdleBackground),
		Padding: widget.Insets{
			Left:   0,
			Right:  0,
			Top:    0,
			Bottom: 0,
		},
	}, nil
}

func CreateGameStyleMenuUI(uiRes *UIResources) (*MenuUI, error) {

	// window-style font size is 12.
	face, _ := GetFontFace(uiRes.FontMap["lishu"], 12)

	background, err := loadImageNineSlice("graphics/header.png", 446, 9)
	if err != nil {
		return nil, err
	}

	return &MenuUI{
		Background: background,
		Face:       face,
		Color:      hexToColor(buttonIdleBackground),
		Padding: widget.Insets{
			Left:   10,
			Right:  10,
			Top:    2,
			Bottom: 2,
		},
	}, nil
}

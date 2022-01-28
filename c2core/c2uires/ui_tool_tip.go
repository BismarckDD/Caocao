package c2uires

import (
	"image/color"

	"github.com/blizzy78/ebitenui/image"
	"github.com/blizzy78/ebitenui/widget"
	"golang.org/x/image/font"
)

type ToolTipUI struct {
	Background *image.NineSlice
	Padding    widget.Insets
	Face       font.Face
	Color      color.Color
}

func NewToolTipResources(faceFont font.Face) (*ToolTipUI, error) {

	bg, _, err := loadImageFromFile("graphics/tool-tip.png")
	if err != nil {
		return nil, err
	}

	return &ToolTipUI{

		Background: image.NewNineSlice(bg, [3]int{19, 6, 13}, [3]int{19, 5, 13}),
		Padding: widget.Insets{
			Left:   15,
			Right:  15,
			Top:    10,
			Bottom: 10,
		},
		Face:  faceFont,
		Color: hexToColor(toolTipColor),
	}, nil
}

func NewToolTipCharAtBattleField(faceFont font.Face) (*ToolTipUI, error) {

	bg, _, err := loadImageFromFile("graphics/tool-tip.png")
	if err != nil {
		return nil, err
	}

	return &ToolTipUI{

		Background: image.NewNineSlice(bg, [3]int{19, 6, 13}, [3]int{19, 5, 13}),
		Padding: widget.Insets{
			Left:   15,
			Right:  15,
			Top:    10,
			Bottom: 10,
		},
		Face:  faceFont,
		Color: hexToColor(toolTipColor),
	}, nil
}

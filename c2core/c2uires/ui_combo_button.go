package c2uires

import (
	"github.com/blizzy78/ebitenui/widget"
	"golang.org/x/image/font"
)

type ComboButtonUI struct {
	Image     *widget.ButtonImage
	TextColor *widget.ButtonTextColor
	Face      font.Face
	Graphic   *widget.ButtonImageImage
	Padding   widget.Insets
}

func NewComboButtonResources(faceFont font.Face) (*ComboButtonUI, error) {

	idle, err := loadImageNineSlice("graphics/combo-button-idle.png", 12, 0)
	if err != nil {
		return nil, err
	}

	hover, err := loadImageNineSlice("graphics/combo-button-hover.png", 12, 0)
	if err != nil {
		return nil, err
	}

	pressed, err := loadImageNineSlice("graphics/combo-button-pressed.png", 12, 0)
	if err != nil {
		return nil, err
	}

	disabled, err := loadImageNineSlice("graphics/combo-button-disabled.png", 12, 0)
	if err != nil {
		return nil, err
	}

	buttonImage := &widget.ButtonImage{
		Idle:     idle,
		Hover:    hover,
		Pressed:  pressed,
		Disabled: disabled,
	}

	arrowDown, err := loadGraphicImages("graphics/arrow-down-idle.png", "graphics/arrow-down-disabled.png")
	if err != nil {
		return nil, err
	}

	return &ComboButtonUI{
		Image: buttonImage,
		TextColor: &widget.ButtonTextColor{
			Idle:     hexToColor(buttonIdleColor),
			Disabled: hexToColor(buttonDisabledColor),
		},
		Face:    faceFont,
		Graphic: arrowDown,
		Padding: widget.Insets{
			Left:  30,
			Right: 30,
		},
	}, nil
}

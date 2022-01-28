package c2uires

import (
	"github.com/blizzy78/ebitenui/widget"
	"golang.org/x/image/font"
)

type TabBookUI struct {
	IdleButton     *widget.ButtonImage
	SelectedButton *widget.ButtonImage
	Face           font.Face
	TextColor      *widget.ButtonTextColor
	Padding        widget.Insets
}

func NewTabBookResources(faceFont font.Face) (*TabBookUI, error) {

	selectedIdle, err := loadImageNineSlice("graphics/button-selected-idle.png", 12, 0)
	if err != nil {
		return nil, err
	}

	selectedHover, err := loadImageNineSlice("graphics/button-selected-hover.png", 12, 0)
	if err != nil {
		return nil, err
	}

	selectedPressed, err := loadImageNineSlice("graphics/button-selected-pressed.png", 12, 0)
	if err != nil {
		return nil, err
	}

	selectedDisabled, err := loadImageNineSlice("graphics/button-selected-disabled.png", 12, 0)
	if err != nil {
		return nil, err
	}

	selected := &widget.ButtonImage{
		Idle:     selectedIdle,
		Hover:    selectedHover,
		Pressed:  selectedPressed,
		Disabled: selectedDisabled,
	}

	idle, err := loadImageNineSlice("graphics/button-idle.png", 12, 0)
	if err != nil {
		return nil, err
	}

	hover, err := loadImageNineSlice("graphics/button-hover.png", 12, 0)
	if err != nil {
		return nil, err
	}

	pressed, err := loadImageNineSlice("graphics/button-pressed.png", 12, 0)
	if err != nil {
		return nil, err
	}

	disabled, err := loadImageNineSlice("graphics/button-disabled.png", 12, 0)
	if err != nil {
		return nil, err
	}

	unselected := &widget.ButtonImage{
		Idle:     idle,
		Hover:    hover,
		Pressed:  pressed,
		Disabled: disabled,
	}

	return &TabBookUI{
		SelectedButton: selected,
		IdleButton:     unselected,
		Face:           faceFont,
		TextColor: &widget.ButtonTextColor{
			Idle:     hexToColor(buttonIdleColor),
			Disabled: hexToColor(buttonDisabledColor),
		},
		Padding: widget.Insets{
			Left:  30,
			Right: 30,
		},
	}, nil
}

package c2uires

import "github.com/blizzy78/ebitenui/widget"

type CheckboxUI struct {
	image   *widget.ButtonImage
	graphic *widget.CheckboxGraphicImage
	spacing int
}

func NewCheckboxResources() (*CheckboxUI, error) {

	idle, err := loadImageNineSlice("graphics/checkbox-idle.png", 20, 0)
	if err != nil {
		return nil, err
	}

	hover, err := loadImageNineSlice("graphics/checkbox-hover.png", 20, 0)
	if err != nil {
		return nil, err
	}

	disabled, err := loadImageNineSlice("graphics/checkbox-disabled.png", 20, 0)
	if err != nil {
		return nil, err
	}

	checked, err := loadGraphicImages("graphics/checkbox-checked-idle.png", "graphics/checkbox-checked-disabled.png")
	if err != nil {
		return nil, err
	}

	unchecked, err := loadGraphicImages("graphics/checkbox-unchecked-idle.png", "graphics/checkbox-unchecked-disabled.png")
	if err != nil {
		return nil, err
	}

	greyed, err := loadGraphicImages("graphics/checkbox-greyed-idle.png", "graphics/checkbox-greyed-disabled.png")
	if err != nil {
		return nil, err
	}

	return &CheckboxUI{
		image: &widget.ButtonImage{
			Idle:     idle,
			Hover:    hover,
			Pressed:  hover,
			Disabled: disabled,
		},

		graphic: &widget.CheckboxGraphicImage{
			Checked:   checked,
			Unchecked: unchecked,
			Greyed:    greyed,
		},

		spacing: 10,
	}, nil
}

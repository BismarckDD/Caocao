package c2uires

import (
	"github.com/blizzy78/ebitenui/image"
	"github.com/blizzy78/ebitenui/widget"
	"golang.org/x/image/font"
)

type ButtonUI struct {
	Image     *widget.ButtonImage
	TextColor *widget.ButtonTextColor
	Face      font.Face
	Padding   widget.Insets
}

func NewButtonResources(uires *UIResources) (*ButtonUI, error) {

	idle, err := loadImageNineSlice("graphics/button-idle.png", 62, 41)
	if err != nil {
		return nil, err
	}

	hover, err := loadImageNineSlice("graphics/button-hover.png", 62, 41)
	if err != nil {
		return nil, err
	}

	pressed, err := loadImageNineSlice("graphics/button-pressed.png", 62, 41)
	if err != nil {
		return nil, err
	}

	disabled, err := loadImageNineSlice("graphics/button-disabled.png", 62, 41)
	if err != nil {
		return nil, err
	}

	buttonImage := &widget.ButtonImage{
		Idle:     idle,
		Hover:    hover,
		Pressed:  pressed,
		Disabled: disabled,
	}

	face, _ := GetFontFace(uires.FontMap["lishu"], 12)

	return &ButtonUI{
		Image: buttonImage,
		TextColor: &widget.ButtonTextColor{
			Idle:     hexToColor(buttonIdleColor),
			Disabled: hexToColor(buttonDisabledColor),
		},
		Face: face,
		Padding: widget.Insets{
			Left:   16,
			Right:  0,
			Top:    16,
			Bottom: 0,
		},
	}, nil
}

func CreateMainMenuButton(res *UIResources) (*ButtonUI, error) {

	idle := image.NewNineSliceColor(hexToColor(buttonIdleBackground))
	hover := image.NewNineSliceColor(hexToColor(buttonHoverBackground))
	pressed := image.NewNineSliceColor(hexToColor(buttonPressedBackground))
	disabled := image.NewNineSliceColor(hexToColor(buttonDisabledBackground))

	buttonImage := &widget.ButtonImage{
		Idle:     idle,
		Hover:    hover,
		Pressed:  pressed,
		Disabled: disabled,
	}

	face, _ := GetFontFace(res.FontMap["kaishu"], 24)

	return &ButtonUI{
		Image: buttonImage,
		TextColor: &widget.ButtonTextColor{
			Idle:     hexToColor(buttonIdleColor),
			Disabled: hexToColor(buttonDisabledColor),
		},
		Face: face,
		Padding: widget.Insets{
			Left:   128,
			Right:  128,
			Top:    30,
			Bottom: 30,
		}, // This is used for text padding of the button.
	}, nil
}

func CreateGameStyleButton(res *UIResources) (*ButtonUI, error) {

	idle, err := loadImageNineSlice("graphics/button-idle.png", 62, 41)
	if err != nil {
		return nil, err
	}

	hover, err := loadImageNineSlice("graphics/button-hover.png", 62, 41)
	if err != nil {
		return nil, err
	}

	pressed, err := loadImageNineSlice("graphics/button-pressed.png", 62, 41)
	if err != nil {
		return nil, err
	}

	disabled, err := loadImageNineSlice("graphics/button-disabled.png", 62, 41)
	if err != nil {
		return nil, err
	}

	buttonImage := &widget.ButtonImage{
		Idle:     idle,
		Hover:    hover,
		Pressed:  pressed,
		Disabled: disabled,
	}

	face, _ := GetFontFace(res.FontMap["kaishu"], 16)

	return &ButtonUI{
		Image: buttonImage,
		TextColor: &widget.ButtonTextColor{
			Idle:     hexToColor(buttonIdleColor),
			Disabled: hexToColor(buttonDisabledColor),
		},
		Face: face,
		Padding: widget.Insets{
			Left:   16,
			Right:  0,
			Top:    16,
			Bottom: 0,
		},
	}, nil
}

func CreateWindowsStyleButton(res *UIResources) (*ButtonUI, error) {

	idle, err := loadImageNineSlice("graphics/button-idle.png", 62, 41)
	if err != nil {
		return nil, err
	}

	hover, err := loadImageNineSlice("graphics/button-hover.png", 62, 41)
	if err != nil {
		return nil, err
	}

	pressed, err := loadImageNineSlice("graphics/button-pressed.png", 62, 41)
	if err != nil {
		return nil, err
	}

	disabled, err := loadImageNineSlice("graphics/button-disabled.png", 62, 41)
	if err != nil {
		return nil, err
	}

	buttonImage := &widget.ButtonImage{
		Idle:     idle,
		Hover:    hover,
		Pressed:  pressed,
		Disabled: disabled,
	}

	face, _ := GetFontFace(res.FontMap["kaishu"], 16)

	return &ButtonUI{
		Image: buttonImage,
		TextColor: &widget.ButtonTextColor{
			Idle:     hexToColor(buttonIdleColor),
			Disabled: hexToColor(buttonDisabledColor),
		},
		Face: face,
		Padding: widget.Insets{
			Left:   16,
			Right:  0,
			Top:    16,
			Bottom: 0,
		},
	}, nil
}

package c2uires

import (
	"github.com/blizzy78/ebitenui/image"
	"github.com/blizzy78/ebitenui/widget"
	"golang.org/x/image/font"
)

type labelResources struct {
	text *widget.LabelColor
	face font.Face
}

type sliderResources struct {
	trackImage *widget.SliderTrackImage
	handle     *widget.ButtonImage
	handleSize int
}

type panelResources struct {
	image   *image.NineSlice
	padding widget.Insets
}

type textInputResources struct {
	image   *widget.TextInputImage
	padding widget.Insets
	face    font.Face
	color   *widget.TextInputColor
}

func NewLabelResources(faceFont font.Face) *labelResources {
	return &labelResources{
		text: &widget.LabelColor{
			Idle:     hexToColor(labelIdleColor),
			Disabled: hexToColor(labelDisabledColor),
		},

		face: faceFont,
	}
}

func NewSliderResources() (*sliderResources, error) {
	idle, _, err := loadImageFromFile("graphics/slider-track-idle.png")
	if err != nil {
		return nil, err
	}

	disabled, _, err := loadImageFromFile("graphics/slider-track-disabled.png")
	if err != nil {
		return nil, err
	}

	handleIdle, _, err := loadImageFromFile("graphics/slider-handle-idle.png")
	if err != nil {
		return nil, err
	}

	handleHover, _, err := loadImageFromFile("graphics/slider-handle-hover.png")
	if err != nil {
		return nil, err
	}

	handleDisabled, _, err := loadImageFromFile("graphics/slider-handle-disabled.png")
	if err != nil {
		return nil, err
	}

	return &sliderResources{
		trackImage: &widget.SliderTrackImage{
			Idle:     image.NewNineSlice(idle, [3]int{0, 19, 0}, [3]int{6, 0, 0}),
			Hover:    image.NewNineSlice(idle, [3]int{0, 19, 0}, [3]int{6, 0, 0}),
			Disabled: image.NewNineSlice(disabled, [3]int{0, 19, 0}, [3]int{6, 0, 0}),
		},

		handle: &widget.ButtonImage{
			Idle:     image.NewNineSliceSimple(handleIdle, 0, 5),
			Hover:    image.NewNineSliceSimple(handleHover, 0, 5),
			Pressed:  image.NewNineSliceSimple(handleHover, 0, 5),
			Disabled: image.NewNineSliceSimple(handleDisabled, 0, 5),
		},

		handleSize: 6,
	}, nil
}

func NewPanelResources() (*panelResources, error) {
	image, err := loadImageNineSlice("graphics/panel-idle.png", 10, 10)
	if err != nil {
		return nil, err
	}

	return &panelResources{
		image: image,
		padding: widget.Insets{
			Left:   20,
			Right:  20,
			Top:    10,
			Bottom: 10,
		},
	}, nil
}

func NewTextInputResources(faceFont font.Face) (*textInputResources, error) {
	idle, _, err := loadImageFromFile("graphics/text-input-idle.png")
	if err != nil {
		return nil, err
	}

	disabled, _, err := loadImageFromFile("graphics/text-input-disabled.png")
	if err != nil {
		return nil, err
	}

	return &textInputResources{
		image: &widget.TextInputImage{
			Idle:     image.NewNineSlice(idle, [3]int{9, 14, 6}, [3]int{9, 14, 6}),
			Disabled: image.NewNineSlice(disabled, [3]int{9, 14, 6}, [3]int{9, 14, 6}),
		},

		padding: widget.Insets{
			Left:   8,
			Right:  8,
			Top:    4,
			Bottom: 4,
		},

		face: faceFont,

		color: &widget.TextInputColor{
			Idle:          hexToColor(textIdleColor),
			Disabled:      hexToColor(textDisabledColor),
			Caret:         hexToColor(textInputCaretColor),
			DisabledCaret: hexToColor(textInputDisabledCaretColor),
		},
	}, nil
}

package c2uires

import (
	"github.com/blizzy78/ebitenui/image"
	"github.com/blizzy78/ebitenui/widget"
	"golang.org/x/image/font"
)

type ListUI struct {
	Image        *widget.ScrollContainerImage
	Track        *widget.SliderTrackImage
	TrackPadding widget.Insets
	Handle       *widget.ButtonImage
	HandleSize   int
	Face         font.Face
	Entry        *widget.ListEntryColor
	EntryPadding widget.Insets
}

func NewListResources(faceFont font.Face) (*ListUI, error) {
	idle, _, err := loadImageFromFile("graphics/list-idle.png")
	if err != nil {
		return nil, err
	}

	disabled, _, err := loadImageFromFile("graphics/list-disabled.png")
	if err != nil {
		return nil, err
	}

	mask, _, err := loadImageFromFile("graphics/list-mask.png")
	if err != nil {
		return nil, err
	}

	trackIdle, _, err := loadImageFromFile("graphics/list-track-idle.png")
	if err != nil {
		return nil, err
	}

	trackDisabled, _, err := loadImageFromFile("graphics/list-track-disabled.png")
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

	return &ListUI{
		Image: &widget.ScrollContainerImage{
			Idle:     image.NewNineSlice(idle, [3]int{25, 12, 22}, [3]int{25, 12, 25}),
			Disabled: image.NewNineSlice(disabled, [3]int{25, 12, 22}, [3]int{25, 12, 25}),
			Mask:     image.NewNineSlice(mask, [3]int{26, 10, 23}, [3]int{26, 10, 26}),
		},

		Track: &widget.SliderTrackImage{
			Idle:     image.NewNineSlice(trackIdle, [3]int{5, 0, 0}, [3]int{25, 12, 25}),
			Hover:    image.NewNineSlice(trackIdle, [3]int{5, 0, 0}, [3]int{25, 12, 25}),
			Disabled: image.NewNineSlice(trackDisabled, [3]int{0, 5, 0}, [3]int{25, 12, 25}),
		},

		TrackPadding: widget.Insets{
			Top:    5,
			Bottom: 24,
		},

		Handle: &widget.ButtonImage{
			Idle:     image.NewNineSliceSimple(handleIdle, 0, 5),
			Hover:    image.NewNineSliceSimple(handleHover, 0, 5),
			Pressed:  image.NewNineSliceSimple(handleHover, 0, 5),
			Disabled: image.NewNineSliceSimple(handleIdle, 0, 5),
		},

		HandleSize: 5,

		Face: faceFont,

		Entry: &widget.ListEntryColor{
			Unselected:         hexToColor(textIdleColor),
			DisabledUnselected: hexToColor(textDisabledColor),

			Selected:         hexToColor(textIdleColor),
			DisabledSelected: hexToColor(textDisabledColor),

			SelectedBackground:         hexToColor(listSelectedBackground),
			DisabledSelectedBackground: hexToColor(listDisabledSelectedBackground),
		},

		EntryPadding: widget.Insets{
			Left:   30,
			Right:  30,
			Top:    2,
			Bottom: 2,
		},
	}, nil
}

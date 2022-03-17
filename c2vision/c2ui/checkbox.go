package c2ui

import (
	"github.com/BismarckDD/Caocao/c2common/c2enum"
	"github.com/BismarckDD/Caocao/c2common/c2interface"
	"github.com/BismarckDD/Caocao/c2common/c2resource"
)

// static check that Checkbox implements ClickableWidget interface
var _ ClickableWidget = &Checkbox{}

// Checkbox represents a checkbox UI element
type Checkbox struct {
	*BaseWidget
	Image        c2interface.Surface
	checkedImage c2interface.Surface
	onClick      func()
	checkState   bool
	enabled      bool
}

// NewCheckbox creates a new instance of a checkbox
func (ui *UIManager) NewCheckbox(checkState bool) *Checkbox {
	var err error

	base := NewBaseWidget(ui)

	result := &Checkbox{
		BaseWidget: base,
		checkState: checkState,
		enabled:    true,
	}

	checkboxSprite, err := ui.NewSprite(c2resource.Checkbox, c2resource.PaletteFechar)
	if err != nil {
		ui.Error(err.Error())
		return nil
	}

	result.width, result.height, err = checkboxSprite.GetFrameSize(0)
	if err != nil {
		ui.Error(err.Error())
		return nil
	}

	checkboxSprite.SetPosition(0, 0)

	result.Image = ui.renderer.NewSurface(result.width, result.height)

	checkboxSprite.RenderSegmented(result.Image, 1, 1, 0)

	result.checkedImage = ui.renderer.NewSurface(result.width, result.height)

	checkboxSprite.RenderSegmented(result.checkedImage, 1, 1, 1)

	ui.addWidget(result)

	return result
}

// Render renders the checkbox
func (v *Checkbox) Render(target c2interface.Surface) {
	target.PushTranslation(v.x, v.y)
	defer target.Pop()

	target.PushFilter(c2enum.FilterNearest)
	defer target.Pop()

	if v.checkState {
		target.Render(v.checkedImage)
	} else {
		target.Render(v.Image)
	}
}

// Advance does nothing for checkboxes
func (v *Checkbox) Advance(_ float64) error {
	return nil
}

// GetEnabled returns the enabled state of the checkbox
func (v *Checkbox) GetEnabled() bool {
	return v.enabled
}

// SetEnabled sets the enabled state of the checkbox
func (v *Checkbox) SetEnabled(enabled bool) {
	v.enabled = enabled
}

// SetPressed does nothing for checkboxes
func (v *Checkbox) SetPressed(_ bool) {
}

// SetCheckState sets the check state of the checkbox
func (v *Checkbox) SetCheckState(checkState bool) {
	v.checkState = checkState
}

// GetCheckState returns the check state of the checkbox
func (v *Checkbox) GetCheckState() bool {
	return v.checkState
}

// GetPressed returns the pressed state of the checkbox
func (v *Checkbox) GetPressed() bool {
	return v.checkState
}

// OnActivated sets the callback function of the click event for the checkbox
func (v *Checkbox) OnActivated(callback func()) {
	v.onClick = callback
}

// Activate activates the checkbox
func (v *Checkbox) Activate() {
	v.checkState = !v.checkState
	if v.onClick == nil {
		return
	}

	v.onClick()
}

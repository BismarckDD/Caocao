package c2ui

import (
	"image"

	"github.com/BismarckDD/Caocao/c2common/c2enum"
	"github.com/BismarckDD/Caocao/c2common/c2interface"
	"github.com/BismarckDD/Caocao/c2common/c2resource"
	"github.com/BismarckDD/Caocao/c2common/c2util"
)

// ButtonType defines the type of button
type ButtonType int

// ButtonType constants
const (
	ButtonTypeWide     ButtonType = 1
	ButtonTypeMedium   ButtonType = 2
	ButtonTypeNarrow   ButtonType = 3
	ButtonTypeCancel   ButtonType = 4
	ButtonTypeTall     ButtonType = 5
	ButtonTypeShort    ButtonType = 6
	ButtonTypeOkCancel ButtonType = 7
	// Game UI
	ButtonTypeSquareMenu   ButtonType = 7
	ButtonTypeRectMainMenu ButtonType = 8 // 用于主界面 【开始新的游戏】【读取保存进度】【环境设定】【结束游戏】

	ButtonNoFixedWidth  int = -1
	ButtonNoFixedHeight int = -1
)

const (
	buttonStatePressed = iota + 1
	buttonStateToggled
	buttonStatePressedToggled
)

const (
	exitButtonBaseFrame        = 0  // base frame offset of the "Exit" image button
	saveButtonBaseFrame        = 3  // base frame offset of the "sell" image button
	loadButtonBaseFrame        = 6  // base frame offset of the "repair" button dc6
	infoButtonBaseFrame        = 9  // base frame offset of the "Info" button
	closeButtonBaseFrame       = 12 // base frame offset of the "close" button dc6
	leftArrowButtonBaseFrame   = 15 // base frame offset of the "leftArrow" button dc6
	rightArrowButtonBaseFrame  = 18 // base frame offset of the "rightArrow" button dc6
	okButtonBaseFrame          = 21 // base frame offset of the "ok" button dc6
	repairAllButtonBaseFrame   = 24 // base frame offset of the "repair all" button dc6
	squelchChatButtonBaseFrame = 27 // base frame offset of the "?" button dc6
)

const (
	greyAlpha100     = 0x646464ff
	lightGreyAlpha75 = 0x808080c3
	whiteAlpha100    = 0xffffffff
)

// ButtonLayout defines the type of buttons
type ButtonLayout struct {
	ResourceName     string
	PaletteName      string
	FontPath         string
	ClickableRect    *image.Rectangle
	XSegments        int
	YSegments        int
	BaseFrame        int
	PressedFrame     int
	DisabledFrame    int
	DisabledColor    uint32
	TextOffset       int
	FixedWidth       int
	FixedHeight      int
	LabelColor       uint32
	Toggleable       bool
	AllowFrameChange bool
	HasImage         bool
	Tooltip          int
	TooltipXOffset   int
	TooltipYOffset   int
}

const (
	buttonTooltipNone int = iota
	buttonTooltipClose
	buttonTooltipOk
	buttonTooltipBuy
	buttonTooltipSell
	buttonTooltipRepair
	buttonTooltipRepairAll
	buttonTooltipLeftArrow
	buttonTooltipRightArrow
	buttonTooltipQuery
	buttonTooltipSquelchChat
)

const (
	buttonMenuSegmentsX     = 2
	buttonMenuSegmentsY     = 1
	buttonMenuDisabledFrame = -1
	buttonWideTextOffset    = 1

	buttonShortSegmentsX     = 1
	buttonShortSegmentsY     = 1
	buttonShortDisabledFrame = -1
	buttonShortTextOffset    = -1

	buttonMediumSegmentsX = 1
	buttonMediumSegmentsY = 1

	buttonTallSegmentsX  = 1
	buttonTallSegmentsY  = 1
	buttonTallTextOffset = 5

	buttonCancelSegmentsX  = 1
	buttonCancelSegmentsY  = 1
	buttonCancelTextOffset = 1

	buttonOkCancelSegmentsX     = 1
	buttonOkCancelSegmentsY     = 1
	buttonOkCancelDisabledFrame = -1

	buttonUpDownArrowSegmentsX     = 1
	buttonUpDownArrowSegmentsY     = 1
	buttonUpDownArrowDisabledFrame = -1
	buttonUpArrowBaseFrame         = 0
	buttonDownArrowBaseFrame       = 2

	buttonBuySellSegmentsX     = 1
	buttonBuySellSegmentsY     = 1
	buttonBuySellDisabledFrame = 1

	buttonSkillTreeTabXSegments = 1
	buttonSkillTreeTabYSegments = 1

	buttonSkillTreeTabDisabledFrame = 7
	buttonSkillTreeTabBaseFrame     = 7
	buttonSkillTreeTabFixedWidth    = 93
	buttonSkillTreeTabFixedHeight   = 107

	buttonTabXSegments = 1
	buttonTabYSegments = 1

	buttonMinipanelOpenCloseBaseFrame = 0
	buttonMinipanelXSegments          = 1
	buttonMinipanelYSegments          = 1

	blankQuestButtonXSegments      = 1
	blankQuestButtonYSegments      = 1
	blankQuestButtonDisabledFrames = 0

	buttonMinipanelCharacterBaseFrame = 0
	buttonMinipanelInventoryBaseFrame = 2
	buttonMinipanelSkilltreeBaseFrame = 4
	buttonMinipanelPartyBaseFrame     = 6
	buttonMinipanelAutomapBaseFrame   = 8
	buttonMinipanelMessageBaseFrame   = 10
	buttonMinipanelQuestBaseFrame     = 12
	buttonMinipanelMenBaseFrame       = 14

	buttonRunSegmentsX     = 1
	buttonRunSegmentsY     = 1
	buttonRunDisabledFrame = -1

	buttonGoldCoinSegmentsX     = 1
	buttonGoldCoinSegmentsY     = 1
	buttonGoldCoinDisabledFrame = -1

	buttonAddSkillSegmentsX     = 1
	buttonAddSkillSegmentsY     = 1
	buttonAddSkillDisabledFrame = 2

	partyButtonSegmentsX     = 1
	partyButtonSegmentsY     = 1
	partyButtonDisabledFrame = -1

	pressedButtonOffset = 2
)

// nolint:funlen // cant reduce
func getButtonLayouts() map[ButtonType]*ButtonLayout {

	return map[ButtonType]*ButtonLayout{
		ButtonTypeWide: {
			XSegments:        buttonWideSegmentsX,
			YSegments:        buttonWideSegmentsY,
			DisabledFrame:    buttonWideDisabledFrame,
			DisabledColor:    lightGreyAlpha75,
			TextOffset:       buttonWideTextOffset,
			ResourceName:     c2resource.WideButtonBlank,
			PaletteName:      c2resource.PaletteUnits,
			FontPath:         c2resource.FontExocet10,
			AllowFrameChange: true,
			HasImage:         true,
			FixedWidth:       ButtonNoFixedWidth,
			FixedHeight:      ButtonNoFixedHeight,
			LabelColor:       greyAlpha100,
		},
		ButtonTypeCancel: {
			XSegments:        buttonCancelSegmentsX,
			YSegments:        buttonCancelSegmentsY,
			DisabledFrame:    0,
			DisabledColor:    lightGreyAlpha75,
			TextOffset:       buttonCancelTextOffset,
			ResourceName:     c2resource.CancelButton,
			PaletteName:      c2resource.PaletteUnits,
			FontPath:         c2resource.FontExocet10,
			AllowFrameChange: true,
			HasImage:         true,
			FixedWidth:       ButtonNoFixedWidth,
			FixedHeight:      ButtonNoFixedHeight,
			LabelColor:       greyAlpha100,
		},
		ButtonTypeShort: {
			XSegments:        buttonShortSegmentsX,
			YSegments:        buttonShortSegmentsY,
			DisabledFrame:    buttonShortDisabledFrame,
			DisabledColor:    lightGreyAlpha75,
			TextOffset:       buttonShortTextOffset,
			ResourceName:     c2resource.ShortButtonBlank,
			PaletteName:      c2resource.PaletteUnits,
			FontPath:         c2resource.FontRediculous,
			AllowFrameChange: true,
			HasImage:         true,
			FixedWidth:       ButtonNoFixedWidth,
			FixedHeight:      ButtonNoFixedHeight,
			LabelColor:       greyAlpha100,
		},
		ButtonTypeMedium: {
			XSegments:        buttonMediumSegmentsX,
			YSegments:        buttonMediumSegmentsY,
			DisabledColor:    lightGreyAlpha75,
			ResourceName:     c2resource.MediumButtonBlank,
			PaletteName:      c2resource.PaletteUnits,
			FontPath:         c2resource.FontExocet10,
			AllowFrameChange: true,
			HasImage:         true,
			FixedWidth:       ButtonNoFixedWidth,
			FixedHeight:      ButtonNoFixedHeight,
			LabelColor:       greyAlpha100,
		},
		ButtonTypeTall: {
			XSegments:        buttonTallSegmentsX,
			YSegments:        buttonTallSegmentsY,
			TextOffset:       buttonTallTextOffset,
			DisabledColor:    lightGreyAlpha75,
			ResourceName:     c2resource.TallButtonBlank,
			PaletteName:      c2resource.PaletteUnits,
			FontPath:         c2resource.FontExocet10,
			AllowFrameChange: true,
			HasImage:         true,
			FixedWidth:       ButtonNoFixedWidth,
			FixedHeight:      ButtonNoFixedHeight,
			LabelColor:       greyAlpha100,
		},
		ButtonTypeOkCancel: {
			XSegments:        buttonOkCancelSegmentsX,
			YSegments:        buttonOkCancelSegmentsY,
			DisabledFrame:    buttonOkCancelDisabledFrame,
			DisabledColor:    lightGreyAlpha75,
			ResourceName:     c2resource.CancelButton,
			PaletteName:      c2resource.PaletteUnits,
			FontPath:         c2resource.FontRediculous,
			AllowFrameChange: true,
			HasImage:         true,
			FixedWidth:       ButtonNoFixedWidth,
			FixedHeight:      ButtonNoFixedHeight,
			LabelColor:       greyAlpha100,
		},
	}
}

// static check to ensure button implements clickable widget
var _ ClickableWidget = &Button{}

// Button defines a standard wide UI button
type Button struct {
	*BaseWidget
	buttonLayout          *ButtonLayout
	normalSurface         c2interface.Surface
	pressedSurface        c2interface.Surface
	toggledSurface        c2interface.Surface
	pressedToggledSurface c2interface.Surface
	disabledSurface       c2interface.Surface
	onClick               func()
	enabled               bool
	pressed               bool
	toggled               bool
	tooltip               *Tooltip
}

// NewButton creates an instance of Button
func (ui *UIResource) NewButton(buttonType ButtonType, text string) *Button {
	buttonLayout := getButtonLayouts()[buttonType]

	btn := ui.createButton(buttonLayout, text)

	return btn
}

// NewDefaultButton creates a new button with default settings
func (ui *UIResource) NewDefaultButton(path string, frame int) *Button {
	layout := &ButtonLayout{
		XSegments:        1,
		YSegments:        1,
		DisabledFrame:    frame,
		DisabledColor:    whiteAlpha100,
		ResourceName:     path,
		PaletteName:      c2resource.PaletteSky,
		BaseFrame:        frame,
		Toggleable:       true,
		FontPath:         c2resource.Font16,
		AllowFrameChange: true,
		HasImage:         true,
		FixedWidth:       ButtonNoFixedWidth,
		FixedHeight:      ButtonNoFixedHeight,
	}

	btn := ui.createButton(layout, "")

	return btn
}

// createButton creates button using input layout and text
func (ui *UIResource) createButton(layout *ButtonLayout, text string) *Button {
	base := NewBaseWidget(ui)
	base.SetVisible(true)

	btn := &Button{
		BaseWidget: base,
		enabled:    true,
		pressed:    false,
	}

	btn.buttonLayout = layout

	lbl := ui.NewLabel(layout.FontPath, c2resource.PaletteUnits)
	lbl.SetText(text)
	lbl.Color[0] = c2util.Color(layout.LabelColor)
	lbl.Alignment = HorizontalAlignCenter

	buttonSprite, err := ui.NewSprite(layout.ResourceName, layout.PaletteName)
	if err != nil {
		ui.Error(err.Error())
		return nil
	}

	if layout.FixedWidth > 0 {
		btn.width = layout.FixedWidth
	} else {
		for i := 0; i < layout.XSegments; i++ {
			w, _, frameSizeErr := buttonSprite.GetFrameSize(i)
			if frameSizeErr != nil {
				ui.Error(frameSizeErr.Error())
				return nil
			}

			btn.width += w
		}
	}

	if layout.FixedHeight > 0 {
		btn.height = layout.FixedHeight
	} else {
		for i := 0; i < layout.YSegments; i++ {
			_, h, frameSizeErr := buttonSprite.GetFrameSize(i * layout.YSegments)
			if frameSizeErr != nil {
				ui.Error(frameSizeErr.Error())
				return nil
			}

			btn.height += h
		}
	}

	btn.normalSurface = ui.renderer.NewSurface(btn.width, btn.height)

	buttonSprite.SetPosition(0, 0)
	buttonSprite.SetEffect(c2enum.DrawEffectModulate)

	btn.createTooltip()

	ui.addWidget(btn) // important that this comes before prerenderStates!

	btn.prerenderStates(buttonSprite, layout, lbl)

	return btn
}

type buttonStateDescriptor struct {
	baseFrame            int
	offsetX, offsetY     int
	prerenderdestination *c2interface.Surface
	fmtErr               string
}

func (v *Button) createTooltip() {
	var toolTip *Tooltip

	switch v.buttonLayout.Tooltip {
	case buttonTooltipNone:
		return
	case buttonTooltipClose:
		t = v.manager.NewTooltip(c2resource.Font16, c2resource.PaletteSky, TooltipXCenter, TooltipYBottom)
		t.SetText(v.manager.asset.TranslateString("strClose"))
	case buttonTooltipOk:
		t = v.manager.NewTooltip(c2resource.Font16, c2resource.PaletteSky, TooltipXCenter, TooltipYBottom)
		t.SetText(v.manager.asset.TranslateString(c2enum.OKLabel))
	case buttonTooltipBuy:
		t = v.manager.NewTooltip(c2resource.Font16, c2resource.PaletteSky, TooltipXCenter, TooltipYBottom)
		t.SetText(v.manager.asset.TranslateString("NPCPurchaseItems"))
	case buttonTooltipSell:
		t = v.manager.NewTooltip(c2resource.Font16, c2resource.PaletteSky, TooltipXCenter, TooltipYBottom)
		t.SetText(v.manager.asset.TranslateString("NPCSellItems"))
	case buttonTooltipRepair:
		t = v.manager.NewTooltip(c2resource.Font16, c2resource.PaletteSky, TooltipXCenter, TooltipYBottom)
		t.SetText(v.manager.asset.TranslateString("NPCRepairItems"))
	case buttonTooltipRepairAll:
		t = v.manager.NewTooltip(c2resource.Font16, c2resource.PaletteSky, TooltipXCenter, TooltipYBottom)
		t.SetText(v.manager.asset.TranslateString(c2enum.RepairAll))
	case buttonTooltipLeftArrow:
		t = v.manager.NewTooltip(c2resource.Font16, c2resource.PaletteSky, TooltipXCenter, TooltipYBottom)
		t.SetText(v.manager.asset.TranslateString("KeyLeft"))
	case buttonTooltipRightArrow:
		t = v.manager.NewTooltip(c2resource.Font16, c2resource.PaletteSky, TooltipXCenter, TooltipYBottom)
		t.SetText(v.manager.asset.TranslateString("KeyRight"))
	case buttonTooltipQuery:
		t = v.manager.NewTooltip(c2resource.Font16, c2resource.PaletteSky, TooltipXCenter, TooltipYBottom)
		t.SetText(v.manager.asset.TranslateString("")) // need to be set up
	case buttonTooltipSquelchChat:
		t = v.manager.NewTooltip(c2resource.Font16, c2resource.PaletteSky, TooltipXCenter, TooltipYBottom)
		t.SetText(v.manager.asset.TranslateString("strParty19")) // need to be verivied
	}

	t.SetVisible(false)
	v.SetTooltip(t)
}

func (v *Button) prerenderStates(btnSprite *Sprite, btnLayout *ButtonLayout, label *Label) {
	numButtonStates := btnSprite.GetFrameCount() / (btnLayout.XSegments * btnLayout.YSegments)

	// buttons always have a base image
	if v.buttonLayout.HasImage {
		btnSprite.RenderSegmented(v.normalSurface, btnLayout.XSegments,
			btnLayout.YSegments, btnLayout.BaseFrame)
	}

	_, labelHeight := label.GetSize()
	textY := half(v.height - labelHeight)
	xOffset := half(v.width)

	label.SetPosition(xOffset, textY)
	label.Render(v.normalSurface)

	if !btnLayout.AllowFrameChange {
		return
	}

	xSeg, ySeg, baseFrame := btnLayout.XSegments, btnLayout.YSegments, btnLayout.BaseFrame

	buttonStateConfigs := make([]*buttonStateDescriptor, 0)

	// pressed button
	if numButtonStates > buttonStatePressed {
		state := &buttonStateDescriptor{
			baseFrame + buttonStatePressed,
			xOffset - pressedButtonOffset, textY + pressedButtonOffset,
			&v.pressedSurface,
			"failed to render button pressedSurface, err: %v\n",
		}

		buttonStateConfigs = append(buttonStateConfigs, state)
	}

	// toggle button
	if numButtonStates > buttonStateToggled {
		buttonStateConfigs = append(buttonStateConfigs, &buttonStateDescriptor{
			baseFrame + buttonStateToggled,
			xOffset, textY,
			&v.toggledSurface,
			"failed to render button toggledSurface, err: %v\n",
		})
	}

	// pressed+toggled
	if numButtonStates > buttonStatePressedToggled {
		buttonStateConfigs = append(buttonStateConfigs, &buttonStateDescriptor{
			baseFrame + buttonStatePressedToggled,
			xOffset, textY,
			&v.pressedToggledSurface,
			"failed to render button pressedToggledSurface, err: %v\n",
		})
	}

	// disabled button
	if btnLayout.DisabledFrame != -1 {
		disabledState := &buttonStateDescriptor{
			btnLayout.DisabledFrame,
			xOffset, textY,
			&v.disabledSurface,
			"failed to render button disabledSurface, err: %v\n",
		}

		buttonStateConfigs = append(buttonStateConfigs, disabledState)
	}

	for stateIdx, w, h := 0, v.width, v.height; stateIdx < len(buttonStateConfigs); stateIdx++ {
		state := buttonStateConfigs[stateIdx]

		if stateIdx > 1 && btnLayout.ResourceName == c2resource.BuySellButton {
			// Without returning early, the button UI gets all subsequent (unrelated) frames
			// stacked on top. Only 2 frames from this sprite are applicable to the button
			// in question. The presentation is incorrect without this hack!
			continue
		}

		surface := v.manager.renderer.NewSurface(w, h)

		*state.prerenderdestination = surface

		btnSprite.RenderSegmented(*state.prerenderdestination, xSeg, ySeg, state.baseFrame)

		label.SetPosition(state.offsetX, state.offsetY)
		label.Render(*state.prerenderdestination)
	}
}

// OnActivated defines the callback handler for the activate event
func (v *Button) OnActivated(callback func()) {
	v.onClick = callback
}

// Activate calls the on activated callback handler, if any
func (v *Button) Activate() {
	if v.onClick == nil {
		return
	}

	v.onClick()
}

// Render renders the button
func (v *Button) Render(target c2interface.Surface) {
	target.PushFilter(c2enum.FilterNearest)
	defer target.Pop()

	target.PushTranslation(v.x, v.y)
	defer target.Pop()

	switch {
	case !v.enabled:
		target.PushColor(c2util.Color(v.buttonLayout.DisabledColor))
		defer target.Pop()

		if v.toggled {
			target.Render(v.toggledSurface)
		} else if v.buttonLayout.HasImage { // it allows to use SetEnabled(false) for non-image budons
			target.Render(v.disabledSurface)
		}
	case v.toggled && v.pressed:
		target.Render(v.pressedToggledSurface)
	case v.pressed:
		if v.buttonLayout.AllowFrameChange {
			target.Render(v.pressedSurface)
		} else {
			target.Render(v.normalSurface)
		}
	case v.toggled:
		target.Render(v.toggledSurface)
	default:
		target.Render(v.normalSurface)
	}
}

// Toggle negates the toggled state of the button
func (v *Button) Toggle() {
	v.toggled = !v.toggled
}

// GetToggled returns the toggled state of the button
func (v *Button) GetToggled() bool {
	return v.toggled
}

// Advance advances the button state
func (v *Button) Advance(_ float64) error {
	return nil
}

// GetEnabled returns the enabled state
func (v *Button) GetEnabled() bool {
	return v.enabled
}

// SetEnabled sets the enabled state
func (v *Button) SetEnabled(enabled bool) {
	v.enabled = enabled
}

// GetPressed returns the pressed state of the button
func (v *Button) GetPressed() bool {
	return v.pressed
}

// SetPressed sets the pressed state of the button
func (v *Button) SetPressed(pressed bool) {
	v.pressed = pressed
}

// SetVisible sets the pressed state of the button
func (v *Button) SetVisible(visible bool) {
	v.BaseWidget.SetVisible(visible)

	if v.isHovered() && !visible {
		v.hoverEnd()
	}
}

// SetPosition sets the position of the widget
func (v *Button) SetPosition(x, y int) {
	v.BaseWidget.SetPosition(x, y)

	if v.buttonLayout.Tooltip != buttonTooltipNone {
		v.tooltip.SetPosition(x+v.buttonLayout.TooltipXOffset, y+v.buttonLayout.TooltipYOffset)
	}
}

// SetTooltip adds a tooltip to the button
func (v *Button) SetTooltip(t *Tooltip) {
	v.tooltip = t
	v.OnHoverStart(func() { v.tooltip.SetVisible(true) })
	v.OnHoverEnd(func() { v.tooltip.SetVisible(false) })
}

func half(n int) int {
	// nolint:gomnd // half is half
	return n / 2
}

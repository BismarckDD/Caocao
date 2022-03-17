package c2ui

import (
	"sort"

	"github.com/BismarckDD/Caocao/c2common/c2enum"
	"github.com/BismarckDD/Caocao/c2common/c2interface"
	"github.com/BismarckDD/Caocao/c2common/c2util"
)

// UIManager manages a collection of UI elements (buttons, textboxes, labels)
type UIManager struct {
	// asset            *c2asset.AssetManager
	renderer c2interface.Renderer
	// inputManager     c2interface.InputManager
	// audio            c2interface.AudioProvider
	widgets          []Widget
	tooltips         []*Tooltip
	widgetsGroups    []*WidgetGroup
	clickableWidgets []ClickableWidget
	cursorButtons    CursorButton
	CursorX          int
	CursorY          int
	pressedWidget    ClickableWidget
	// clickSfx         c2interface.SoundEffect
	*c2util.Logger
}

// Note: methods for creating buttons and stuff are in their respective files

// Initialize is meant to be called after the game loads all of the necessary files
// for sprites and audio
func (ui *UIManager) Initialize() {
	// sfx, err := ui.audio.LoadSound(c2resource.SFXButtonClick, false, false)
	// if err != nil {
	// 	ui.Fatalf("failed to initialize ui: %v", err)
	// }

	// ui.clickSfx = sfx

	// if err := ui.inputManager.BindHandler(ui); err != nil {
	// 	ui.Fatalf("failed to initialize ui: %v", err)
	// }
}

// Reset resets the state of the UI manager. Typically called for new screens
func (ui *UIManager) Reset() {
	ui.widgets = nil
	ui.clickableWidgets = nil
	ui.pressedWidget = nil
	ui.widgetsGroups = nil
	ui.tooltips = nil
}

func (ui *UIManager) addClickable(widget ClickableWidget) {
	ui.clickableWidgets = append(ui.clickableWidgets, widget)
}

// addWidget adds a widget to the UI manager
func (ui *UIManager) addWidget(widget Widget) {
	// err := ui.inputManager.BindHandler(widget)
	// if err != nil {
	// 	ui.Error(err.Error())
	// }

	clickable, ok := widget.(ClickableWidget)
	if ok {
		ui.addClickable(clickable)
	}

	if widgetGroup, ok := widget.(*WidgetGroup); ok {
		ui.widgetsGroups = append(ui.widgetsGroups, widgetGroup)
	}

	ui.widgets = append(ui.widgets, widget)

	sort.SliceStable(ui.widgets, func(i, j int) bool {
		return ui.widgets[i].GetRenderPriority() < ui.widgets[j].GetRenderPriority()
	})

	widget.bindManager(ui)
}

// addTooltip adds a widget to the UI manager
func (ui *UIManager) addTooltip(t *Tooltip) {
	ui.tooltips = append(ui.tooltips, t)
}

// OnMouseButtonUp is an event handler for input
func (ui *UIManager) OnMouseButtonUp(event c2interface.MouseEvent) bool {
	ui.CursorX, ui.CursorY = event.X(), event.Y()
	if event.Button() == c2enum.MouseButtonLeft {
		ui.cursorButtons |= CursorButtonLeft
		// activate previously pressed widget if cursor is still hovering
		w := ui.pressedWidget

		if w != nil && w.Contains(ui.CursorX, ui.CursorY) && w.GetVisible() && w.GetEnabled() {
			w.Activate()
		}

		// unpress all widgets that are pressed
		for _, w := range ui.clickableWidgets {
			w.SetPressed(false)
		}
	}

	return false
}

// OnMouseMove is the mouse move event handler
func (ui *UIManager) OnMouseMove(event c2interface.MouseMoveEvent) bool {
	for _, w := range ui.widgetsGroups {
		if w.GetVisible() {
			w.OnMouseMove(event.X(), event.Y())
		}
	}

	for _, w := range ui.widgets {
		if w.GetVisible() {
			w.OnMouseMove(event.X(), event.Y())
		}
	}

	return false
}

// OnMouseButtonDown is the mouse button down event handler
func (ui *UIManager) OnMouseButtonDown(event c2interface.MouseEvent) bool {
	ui.CursorX, ui.CursorY = event.X(), event.Y()
	if event.Button() == c2enum.MouseButtonLeft {
		// find and press a widget on screen
		ui.pressedWidget = nil
		for _, w := range ui.clickableWidgets {
			if w.Contains(ui.CursorX, ui.CursorY) && w.GetVisible() && w.GetEnabled() {
				w.SetPressed(true)
				ui.pressedWidget = w
				// ui.clickSfx.Play()

				break
			}
		}
	}

	if event.Button() == c2enum.MouseButtonRight {
		ui.cursorButtons |= CursorButtonRight
	}

	return false
}

// Render renders all of the UI elements
func (ui *UIManager) Render(target c2interface.Surface) {
	for _, widget := range ui.widgets {
		if widget.GetVisible() {
			widget.Render(target)
		}
	}

	for _, tooltip := range ui.tooltips {
		if tooltip.GetVisible() {
			tooltip.Render(target)
		}
	}
}

// Advance updates all of the UI elements
func (ui *UIManager) Advance(elapsed float64) {
	for _, widget := range ui.widgets {
		if widget.GetVisible() {
			err := widget.Advance(elapsed)
			if err != nil {
				ui.Error(err.Error())
			}
		}
	}
}

// CursorButtonPressed determines if the specified button has been pressed
func (ui *UIManager) CursorButtonPressed(button CursorButton) bool {
	return ui.cursorButtons&button > 0
}

// CursorPosition returns the current cursor position
func (ui *UIManager) CursorPosition() (x, y int) {
	return ui.CursorX, ui.CursorY
}

// Renderer returns the renderer for this ui manager
func (ui *UIManager) Renderer() c2interface.Renderer {
	return ui.renderer
}

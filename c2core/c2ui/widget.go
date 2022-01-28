package c2ui

// RenderPriority determines in which order ui elements are drawn.
// The higher the number the later an element is drawn.
type RenderPriority int

// Render priorities that determine the order in which widgets/widgetgroups are
// rendered. The higher the later it is rendered
const (
	RenderPriorityBackground RenderPriority = iota
	RenderPriorityMap
	RenderPriorityUnitStatsPanel
	RenderPriorityConversation
	RenderPriorityInventory
	RenderPriorityMiniMap
	RenderPriorityHelpPanel
	RenderPriorityForeground
)

// Widget defines an object that is a UI widget
type Widget interface {
	Drawable
	SetManager(ui *UIManager)
	GetManager() (ui *UIManager)
	OnMouseMove(x int, y int)
	OnHoverStart(callback func())
	OnHoverEnd(callback func())
	isHovered() bool
	hoverStart()
	hoverEnd()
	Contains(x, y int) (contained bool)
}

// ClickableWidget defines an object that can be clicked
type ClickableWidget interface {
	Widget
	SetEnabled(enabled bool)
	SetPressed(pressed bool)
	GetEnabled() bool
	GetPressed() bool
	OnActivated(callback func())
	Activate()
}

// BaseWidget contains default functionality that all widgets share
type BaseWidget struct {
	manager        *UIManager
	x              int
	y              int
	width          int
	height         int
	renderPriority RenderPriority
	visible        bool

	hovered        bool
	onHoverStartCb func()
	onHoverEndCb   func()
}

// NewBaseWidget creates a new BaseWidget with defaults
func NewBaseWidget(manager *UIManager) *BaseWidget {
	return &BaseWidget{
		manager:        manager,
		x:              0,
		y:              0,
		width:          0,
		height:         0,
		visible:        true,
		renderPriority: RenderPriorityBackground,
	}
}

// this seems a public function
func (widget *BaseWidget) SetManager(manager *UIManager) {
	widget.manager = manager
}

// GetManager returns the uiManager
func (widget *BaseWidget) GetManager() (ui *UIManager) {
	return widget.manager
}

// GetSize returns the size of the widget
func (widget *BaseWidget) GetSize() (width, height int) {
	return widget.width, widget.height
}

// SetPosition sets the position of the widget
func (widget *BaseWidget) SetPosition(x, y int) {
	widget.x, widget.y = x, y
}

// OffsetPosition moves the widget by x and y
func (widget *BaseWidget) OffsetPosition(x, y int) {
	widget.x += x
	widget.y += y
}

// GetPosition returns the position of the widget
func (widget *BaseWidget) GetPosition() (x, y int) {
	return widget.x, widget.y
}

// GetVisible returns whether the widget is visible
func (widget *BaseWidget) GetVisible() (visible bool) {
	return widget.visible
}

// SetVisible make the widget visible, not visible
func (widget *BaseWidget) SetVisible(visible bool) {
	widget.visible = visible
}

// GetRenderPriority returns the order in which this widget is rendered
func (widget *BaseWidget) GetRenderPriority() RenderPriority {
	return widget.renderPriority
}

// SetRenderPriority sets the order in which this widget is rendered
func (widget *BaseWidget) SetRenderPriority(priority RenderPriority) {
	widget.renderPriority = priority
}

// OnHoverStart sets a function that is called if the hovering of the widget starts
func (widget *BaseWidget) OnHoverStart(callback func()) {
	widget.onHoverStartCb = callback
}

// HoverStart is called when the hovering of the widget starts
func (widget *BaseWidget) hoverStart() {
	widget.hovered = true
	if widget.onHoverStartCb != nil {
		widget.onHoverStartCb()
	}
}

// OnHoverEnd sets a function that is called if the hovering of the widget ends
func (widget *BaseWidget) OnHoverEnd(callback func()) {
	widget.onHoverEndCb = callback
}

// hoverEnd is called when the widget hovering ends
func (widget *BaseWidget) hoverEnd() {
	widget.hovered = false
	if widget.onHoverEndCb != nil {
		widget.onHoverEndCb()
	}
}

func (widget *BaseWidget) isHovered() bool {
	return widget.hovered
}

// Contains determines whether a given x,y coordinate lands within a Widget
func (widget *BaseWidget) Contains(x, y int) bool {
	widgetX, widgetY := widget.GetPosition()
	widgetWidth, widgetHeight := widget.GetSize()

	return x >= widgetX && x <= widgetX+widgetWidth && y >= widgetY && y <= widgetY+widgetHeight
}

// OnMouseMove is called when the mouse is moved
func (b *BaseWidget) OnMouseMove(x, y int) {
	if b.Contains(x, y) {
		if !b.isHovered() {
			b.hoverStart()
		}
	} else if b.isHovered() {
		b.hoverEnd()
	}
}

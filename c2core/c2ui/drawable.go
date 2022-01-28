package c2ui

import (
	"github.com/hajimehoshi/ebiten"
)

// Drawable represents an instance that can be drawn
type Drawable interface {
	Render(target ebiten.Image)
	Advance(elapsed float64) error
	GetSize() (width, height int)
	SetPosition(x, y int)
	GetPosition() (x, y int)
	OffsetPosition(xo, yo int)
	GetVisible() bool
	SetVisible(visible bool)
	SetRenderPriority(priority RenderPriority)
	GetRenderPriority() (priority RenderPriority)
}

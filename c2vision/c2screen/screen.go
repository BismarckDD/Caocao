package c2screen

import (
	"github.com/BismarckDD/Caocao/c2common/c2interface"
)

// Screen is an exported interface
type Screen interface{}

// ScreenLoadHandler is an exported interface
type ScreenLoadHandler interface {
	// OnLoad performs all necessary loading to prepare a screen to be shown such as loading assets, placing and binding
	// of ui elements, etc. This loading is done asynchronously. The provided channel will allow implementations to
	// provide progress via Error, Progress, or Done
	OnLoad(loading LoadingState)
}

// ScreenUnloadHandler is an exported interface
type ScreenUnloadHandler interface {
	OnUnload() error
}

// ScreenRenderHandler is an exported interface
type ScreenRenderHandler interface {
	Render(target c2interface.Surface)
}

// ScreenAdvanceHandler is an exported interface
type ScreenAdvanceHandler interface {
	Advance(elapsed float64) error
}

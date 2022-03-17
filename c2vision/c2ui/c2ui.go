package c2ui

import (
	"github.com/BismarckDD/Caocao/c2common/c2interface"
	"github.com/BismarckDD/Caocao/c2common/c2util"
)

// CursorButton represents a mouse button
type CursorButton uint8

const (
	logPrefix = "UI Manager"
)

const (
	// CursorButtonLeft represents the left mouse button
	CursorButtonLeft CursorButton = 1
	// CursorButtonRight represents the right mouse button
	CursorButtonRight CursorButton = 2
)

// HorizontalAlign type, determines alignment along x-axis within a layout
type HorizontalAlign int

// Horizontal alignment types
const (
	HorizontalAlignLeft HorizontalAlign = iota
	HorizontalAlignCenter
	HorizontalAlignRight
)

// NewUIManager creates a UIManager instance with the given input and audio provider
func NewUIManager(
	// asset *c2asset.AssetManager,
	renderer c2interface.Renderer,
	// input c2interface.InputManager,
	l c2util.LogLevel,
	// audio c2interface.AudioProvider,
) *UIManager {
	ui := &UIManager{
		// asset:        asset,
		renderer: renderer,
		// inputManager: input,
		// audio:        audio,
	}

	ui.Logger = c2util.NewLogger()
	ui.Logger.SetPrefix(logPrefix)
	ui.Logger.SetLevel(l)

	return ui
}

package c2ebiten

import (
	"errors"
	"image"

	"github.com/BismarckDD/Caocao/c2common/c2enum"
	"github.com/BismarckDD/Caocao/c2common/c2interface"
	"github.com/BismarckDD/Caocao/c2common/c2util"
	"github.com/BismarckDD/Caocao/c2core/c2config"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth       = 1280
	screenHeight      = 960
	defaultSaturation = 1.0
	defaultBrightness = 1.0
	defaultSkewX      = 0.0
	defaultSkewY      = 0.0
	defaultScaleX     = 1.0
	defaultScaleY     = 1.0
)

type renderCallback = func(surface c2interface.Surface) error

type updateCallback = func() error

// static check that we implement our renderer interface
var _ c2interface.Renderer = &Renderer{}

// Renderer is an implementation of a renderer
type Renderer struct {
	updateCallback // used for update game state.
	renderCallback // used for
	*c2util.GlyphPrinter
	lastRenderError error
}

// Update calls the game's logical update function (the `Advance` method)
func (r *Renderer) Update() error {
	if r.updateCallback == nil {
		return errors.New("no update callback defined for ebiten renderer")
	}

	return r.updateCallback()
}

const drawError = "no render callback defined for ebiten renderer"

// Draw updates the screen with the given *ebiten.Image
func (r *Renderer) Draw(screen *ebiten.Image) {
	r.lastRenderError = nil

	if r.renderCallback == nil {
		r.lastRenderError = errors.New(drawError)
		return
	}

	r.lastRenderError = r.renderCallback(CreateEbitenSurface(r, screen))
}

// Layout returns the renderer screen width and height
func (r *Renderer) Layout(_, _ int) (width, height int) {
	return screenWidth, screenHeight
}

// CreateRenderer creates an ebiten renderer instance
func CreateRenderer(config *c2config.Configuration) (*Renderer, error) {
	result := &Renderer{
		GlyphPrinter: c2util.NewDebugPrinter(),
	}

	if config != nil {
		ebiten.SetCursorMode(ebiten.CursorModeHidden)
		ebiten.SetFullscreen(config.FullScreen)
		ebiten.SetRunnableOnUnfocused(config.RunInBackground)
		ebiten.SetVsyncEnabled(config.VsyncEnabled)
		ebiten.SetMaxTPS(config.TicksPerSecond)
	}

	return result, nil
}

// GetRendererName returns the name of the renderer
func (*Renderer) GetRendererName() string {
	return "EbitenRenderer."
}

// SetWindowIcon sets the icon for the window, visible in the chrome of the window
func (*Renderer) SetWindowIcon(fileName string) {
	_, iconImage, err := ebitenutil.NewImageFromFile(fileName)
	if err == nil {
		ebiten.SetWindowIcon([]image.Image{iconImage})
	}
}

// IsDrawingSkipped returns a bool for whether or not the drawing has been skipped
func (r *Renderer) IsDrawingSkipped() bool {
	return r.lastRenderError != nil
}

// Renderer could be run by ebiten.RunGame
// so besides hitage the Renderer interface, it has all Draw, Layout, Update methods.
// Draw: use renderCallback member to draw (a EbitenSurface)
// Layout: lock the screen Width & Height.
// Update: use updateCallback member to update (is also seen as "advance" outer place)
func (r *Renderer) Run(f renderCallback, u updateCallback, width, height int, title string) error {
	r.renderCallback = f
	r.updateCallback = u

	ebiten.SetWindowTitle(title)
	ebiten.SetWindowResizable(false)
	ebiten.SetWindowSize(width, height)

	return ebiten.RunGame(r)
}

// CreateSurface creates a renderer surface from an existing surface
func (r *Renderer) CreateSurface(surface c2interface.Surface) (c2interface.Surface, error) {
	// surface -> EbitenSurface
	img := surface.(*EbitenSurface).image
	surfaceState := SurfaceState{
		filter:     ebiten.FilterNearest,
		effect:     c2enum.DrawEffectNone,
		saturation: defaultSaturation,
		brightness: defaultBrightness,
		skewX:      defaultSkewX,
		skewY:      defaultSkewY,
		scaleX:     defaultScaleX,
		scaleY:     defaultScaleY,
	}
	result := CreateEbitenSurface(r, img, surfaceState)

	return result, nil
}

// NewSurface creates a new surface
func (r *Renderer) NewSurface(width, height int) c2interface.Surface {
	img := ebiten.NewImage(width, height)

	return CreateEbitenSurface(r, img)
}

// IsFullScreen returns a boolean for whether or not the renderer is currently set to fullscreen
func (r *Renderer) IsFullScreen() bool {
	return ebiten.IsFullscreen()
}

// SetFullScreen sets the renderer to fullscreen, given a boolean
func (r *Renderer) SetFullScreen(fullScreen bool) {
	ebiten.SetFullscreen(fullScreen)
}

// SetVSyncEnabled enables vsync, given a boolean
func (r *Renderer) SetVSyncEnabled(vsync bool) {
	ebiten.SetVsyncEnabled(vsync)
}

// GetVSyncEnabled returns a boolean for whether or not vsync is enabled
func (r *Renderer) GetVSyncEnabled() bool {
	return ebiten.IsVsyncEnabled()
}

// GetCursorPos returns the current cursor position x,y coordinates
func (r *Renderer) GetCursorPos() (x, y int) {
	return ebiten.CursorPosition()
}

// CurrentFPS returns the current frames per second of the renderer
func (r *Renderer) CurrentFPS() float64 {
	return ebiten.CurrentFPS()
}

// ShowPanicScreen shows a panic message in a forever loop
func (r *Renderer) ShowPanicScreen(message string) {
	errorScreen := CreatePanicScreen(message)

	err := ebiten.RunGame(errorScreen)
	if err != nil {
		panic(err)
	}
}

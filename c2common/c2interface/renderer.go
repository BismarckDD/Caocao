package c2interface

type renderCallback = func(Surface) error

type updateCallback = func() error

// Renderer interface defines the functionality of a renderer
type Renderer interface {
	GetRendererName() string
	SetWindowIcon(fileName string)
	Run(r renderCallback, u updateCallback, width, height int, title string) error
	IsDrawingSkipped() bool
	CreateSurface(surface Surface) (Surface, error)
	NewSurface(width, height int) Surface
	IsFullScreen() bool
	SetFullScreen(fullScreen bool)
	SetVSyncEnabled(vsync bool)
	GetVSyncEnabled() bool
	GetCursorPos() (int, int)
	CurrentFPS() float64
	ShowPanicScreen(message string)
	Print(target interface{}, str string) error
	PrintAt(target interface{}, str string, x, y int)
}

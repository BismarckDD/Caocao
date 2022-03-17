package c2config

const (
	NormalScreenX          = 640 // resolution
	NormalScreenY          = 480 // resolution
	NormalMenuHeight       = 80
	HighResScreenX         = 1280
	HighResScreenY         = 960
	HighResMenuHeight      = 160
	HighResBattleScreenX   = 1280
	HighResBattleScreenY   = 960
	OffsetYOfFace          = 80 //
	DefaultFramesPerAction = 12
)

type Configuration struct {
	ScreenX            int
	ScreenY            int
	BattleFieldScreenX int
	BattleFieldScreenY int
	MenuHeight         int
	FramesPerAction    int
	FullScreen         bool
	RunInBackground    bool
	VsyncEnabled       bool
	TicksPerSecond     int
}

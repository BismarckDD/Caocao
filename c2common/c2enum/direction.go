package c2enum

// AnimationFrameDirection enumerates animation frame directions used in
// c2datadict.UnitSequenceFrame
type AnimationFrameDirection int

// Animation frame directions
const (
	Up AnimationFrameDirection = iota
	Right
	Down
	Left
)

package c2interface

import "github.com/BismarckDD/Caocao/c2common/c2enum"

// InputService represents an interface offering Keyboard and Mouse interactions.
type InputService interface {
	// CursorPosition returns a position of a mouse cursor relative to the game screen (window).
	CursorPosition() (x int, y int)
	// InputChars return "printable" runes read from the keyboard at the time update is called.
	InputChars() []rune
	// IsKeyPressed checks if the provided key is down.
	IsKeyPressed(key c2enum.Key) bool
	// IsKeyJustPressed checks if the provided key is just transitioned from up to down.
	IsKeyJustPressed(key c2enum.Key) bool
	// IsKeyJustReleased checks if the provided key is just transitioned from down to up.
	IsKeyJustReleased(key c2enum.Key) bool
	// IsMouseButtonPressed checks if the provided mouse button is down.
	IsMouseButtonPressed(button c2enum.MouseButton) bool
	// IsMouseButtonJustPressed checks if the provided mouse button is just transitioned from up to down.
	IsMouseButtonJustPressed(button c2enum.MouseButton) bool
	// IsMouseButtonJustReleased checks if the provided mouse button is just transitioned from down to up.
	IsMouseButtonJustReleased(button c2enum.MouseButton) bool
	// KeyPressDuration returns how long the key is pressed in frames.
	KeyPressDuration(key c2enum.Key) int
}

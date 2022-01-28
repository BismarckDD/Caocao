package c2interface

import "github.com/BismarckDD/Caocao/c2common/c2enum"

// HandlerEvent holds the qualifiers for a key or mouse event
type HandlerEvent interface {
	KeyMod() c2enum.KeyMod
	ButtonMod() c2enum.MouseButtonMod
	X() int
	Y() int
}

// KeyEvent represents an event associated with a keyboard key
type KeyEvent interface {
	HandlerEvent
	Key() c2enum.Key
	// Duration represents the number of frames this key has been pressed for
	Duration() int
}

// KeyCharsEvent represents an event associated with a keyboard character being pressed
type KeyCharsEvent interface {
	HandlerEvent
	Chars() []rune
}

// MouseEvent represents a mouse event
type MouseEvent interface {
	HandlerEvent
	Button() c2enum.MouseButton
}

// MouseMoveEvent represents a mouse movement event
type MouseMoveEvent interface {
	HandlerEvent
}

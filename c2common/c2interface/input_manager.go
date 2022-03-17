package c2interface

import "github.com/BismarckDD/Caocao/c2common/c2enum"

// InputManager manages an InputService
type InputManager interface {
	Advance(elapsedTime, currentTime float64) error
	BindHandlerWithPriority(InputEventHandler, c2enum.Priority) error
	BindHandler(h InputEventHandler) error
	UnbindHandler(handler InputEventHandler) error
}

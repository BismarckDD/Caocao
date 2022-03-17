package c2asset

import (
	"errors"
	"image"
	"image/color"
	"log"
	"math"

	"github.com/BismarckDD/Caocao/c2common/c2enum"
	"github.com/BismarckDD/Caocao/c2common/c2interface"
	"github.com/BismarckDD/Caocao/c2common/c2math"
)

type playMode int

const (
	playModePause playMode = iota
	playModeForward
	playModeBackward
)

const defaultPlayLength = 1.0

type animationFrame struct {
	decoded bool

	width   int
	height  int
	offsetX int
	offsetY int

	image c2interface.Surface
}

type animationDirection struct {
	decoded bool
	frames  []animationFrame
}

// static check that we implement the animation interface
var _ c2interface.Animation = &Animation{}

// Animation has directionality, play modes, and frame counting
type Animation struct {
	renderer       c2interface.Renderer
	onBindRenderer func(renderer c2interface.Renderer) error
	directions     []animationDirection
	effect         c2enum.DrawEffect
	colorMod       color.Color
	frameIndex     int
	directionIndex int
	lastFrameTime  float64
	playedCount    int
	playMode
	playLength       float64
	subStartingFrame int
	subEndingFrame   int
	originAtBottom   bool
	playLoop         bool
	hasSubLoop       bool // runs after first animation ends
	hasShadow        bool
}

// SetSubLoop sets a sub loop for the animation
func (a *Animation) SetSubLoop(startFrame, endFrame int) {
	a.subStartingFrame = startFrame
	a.subEndingFrame = endFrame
	a.hasSubLoop = true
}

// Advance advances the animation state
func (a *Animation) Advance(elapsed float64) error {
	if a.playMode == playModePause {
		return nil
	}

	frameCount := a.GetFrameCount()
	frameLength := a.playLength / float64(frameCount)
	a.lastFrameTime += elapsed
	framesAdvanced := int(a.lastFrameTime / frameLength)
	a.lastFrameTime -= float64(framesAdvanced) * frameLength

	for i := 0; i < framesAdvanced; i++ {
		startIndex := 0
		endIndex := frameCount

		if a.hasSubLoop && a.playedCount > 0 {
			startIndex = a.subStartingFrame
			endIndex = a.subEndingFrame
		}

		switch a.playMode {
		case playModeForward:
			a.frameIndex++
			if a.frameIndex >= endIndex {
				a.playedCount++
				if a.playLoop {
					a.frameIndex = startIndex
				} else {
					a.frameIndex = endIndex - 1
					break
				}
			}
		case playModeBackward:
			a.frameIndex--
			if a.frameIndex < startIndex {
				a.playedCount++
				if a.playLoop {
					a.frameIndex = endIndex - 1
				} else {
					a.frameIndex = startIndex
					break
				}
			}
		}
	}

	return nil
}

const (
	full = 1.0
	half = 0.5
	zero = 0.0
)

func (a *Animation) renderShadow(target c2interface.Surface) {
	direction := a.directions[a.directionIndex]
	frame := direction.frames[a.frameIndex]

	target.PushFilter(c2enum.FilterLinear)
	defer target.Pop()

	target.PushTranslation(frame.offsetX, int(float64(frame.offsetY)*half))
	defer target.Pop()

	target.PushScale(full, half)
	defer target.Pop()

	target.PushSkew(half, zero)
	defer target.Pop()

	target.PushEffect(c2enum.DrawEffectPctTransparency25)
	defer target.Pop()

	target.PushBrightness(zero)
	defer target.Pop()

	target.Render(frame.image)
}

// GetCurrentFrameSurface returns the surface for the current frame of the
// animation
func (a *Animation) GetCurrentFrameSurface() c2interface.Surface {
	return a.directions[a.directionIndex].frames[a.frameIndex].image
}

// Render renders the animation to the given surface
func (a *Animation) Render(target c2interface.Surface) {
	if a.renderer == nil {
		a.BindRenderer(target.Renderer())
	}

	direction := a.directions[a.directionIndex]
	frame := direction.frames[a.frameIndex]

	target.PushTranslation(frame.offsetX, frame.offsetY)
	defer target.Pop()

	target.PushEffect(a.effect)
	defer target.Pop()

	target.PushColor(a.colorMod)
	defer target.Pop()

	target.Render(frame.image)
}

// BindRenderer binds the given renderer to the animation so that it can initialize
// the required surfaces
func (a *Animation) BindRenderer(r c2interface.Renderer) {
	if a.onBindRenderer == nil {
		return
	}

	if err := a.onBindRenderer(r); err != nil {
		log.Println(err)
	}
}

// RenderFromOrigin renders the animation from the animation origin
func (a *Animation) RenderFromOrigin(target c2interface.Surface, shadow bool) {
	if a.renderer == nil {
		a.BindRenderer(target.Renderer())
	}

	if a.originAtBottom {
		direction := a.directions[a.directionIndex]
		frame := direction.frames[a.frameIndex]
		target.PushTranslation(0, -frame.height)

		defer target.Pop()
	}

	if shadow && !a.effect.Transparent() && a.hasShadow {
		_, height := a.GetFrameBounds()
		height = int(math.Abs(float64(height)))
		halfHeight := height / 2 //nolint:gomnd // this ain't rocket surgery...

		target.PushTranslation(-halfHeight, 0)
		defer target.Pop()

		a.renderShadow(target)

		return
	}

	a.Render(target)
}

// RenderSection renders the section of the animation frame enclosed by bounds
func (a *Animation) RenderSection(target c2interface.Surface, bound image.Rectangle) {
	if a.renderer == nil {
		a.BindRenderer(target.Renderer())
	}

	direction := a.directions[a.directionIndex]
	frame := direction.frames[a.frameIndex]

	target.PushTranslation(frame.offsetX, frame.offsetY)
	defer target.Pop()

	target.PushEffect(a.effect)
	defer target.Pop()

	target.PushColor(a.colorMod)
	defer target.Pop()

	target.RenderSection(frame.image, bound)
}

// GetFrameSize gets the Size(width, height) of a indexed frame.
func (a *Animation) GetFrameSize(frameIndex int) (width, height int, err error) {
	direction := a.directions[a.directionIndex]
	if frameIndex >= len(direction.frames) {
		return 0, 0, errors.New("invalid frame index")
	}

	frame := direction.frames[frameIndex]

	return frame.width, frame.height, nil
}

// GetCurrentFrameSize gets the Size(width, height) of the current frame.
func (a *Animation) GetCurrentFrameSize() (width, height int) {
	width, height, err := a.GetFrameSize(a.frameIndex)
	if err != nil {
		log.Print(err)
	}

	return width, height
}

// GetFrameBounds gets maximum Size(width, height) of all frame.
func (a *Animation) GetFrameBounds() (maxWidth, maxHeight int) {
	maxWidth, maxHeight = 0, 0

	direction := a.directions[a.directionIndex]

	for _, frame := range direction.frames {
		maxWidth = c2math.MaxInt(maxWidth, frame.width)
		maxHeight = c2math.MaxInt(maxHeight, frame.height)
	}

	return maxWidth, maxHeight
}

// GetCurrentFrame gets index of current frame in animation
func (a *Animation) GetCurrentFrame() int {
	return a.frameIndex
}

// GetFrameCount gets number of frames in animation
func (a *Animation) GetFrameCount() int {
	direction := a.directions[a.directionIndex]
	return len(direction.frames)
}

// IsOnFirstFrame gets if the animation on its first frame
func (a *Animation) IsOnFirstFrame() bool {
	return a.frameIndex == 0
}

// IsOnLastFrame gets if the animation on its last frame
func (a *Animation) IsOnLastFrame() bool {
	return a.frameIndex == a.GetFrameCount()-1
}

// GetDirectionCount gets the number of animation direction
func (a *Animation) GetDirectionCount() int {
	return len(a.directions)
}

// SetDirection places the animation in the direction of an animation
func (a *Animation) SetDirection(directionIndex int) error {
	const smallestInvalidDirectionIndex = 64
	if directionIndex >= smallestInvalidDirectionIndex {
		return errors.New("invalid direction index")
	}

	a.directionIndex = directionIndex
	a.frameIndex = 0

	return nil
}

// GetDirection get the current animation direction
func (a *Animation) GetDirection() int {
	return a.directionIndex
}

// SetCurrentFrame sets animation at a specific frame
func (a *Animation) SetCurrentFrame(frameIndex int) error {
	if frameIndex >= a.GetFrameCount() {
		return errors.New("invalid frame index")
	}

	a.frameIndex = frameIndex
	a.lastFrameTime = 0

	return nil
}

// Rewind animation to beginning
func (a *Animation) Rewind() {
	err := a.SetCurrentFrame(0)
	if err != nil {
		log.Print(err)
	}
}

// PlayForward plays animation forward
func (a *Animation) PlayForward() {
	a.playMode = playModeForward
	a.lastFrameTime = 0
}

// PlayBackward plays animation backward
func (a *Animation) PlayBackward() {
	a.playMode = playModeBackward
	a.lastFrameTime = 0
}

// Pause animation
func (a *Animation) Pause() {
	a.playMode = playModePause
	a.lastFrameTime = 0
}

// SetPlayLoop sets whether to loop the animation
func (a *Animation) SetPlayLoop(loop bool) {
	a.playLoop = loop
}

// SetPlaySpeed sets play speed of the animation
func (a *Animation) SetPlaySpeed(playSpeed float64) {
	a.SetPlayLength(playSpeed * float64(a.GetFrameCount()))
}

// SetPlayLength sets the Animation's play length in seconds
func (a *Animation) SetPlayLength(playLength float64) {
	a.playLength = playLength
	a.lastFrameTime = 0
}

// SetColorMod sets the Animation's color mod
func (a *Animation) SetColorMod(colorMod color.Color) {
	a.colorMod = colorMod
}

// GetPlayedCount gets the number of times the application played
func (a *Animation) GetPlayedCount() int {
	return a.playedCount
}

// ResetPlayedCount resets the play count
func (a *Animation) ResetPlayedCount() {
	a.playedCount = 0
}

// SetEffect sets the draw effect for the animation
func (a *Animation) SetEffect(e c2enum.DrawEffect) {
	a.effect = e
}

// SetShadow sets bool for whether or not to draw a shadow
func (a *Animation) SetShadow(shadow bool) {
	a.hasShadow = shadow
}

// Clone creates a copy of the Animation
func (a *Animation) Clone() c2interface.Animation {
	clone := *a
	copy(clone.directions, a.directions)

	return &clone
}

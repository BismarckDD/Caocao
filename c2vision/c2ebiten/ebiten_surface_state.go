package c2ebiten

import (
	"image/color"

	"github.com/BismarckDD/Caocao/c2common/c2enum"
	"github.com/hajimehoshi/ebiten/v2"
)

type SurfaceState struct {
	x              int
	y              int
	filter         ebiten.Filter
	color          color.Color
	brightness     float64
	saturation     float64
	effect         c2enum.DrawEffect
	skewX, skewY   float64
	scaleX, scaleY float64
}

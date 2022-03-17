package c2ebiten

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/BismarckDD/Caocao/c2common/c2enum"
)

// Filter参数转换，Filter用于ebiten.LoadImage方法的参数，用途尚不明确
func ToEbitenFilter(filter c2enum.Filter) ebiten.Filter {
	switch filter {
	case c2enum.FilterNearest:
		return ebiten.FilterNearest
	default:
		return ebiten.FilterLinear
	}
}

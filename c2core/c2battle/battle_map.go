package c2battle

import "image"

type BattleMap struct {
	Width            uint16 // battleImage.width() / 64
	Height           uint16 // battleImage.height() / 64
	BattleImage      *image.Image
	SmallBattleImage *image.Image
	MapTerrain       [][]uint32
}

func CreateBattleMap(battleId uint16) *BattleMap {
	var battleImage *image.Image = nil
	var smallBattleImage *image.Image = nil
	width, height = battleImage. // px
	return &BattleMap{
		Width:            width,
		Height:           height,
		BattleImage:      battleImage,
		SmallBattleImage: smallBattleImage,
	}
}

package c2char

import "image"

var BattleFieldFigureList = make([]BattleFieldFigure, 200)

type BattleFieldFigure struct {
	StandImage       []image.Image // 2
	MoveDownImage    []image.Image // 2
	MoveUpImage      []image.Image // 2
	MoveLeftImage    []image.Image // 2
	MoveRightImage   []image.Image // 2
	AttackDownImage  []image.Image
	AttackUpImage    []image.Image
	AttackLeftImage  []image.Image
	AttackRightImage []image.Image
	DefendDownImage  []image.Image
	DefendUpImage    []image.Image
	DefendLeftImage  []image.Image
	DefendRightImage []image.Image
	DeadImage        []image.Image
	ExcitationImage  []image.Image
}

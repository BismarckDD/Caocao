package c2char

import "image"

var ScenarioFigureList = make([]ScenarioFigure, 200)

type ScenarioFigure struct {
	StandLeftUp    []image.Image
	StandLeftDown  []image.Image
	StandRightUp   []image.Image
	StandRightDOwn []image.Image
	MoveLeftUp     []image.Image
	MoveLeftDown   []image.Image
	MoveRightUp    []image.Image
	MoveRightDown  []image.Image
}

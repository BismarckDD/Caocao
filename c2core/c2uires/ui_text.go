package c2uires

import (
	"image/color"

	"golang.org/x/image/font"
)

type TextUI struct {
	IdleColor     color.Color
	DisabledColor color.Color
	Face          font.Face
	TitleFace     font.Face
	BigTitleFace  font.Face
	SmallFace     font.Face
}

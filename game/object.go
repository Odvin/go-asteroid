package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Object struct {
	sprite     *ebiten.Image
	position   Vector
	width      float64
	height     float64
	halfWidth  float64
	halfHeight float64
	rotation   float64
}

func NewObject(sprite *ebiten.Image) *Object {
	width := float64(sprite.Bounds().Dx())
	height := float64(sprite.Bounds().Dy())

	halfWidth := width / 2
	halfHeight := height / 2

	return &Object{
		sprite:     sprite,
		width:      width,
		height:     height,
		halfWidth:  halfWidth,
		halfHeight: halfHeight,
		rotation:   0,
	}
}

package game

import (
	"math"

	"github.com/Odvin/go-asteroid/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	bulletSpeedPerSecond = 350.0
)

type Bullet struct {
	sprite     *ebiten.Image
	width      float64
	height     float64
	halfWidth  float64
	halfHeight float64
	position   Vector
	rotation   float64
}

func NewBullet(position Vector, rotation float64) *Bullet {
	sprite := assets.LaserSprite

	bounds := sprite.Bounds()
	width := float64(bounds.Dx())
	height := float64(bounds.Dx())
	halfWidth := width / 2
	halfHeight := height / 2

	position.X -= halfWidth
	position.Y -= halfHeight

	return &Bullet{
		sprite:     sprite,
		width:      width,
		height:     height,
		halfWidth:  halfWidth,
		halfHeight: halfHeight,
		position:   position,
		rotation:   rotation,
	}
}

func (b *Bullet) Update() {
	speed := bulletSpeedPerSecond / float64(ebiten.TPS())

	b.position.X += math.Sin(b.rotation) * speed
	b.position.Y += math.Cos(b.rotation) * -speed
}

func (b *Bullet) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(-b.halfWidth, -b.halfHeight)
	op.GeoM.Rotate(b.rotation)
	op.GeoM.Translate(b.halfWidth, b.halfHeight)

	op.GeoM.Translate(b.position.X, b.position.Y)

	screen.DrawImage(b.sprite, op)
}

func (b *Bullet) Collider() Rect {
	return NewRect(b.position.X, b.position.Y, b.width, b.height)
}

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
	Object
	position Vector
}

func NewBullet(position Vector, rotation float64) *Bullet {
	object := NewObject(assets.LaserSprite)
	object.rotation = rotation

	position.X -= object.halfWidth
	position.Y -= object.halfHeight

	return &Bullet{
		Object:   *object,
		position: position,
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

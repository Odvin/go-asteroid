package game

import (
	"math"
	"math/rand/v2"

	"github.com/Odvin/go-asteroid/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	rotationSpeedMin = -0.02
	rotationSpeedMax = 0.02
)

type Meteor struct {
	sprite        *ebiten.Image
	position      Vector
	width         float64
	height        float64
	halfWidth     float64
	halfHeight    float64
	rotation      float64
	rotationSpeed float64
	movement      Vector
}

func NewMeteor(baseVelocity float64) *Meteor {
	sprite := assets.MeteorSprites[rand.IntN(len(assets.MeteorSprites))]

	bounds := sprite.Bounds()

	width := float64(bounds.Dx())
	halfWidth := width / 2
	height := float64(bounds.Dy())
	halfHeight := height / 2

	target := Vector{
		X: screenWidth / 2,
		Y: screenHeight / 2,
	}

	angle := rand.Float64() * 2 * math.Pi

	r := screenWidth / 2.0

	position := Vector{
		X: target.X + math.Cos(angle)*r,
		Y: target.Y + math.Sin(angle)*r,
	}

	velocity := baseVelocity + rand.Float64()*1.5

	direction := Vector{
		X: target.X - position.X,
		Y: target.Y - position.Y,
	}

	normalizedDirection := direction.Normalize()

	movement := Vector{
		X: normalizedDirection.X * velocity,
		Y: normalizedDirection.Y * velocity,
	}

	return &Meteor{
		sprite:        sprite,
		position:      position,
		width:         width,
		height:        height,
		halfWidth:     halfWidth,
		halfHeight:    halfHeight,
		rotation:      0,
		rotationSpeed: rotationSpeedMin + rand.Float64()*(rotationSpeedMax-rotationSpeedMin),
		movement:      movement,
	}
}

func (m *Meteor) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-m.halfWidth, -m.halfHeight)
	op.GeoM.Rotate(m.rotation)
	op.GeoM.Translate(m.halfWidth, m.halfHeight)

	op.GeoM.Translate(m.position.X, m.position.Y)

	screen.DrawImage(m.sprite, op)
}

func (m *Meteor) Update() {
	m.position.X += m.movement.X
	m.position.Y += m.movement.Y
	m.rotation += m.rotationSpeed
}

func (m *Meteor) Collider() Rect {
	return NewRect(m.position.X, m.position.Y, m.width, m.height)
}

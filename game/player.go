package game

import (
	"math"
	"time"

	"github.com/Odvin/go-asteroid/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	rotationPerSecond = math.Pi
	shootCooldown     = time.Millisecond * 500
	bulletSpawnOffset = 50.0
)

type Player struct {
	sprite     *ebiten.Image
	position   Vector
	width      float64
	height     float64
	halfWidth  float64
	halfHeight float64
	rotation   float64

	shootCooldown *Timer
	bullets       []*Bullet
}

func NewPlayer() *Player {
	sprite := assets.PlayerSprite

	bounds := sprite.Bounds()
	width := float64(bounds.Dx())
	halfWidth := width / 2
	height := float64(bounds.Dy())
	halfHeight := height / 2

	position := Vector{
		X: screenWidth/2 - halfWidth,
		Y: screenHeight/2 - halfHeight,
	}

	return &Player{
		sprite:        sprite,
		position:      position,
		width:         width,
		height:        height,
		halfWidth:     halfWidth,
		halfHeight:    halfHeight,
		rotation:      0,
		shootCooldown: NewTimer(shootCooldown),
	}
}

func (p *Player) Update() {
	speed := rotationPerSecond / float64(ebiten.TPS())

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.rotation -= speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.rotation += speed
	}

	p.shootCooldown.Update()
	if p.shootCooldown.IsReady() && ebiten.IsKeyPressed(ebiten.KeySpace) {
		p.shootCooldown.Reset()

		shootPosition := Vector{
			p.position.X + p.halfWidth + math.Sin(p.rotation)*bulletSpawnOffset,
			p.position.Y + p.halfHeight + math.Cos(p.rotation)*-bulletSpawnOffset,
		}

		bullet := NewBullet(shootPosition, p.rotation)
		p.bullets = append(p.bullets, bullet)
	}

	for _, b := range p.bullets {
		b.Update()
	}

}

func (p *Player) Draw(screen *ebiten.Image) {

	op := &ebiten.DrawImageOptions{}

	if p.rotation != 0 {
		op.GeoM.Translate(-p.halfWidth, -p.halfHeight)
		op.GeoM.Rotate(p.rotation)
		op.GeoM.Translate(p.halfWidth, p.halfHeight)
	}

	op.GeoM.Translate(p.position.X, p.position.Y)

	screen.DrawImage(p.sprite, op)

	for _, b := range p.bullets {
		b.Draw(screen)
	}
}

func (p *Player) Collider() Rect {
	return NewRect(p.position.X, p.position.Y, p.width, p.height)
}

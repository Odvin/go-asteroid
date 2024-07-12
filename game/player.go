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
	Object
	position Vector

	shootCooldown *Timer
	bullets       []*Bullet
}

func NewPlayer() *Player {
	object := NewObject(assets.PlayerSprite)

	position := Vector{
		X: screenWidth/2 - object.halfWidth,
		Y: screenHeight/2 - object.halfHeight,
	}

	return &Player{
		Object:   *object,
		position: position,

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

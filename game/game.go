package game

import (
	"fmt"
	"image/color"
	"time"

	"github.com/Odvin/go-asteroid/assets"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

const (
	screenWidth  = 800
	screenHeight = 600

	meteorSpawnTime = 1 * time.Second

	baseMeteorVelocity = 0.25
)

type Game struct {
	player *Player

	score int

	baseVelocity     float64
	meteorSpawnTimer *Timer
	meteors          []*Meteor
}

func NewGame() *Game {
	return &Game{
		player:           NewPlayer(),
		baseVelocity:     baseMeteorVelocity,
		meteorSpawnTimer: NewTimer(meteorSpawnTime),
	}
}

func (g *Game) Update() error {
	g.player.Update()

	g.meteorSpawnTimer.Update()
	if g.meteorSpawnTimer.IsReady() {
		g.meteorSpawnTimer.Reset()

		m := NewMeteor(g.baseVelocity)
		g.meteors = append(g.meteors, m)
	}

	for _, m := range g.meteors {
		m.Update()
	}

	// Check for meteor/bullet collisions
	for i, m := range g.meteors {
		for j, b := range g.player.bullets {
			if m.Collider().Intersects(b.Collider()) {
				g.meteors = append(g.meteors[:i], g.meteors[i+1:]...)
				g.player.bullets = append(g.player.bullets[:j], g.player.bullets[j+1:]...)
				g.score++
			}
		}
	}

	// Check for meteor/player collisions
	for _, m := range g.meteors {
		if m.Collider().Intersects(g.player.Collider()) {
			g.Reset()
			break
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)

	for _, m := range g.meteors {
		m.Draw(screen)
	}

	op := &text.DrawOptions{}
	op.GeoM.Translate(10, 50)
	op.ColorScale.ScaleWithColor(color.White)
	text.Draw(screen, fmt.Sprintf("%06d", g.score), &text.GoTextFace{
		Source: assets.ScoreFont,
		Size:   24,
	}, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) Reset() {
	g.player = NewPlayer()
	g.meteors = nil
	g.score = 0
	g.meteorSpawnTimer.Reset()
}

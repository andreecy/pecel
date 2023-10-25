package main

import (
	"github.com/andreecy/pecel"
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	Position pecel.Vec2
}

func (p *Player) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		p.Position.X -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		p.Position.X += 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		p.Position.Y -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		p.Position.Y += 1
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	// draw character
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.Position.X, p.Position.Y)
	// centering image
	op.GeoM.Translate(-16, -16)
	screen.DrawImage(charImage, op)
}

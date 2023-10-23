package main

import (
	"bytes"
	"fmt"
	"image"
  _ "image/png"
	"log"

	"github.com/andreecy/pecel"
	"github.com/andreecy/pecel/examples/resources/images"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type GameplayScene struct {
	playerPos  pecel.Vec2
	keyPressed []ebiten.Key
	pecel.BasicScene
}

var (
	tilesImage *ebiten.Image
	charImage  *ebiten.Image
)

func init() {
	img, _, err := image.Decode(bytes.NewReader(images.Grass_png))
	if err != nil {
		log.Fatal(err)
	}
	tilesImage = ebiten.NewImageFromImage(img)

	charImg, _, err := image.Decode(bytes.NewReader(images.Character_png))
	if err != nil {
		log.Fatal(err)
	}
	charImage = ebiten.NewImageFromImage(charImg)
}

func (s *GameplayScene) Update(state *pecel.GameState) error {
	if s.World == nil {
		s.World = ebiten.NewImage(200, 200)
	}

	if s.Camera == nil {
		s.Camera = pecel.CreateCamera(screenWidth, screenHeight)
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		s.playerPos.X -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		s.playerPos.X += 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		s.playerPos.Y -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		s.playerPos.Y += 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		s.Camera.ZoomFactor -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyE) {
		s.Camera.ZoomFactor += 1
	}

	// follow player
	s.Camera.Position = s.playerPos
	return nil
}

func (s *GameplayScene) Draw(screen *ebiten.Image) {
	s.World.Clear()

	// draw grounds
	op := &ebiten.DrawImageOptions{}
	s.World.DrawImage(tilesImage, op)

	// draw character
	charOp := &ebiten.DrawImageOptions{}
	charOp.GeoM.Translate(s.playerPos.X, s.playerPos.Y)
	charOp.GeoM.Translate(-16, -16)
	s.World.DrawImage(charImage, charOp)

	s.Camera.Render(s.World, screen)

	// debug
	wordPos := s.Camera.ScreenToWorld(pecel.NewVec2(ebiten.CursorPosition()))
	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %0.1f", ebiten.ActualFPS()))
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%s\nCursor World Pos: %.2f,%.2f", s.Camera.String(), wordPos.X, wordPos.Y), 0, screenHeight-32)
}

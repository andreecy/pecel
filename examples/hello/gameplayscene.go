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

type GameplayScene struct{}

var (
	charImage *ebiten.Image
)

func init() {
	img, _, err := image.Decode(bytes.NewReader(images.Character_png))
	if err != nil {
		log.Fatal(err)
	}
	charImage = ebiten.NewImageFromImage(img)
}

func (s *GameplayScene) Update(state *pecel.GameState) error {

	return nil
}

func (s *GameplayScene) Draw(screen *ebiten.Image) {
	// draw sprite
	op := &ebiten.DrawImageOptions{}
  // center to the screen
	op.GeoM.Translate(screenWidth/2, screenHeight/2)
  // center of sprite
	op.GeoM.Translate(-float64(charImage.Bounds().Dx())/2, -float64(charImage.Bounds().Dy())/2)
	screen.DrawImage(charImage, op)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %0.1f", ebiten.ActualFPS()))
}

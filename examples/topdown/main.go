package main

import (
	_ "image/png"
	"log"

	"github.com/andreecy/pecel"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 320
	screenHeight = 240
)


type Game struct {
	sceneManager *pecel.SceneManager
}

func (g *Game) Update() error {
	if g.sceneManager == nil {
		g.sceneManager = &pecel.SceneManager{}
		g.sceneManager.LoadScene(&GameplayScene{})
	}
	if err := g.sceneManager.Update(); err != nil {
		return err
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.sceneManager.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	g := &Game{}

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

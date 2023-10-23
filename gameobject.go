package pecel

import "github.com/hajimehoshi/ebiten/v2"

type GameObject interface {
	Draw(*ebiten.Image) error
}

// BasicObject is bare implementation of GameObject
type BasicObject struct {
	Width, Height int
	Position      *Vec2
	Scale         *Vec2
	Rotation      int
}

func (obj *BasicObject) Draw(screen *ebiten.Image) error {
  return nil
}

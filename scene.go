package pecel

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Scene interface {
	Update(state *GameState) error
	Draw(screen *ebiten.Image)
}

type SceneManager struct {
	current Scene
}

type GameState struct {
	SceneManager *SceneManager
}

// do update current active scene
func (s *SceneManager) Update() error {
	s.current.Update(&GameState{
		SceneManager: s,
	})

	return nil
}

// draw current active scene
func (s *SceneManager) Draw(r *ebiten.Image) {
	s.current.Draw(r)
}

// load a scene
func (s *SceneManager) LoadScene(scene Scene) {
	s.current = scene
}

type BasicScene struct {
	World       *ebiten.Image
	Camera      *Camera
	GameObjects []GameObject
}

func (s *BasicScene) Update() {
}

func (s *BasicScene) Draw(screen *ebiten.Image) {
	s.World.Clear()
	s.World.Fill(color.White)

	for _, obj := range s.GameObjects {
		obj.Draw(s.World)
	}
	s.Camera.Render(s.World, screen)
}

package pecel

import (
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

package pecel

import (
	// "image/color"

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

// type BasicScene struct {
// 	world       *ebiten.Image
// 	camera      *Camera
// 	gameObjects []GameObject
// }
//
// func (s *BasicScene) Update() {
// }
//
// func (s *BasicScene) Draw(screen *ebiten.Image) {
// 	s.world.Clear()
// 	s.world.Fill(color.White)
//
// 	for _, obj := range s.gameObjects {
// 		obj.Draw(s.world)
// 	}
// 	s.camera.Render(s.world, screen)
// }

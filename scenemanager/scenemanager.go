package scenemanager

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Scene is an interface with Draw and Update.
type Scene interface {
	Update(sceneManager *SceneManager) error
	Draw(r *ebiten.Image)
}

// SceneManager is a struct that holds the current scene and next scene
// Use it as a member variable of ebiten.Game interface.
type SceneManager struct {
	current        Scene
	next           Scene
	transitionFlag bool
}

// Same as Update() for ebiten.Game .
func (s *SceneManager) Update() error {

	if !s.transitionFlag {
		return s.current.Update(s)
	}

	s.transitionFlag = false

	s.current = s.next
	s.next = nil
	return nil
}

// Same as Draw() for ebiten.Game .
func (s *SceneManager) Draw(r *ebiten.Image) {
	s.current.Draw(r)
}

// Change current in SceneManager to the scene of the argument.
func (s *SceneManager) GoTo(scene Scene) {
	if s.current == nil {
		s.current = scene
	} else {
		s.next = scene
		s.transitionFlag = true
	}
}

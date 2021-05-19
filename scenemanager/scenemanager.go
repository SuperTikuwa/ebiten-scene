package scenemanager

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Scene interface {
	Update(state *GameState) error
	Draw(r *ebiten.Image)
}

type SceneManager struct {
	current        Scene
	next           Scene
	transitionFlag bool
}

type GameState struct {
	SceneManager *SceneManager
}

func (s *SceneManager) Update() error {

	if !s.transitionFlag {
		return s.current.Update(&GameState{SceneManager: s})
	}

	s.transitionFlag = false

	s.current = s.next
	s.next = nil
	return nil
}

func (s *SceneManager) Draw(r *ebiten.Image) {
	s.current.Draw(r)
}

func (s *SceneManager) GoTo(scene Scene) {
	if s.current == nil {
		s.current = scene
	} else {
		s.next = scene
		s.transitionFlag = true
	}
}

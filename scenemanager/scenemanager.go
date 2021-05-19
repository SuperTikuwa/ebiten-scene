package scenemanager

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// SceneはDrawとUpdateを持つインタフェースです．
type Scene interface {
	Update(sceneManager *SceneManager) error
	Draw(r *ebiten.Image)
}

// SceneManagerは現在のシーンと次のシーンを持つ構造体
// ebiten.Game interfaceのメンバ変数にして使う
type SceneManager struct {
	current        Scene
	next           Scene
	transitionFlag bool
}

// ebiten.GameのUpdateと同じ．
func (s *SceneManager) Update() error {

	if !s.transitionFlag {
		return s.current.Update(s)
	}

	s.transitionFlag = false

	s.current = s.next
	s.next = nil
	return nil
}

// ebiten.GameのDrawと同じ．
func (s *SceneManager) Draw(r *ebiten.Image) {
	s.current.Draw(r)
}

// SceneManagerのcurrentを引数のsceneに変更する．
func (s *SceneManager) GoTo(scene Scene) {
	if s.current == nil {
		s.current = scene
	} else {
		s.next = scene
		s.transitionFlag = true
	}
}

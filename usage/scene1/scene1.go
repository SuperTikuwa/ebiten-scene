/*
This package manages one of the scenes.
The methods of the structure are switched by 'sceneManager.GoTo'
and executed in the 'Game' structure of main.go.
*/
package scene1

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/supertikuwa/ebiten-scene/scene"
	"github.com/supertikuwa/usage/scene2"
)

type scene1 struct {
	color color.Color
}

func NewScene1() scene.Scene {
	ebiten.SetWindowTitle("Scene1")
	return &scene1{color: color.Black}
}

// Fill the background with black.
func (s scene1) Draw(screen *ebiten.Image) {
	screen.Fill(s.color)
}
func (s scene1) Update(sm *scene.SceneManager) error {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		sm.GoTo(scene2.NewScene2())
	}
	return nil
}

/*
This package manages the title scene.
The methods of the title structure are switched by 'sceneManager.GoTo'
and executed from the Game structure in main.go.
*/
package title

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/supertikuwa/ebiten-scene/scene"
	"github.com/supertikuwa/usage/scene1"
)

type title struct {
	color color.Color
}

func NewTitle() scene.Scene {
	ebiten.SetWindowTitle("Title")
	return &title{color: color.White}
}

func (t *title) Update(sceneManager *scene.SceneManager) error {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		sceneManager.GoTo(scene1.NewScene1())
	}
	return nil
}

// Fill the background with white.
func (t *title) Draw(screen *ebiten.Image) {
	screen.Fill(t.color)
}

/*
This package manages one of the scenes.
The methods of the structure are switched by 'sceneManager.GoTo'
and executed in the 'Game' structure of main.go.
*/
package scene2

import (
	"fmt"
	"image/color"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/supertikuwa/ebiten-scene/scene"
)

type scene2 struct {
	color color.Color
}

func NewScene2() scene.Scene {
	ebiten.SetWindowTitle("Scene2")
	return &scene2{color: color.RGBA{R: 255, G: 0, B: 0, A: 255}}
}

func (s *scene2) Update(sceneManager *scene.SceneManager) error {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		fmt.Println("Game Finished")
		os.Exit(0)
	}
	return nil
}

// Fill the background with red.
func (s *scene2) Draw(screen *ebiten.Image) {
	screen.Fill(s.color)
}

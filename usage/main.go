// main.go must have a struct likes 'Game'.
// Game struct that is an interface of 'ebiten.Game'.
// It struct must have a member variable of type '*scene.SceneManager'

package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/supertikuwa/ebiten-scene/scene"
	"github.com/supertikuwa/usage/title"
)

type Game struct {
	sceneManager *scene.SceneManager
}

func (g *Game) Update() error {
	if g.sceneManager == nil {
		g.sceneManager = &scene.SceneManager{}
		g.sceneManager.GoTo(title.NewTitle())
	}

	err := g.sceneManager.Update()
	return err
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.sceneManager.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	g := &Game{}
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

package scene_test

import (
	"image/color"
	"log"
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/supertikuwa/ebiten-scene/scene"
)

type Game struct {
	sceneManager *scene.SceneManager
}

func (g *Game) Update() error {
	if g.sceneManager == nil {
		g.sceneManager = &scene.SceneManager{}
		g.sceneManager.GoTo(&First{})
		return nil
	}

	g.sceneManager.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.sceneManager.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

// Scene interface.
type First struct {
}

func (f *First) Update(sceneManager *scene.SceneManager) error {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		sceneManager.GoTo(NewSecondScene())
		return nil
	}

	return nil
}

func (f *First) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
	ebitenutil.DebugPrint(screen, "Scene1")
}

type Second struct {
	message string
}

func (s *Second) Update(sceneManager *scene.SceneManager) error {

	return nil
}

func (s *Second) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)
	ebitenutil.DebugPrint(screen, s.message)
}

func NewSecondScene() scene.Scene {
	return &Second{message: "second"}
}

func Example() {
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}

func TestMain(m *testing.M) {
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}

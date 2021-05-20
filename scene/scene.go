package scene

import "github.com/hajimehoshi/ebiten/v2"

// Scene is an interface with Draw and Update.
type Scene interface {
	Update(sceneManager *SceneManager) error
	Draw(r *ebiten.Image)
}

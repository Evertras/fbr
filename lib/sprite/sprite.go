package sprite

import "github.com/hajimehoshi/ebiten"

// Sprite is some image that knows how to draw itself to a target
type Sprite interface {
	Draw(target *ebiten.Image, m ebiten.GeoM) error
}

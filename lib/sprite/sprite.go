package sprite

import (
	"time"

	"github.com/hajimehoshi/ebiten"
)

// Sprite is some image that knows how to draw itself to a target
type Sprite interface {
	Draw(target *ebiten.Image, m ebiten.GeoM)
	Update(delta time.Duration)
}

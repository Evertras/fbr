package sprite

import (
	"time"

	"github.com/hajimehoshi/ebiten"
)

// Sprite is some image that knows how to draw itself to a target
type Sprite interface {
	// Draws the sprite to the target using the given matrix for a transform
	Draw(target *ebiten.Image, m ebiten.GeoM)

	// Updates the sprite, advancing animations
	Update(delta time.Duration)

	// Returns true if the sprite is non-looping and has completed its animation
	Complete() bool
}

package sprite

import (
	"time"

	"github.com/hajimehoshi/ebiten"
)

// static is a static sprite with no animations
type static struct {
	img *ebiten.Image
}

// NewStatic creates a new Static sprite using the supplied image
func NewStatic(img *ebiten.Image) Sprite {
	return &static{
		img: img,
	}
}

// Draw draws the sprite to the supplied ebiten image
func (s *static) Draw(target *ebiten.Image, m ebiten.GeoM) {
	// DrawImage technically returns an error, but according to documentation
	// this is always nil in the version we're using so we don't care
	target.DrawImage(s.img, &ebiten.DrawImageOptions{
		GeoM: m,
	})
}

// Update is a no-op for a static sprite
func (s *static) Update(_ time.Duration) {}

package sprite

import "github.com/hajimehoshi/ebiten"

// Static is a static sprite with no animations
type Static struct {
	img *ebiten.Image
}

// NewStatic creates a new Static sprite using the supplied image
func NewStatic(img *ebiten.Image) Sprite {
	return &Static{
		img: img,
	}
}

// Draw draws the sprite to the supplied ebiten image
func (s *Static) Draw(target *ebiten.Image, m ebiten.GeoM) error {
	return target.DrawImage(s.img, &ebiten.DrawImageOptions{
		GeoM: m,
	})
}

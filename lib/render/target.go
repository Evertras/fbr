package render

import "github.com/hajimehoshi/ebiten"

// Target represents a target to render things to.  This
// generally targets Ebiten, but allows us to mock it
// out for unit testing.
type Target interface {
	DrawImage(img *ebiten.Image, opts *ebiten.DrawImageOptions) error
}

// Ensure an Ebiten Image fits the interface
var _ Target = &ebiten.Image{}

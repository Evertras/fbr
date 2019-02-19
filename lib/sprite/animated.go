package sprite

import (
	"image"
	"time"

	"github.com/hajimehoshi/ebiten"
)

type animated struct {
	sheet     *ebiten.Image
	frames    []image.Rectangle
	numFrames float64
	opts      AnimationOptions

	// Round down to get the actual frame
	currentFrame float64
}

// AnimationOptions can specify various options for animation of sprites
type AnimationOptions struct {
	FPS float64
}

// NewAnimated creates an animated sprite that works via sprite sheets
func NewAnimated(sheet *ebiten.Image, frames []image.Rectangle, opts AnimationOptions) Sprite {
	return &animated{
		sheet:     sheet,
		frames:    frames,
		numFrames: float64(len(frames)),
		opts:      opts,
	}
}

func (a *animated) Draw(target *ebiten.Image, m ebiten.GeoM) {
	target.DrawImage(a.sheet, &ebiten.DrawImageOptions{
		GeoM: m,

		// Note: This method is deprecated as of Ebiten 1.9 (unreleased as of typing this),
		// but we're using 1.8 until the new stuff is stable/released.  Update this at that point.
		SourceRect: &a.frames[int(a.currentFrame)],
	})
}

func (a *animated) Update(delta time.Duration) {
	a.currentFrame += delta.Seconds() * a.opts.FPS

	// Wrap around smoothly
	for a.currentFrame >= a.numFrames {
		a.currentFrame -= a.numFrames
	}
}

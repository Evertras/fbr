package components

import (
	"image"

	"github.com/Evertras/fbr/lib/ecs"
	"github.com/hajimehoshi/ebiten"
)

// SpriteAnimationOptions can specify various options for animation of sprites
type SpriteAnimationOptions struct {
	FPS   float64
	Loops bool
}

// Sprite defines a sprite to draw for the entity
type Sprite struct {
	ecs.BaseComponent

	Sheet     *ebiten.Image
	Frames    []image.Rectangle
	NumFrames float64
	Opts      SpriteAnimationOptions

	// For non-looping animations, set this when done
	Completed bool

	// Round down to get the actual frame
	CurrentFrame float64
}

// NewSpriteAnimated creates an animated sprite that works via sprite sheets
func NewSpriteAnimated(sheet *ebiten.Image, frames []image.Rectangle, opts SpriteAnimationOptions) *Sprite {
	return &Sprite{
		Sheet:     sheet,
		Frames:    frames,
		NumFrames: float64(len(frames)),
		Opts:      opts,
		Completed: false,
	}
}

// NewSpriteStatic creates a new static sprite using the supplied image
func NewSpriteStatic(img *ebiten.Image) *Sprite {
	width, height := img.Size()

	return &Sprite{
		Sheet: img,
		Frames: []image.Rectangle{
			image.Rect(0, 0, width, height),
		},
		CurrentFrame: 0,
		Opts: SpriteAnimationOptions{
			FPS:   0,
			Loops: true,
		},
		Completed: false,
	}
}

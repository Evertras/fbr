package systems

import (
	"time"

	"github.com/Evertras/fbr/lib/ecs"
	"github.com/Evertras/fbr/lib/game/components"
)

type spriteAnimations struct {
	componentSprite ecs.ComponentType
}

// NewSpriteAnimations creates a system that controls sprite animations
func NewSpriteAnimations(componentSprite ecs.ComponentType) ecs.System {
	return &spriteAnimations{
		componentSprite: componentSprite,
	}
}

func (s *spriteAnimations) ActOn(w *ecs.World, delta time.Duration) {
	// Note: we do care about animations even if Ebiten isn't drawing, since
	// a completed animation may trigger reaping of an entity, so don't check
	// IsDrawingSkipped here!

	sprites := w.GetComponents(s.componentSprite)

	for _, sData := range sprites {
		sprite := sData.(*components.Sprite)
		if !sprite.Completed {
			sprite.CurrentFrame += delta.Seconds() * sprite.Opts.FPS

			if sprite.Opts.Loops {
				// Wrap around smoothly
				for sprite.CurrentFrame >= sprite.NumFrames {
					sprite.CurrentFrame -= sprite.NumFrames
				}
			} else if sprite.CurrentFrame >= sprite.NumFrames {
				sprite.CurrentFrame = sprite.NumFrames - 0.01
				sprite.Completed = true
			}
		}
	}
}

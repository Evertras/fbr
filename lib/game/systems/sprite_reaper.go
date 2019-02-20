package systems

import (
	"time"

	"github.com/Evertras/fbr/lib/ecs"
	"github.com/Evertras/fbr/lib/game/components"
)

type spriteReaper struct {
	componentSprite ecs.ComponentType
}

// NewSpriteReaper creates a system that deletes entities once their animated sprite completes
func NewSpriteReaper(componentSprite ecs.ComponentType) ecs.System {
	return &spriteReaper{
		componentSprite: componentSprite,
	}
}

func (s *spriteReaper) ActOn(w *ecs.World, delta time.Duration) {
	for _, raw := range w.GetComponents(s.componentSprite) {
		sprite := raw.(*components.Sprite)

		if sprite.Completed {
			w.MarkEntityDeleted(sprite.GetOwner())
		}
	}
}

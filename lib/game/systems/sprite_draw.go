package systems

import (
	"github.com/Evertras/fbr/lib/ecs"
	"github.com/Evertras/fbr/lib/game/components"
	"github.com/hajimehoshi/ebiten"
)

type spriteDraw struct {
	componentSprite   ecs.ComponentType
	componentPosition ecs.ComponentType
}

// NewSpriteDraw returns a draw system that draws any entities with Sprite components to the screen
func NewSpriteDraw(componentSprite, componentPosition ecs.ComponentType) ecs.SystemDraw {
	return &spriteDraw{
		componentSprite:   componentSprite,
		componentPosition: componentPosition,
	}
}

func (s *spriteDraw) Draw(w *ecs.World, target *ebiten.Image) {
	// This whole function is bad but wins with simplicity until the ECS system handles multiple components better
	sprites := w.GetComponents(s.componentSprite)

	for _, sData := range sprites {
		e := sData.GetOwner()

		if pData, ok := w.GetComponent(e, s.componentPosition); ok {
			sprite := sData.(*components.Sprite)
			pos := pData.(*components.Position)

			m := ebiten.GeoM{}
			m.Translate(pos.X, pos.Y)

			target.DrawImage(sprite.Sheet, &ebiten.DrawImageOptions{
				GeoM: m,

				// Note: This method is deprecated as of Ebiten 1.9 (unreleased as of typing this),
				// but we're using 1.8 until the new stuff is stable/released.  Update this at that point.
				SourceRect: &sprite.Frames[int(sprite.CurrentFrame)],
			})
		}
	}
}

package systems

import (
	"time"

	"github.com/Evertras/fbr/lib/ecs"
	"github.com/Evertras/fbr/lib/game/components"
	"github.com/hajimehoshi/ebiten"
)

type inputLocal struct {
	inputLocalType ecs.ComponentType
	positionType   ecs.ComponentType
}

// NewInputLocal creates a system that allows entities to be controlled
// by the local environment directly.
func NewInputLocal(inputLocalType ecs.ComponentType, positionType ecs.ComponentType) ecs.System {
	return &inputLocal{
		inputLocalType: inputLocalType,
		positionType:   positionType,
	}
}

func (i *inputLocal) ActOn(w *ecs.World, delta time.Duration) {
	all := w.GetComponents(i.inputLocalType)

	// Would be a bit surprising to have more than one, but don't care for now
	for _, raw := range all {
		input := raw.(*components.InputLocal)

		e := input.GetOwner()

		rawPos, ok := w.GetComponent(e, i.positionType)

		if !ok {
			continue
		}

		pos := rawPos.(*components.Position)
		adj := delta.Seconds() * input.Speed

		if ebiten.IsKeyPressed(ebiten.KeyRight) {
			pos.X += adj
		}

		if ebiten.IsKeyPressed(ebiten.KeyLeft) {
			pos.X -= adj
		}

		if ebiten.IsKeyPressed(ebiten.KeyUp) {
			pos.Y -= adj
		}

		if ebiten.IsKeyPressed(ebiten.KeyDown) {
			pos.Y += adj
		}
	}
}

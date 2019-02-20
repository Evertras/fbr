package game

import (
	"github.com/Evertras/fbr/lib/asset"
	"github.com/Evertras/fbr/lib/ecs"
	"github.com/Evertras/fbr/lib/game/components"
	"github.com/pkg/errors"
)

// SpawnFire creates a fire visual effect at the given location
func (i *Instance) SpawnFire(x, y float64) (ecs.EntityID, error) {
	e := i.world.NewEntity()

	i.world.AddComponent(e, i.componentPosition, &components.Position{
		X: x,
		Y: y,
	})

	fireSheet, err := asset.LoadImageFromPath("http://localhost:8000/assets/fire.png")

	if err != nil {
		return 0, errors.Wrap(err, "could not load fire sprite sheet")
	}

	fireFrames, err := asset.LoadFramesFromPath("http://localhost:8000/assets/fire.frames")

	if err != nil {
		return 0, errors.Wrap(err, "could not load fire frames")
	}

	i.world.AddComponent(
		e,
		i.componentSprite,
		components.NewSpriteAnimated(
			fireSheet,
			fireFrames,
			components.SpriteAnimationOptions{
				FPS:   60,
				Loops: true,
			},
		),
	)

	return e, nil
}

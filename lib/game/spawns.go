package game

import (
	"github.com/Evertras/fbr/lib/asset"
	"github.com/Evertras/fbr/lib/ecs"
	"github.com/Evertras/fbr/lib/game/components"
	"github.com/pkg/errors"
)

// SpawnFire creates a fire visual effect at the given location
func (i *Instance) SpawnFire(x, y float64) (ecs.EntityID, error) {
	fireSheet, err := asset.LoadImageFromPath("assets/fire.png")

	if err != nil {
		return 0, errors.Wrap(err, "could not load fire sprite sheet")
	}

	fireFrames, err := asset.LoadFramesFromPath("assets/fire.frames")

	if err != nil {
		return 0, errors.Wrap(err, "could not load fire frames")
	}

	e := i.world.NewEntity()

	i.world.AddComponent(e, i.componentPosition, &components.Position{
		X: x,
		Y: y,
	})

	i.world.AddComponent(
		e,
		i.componentSprite,
		components.NewSpriteAnimated(
			fireSheet,
			fireFrames,
			components.SpriteAnimationOptions{
				FPS:   60,
				Loops: false,
			},
			i.layerObjects,
		),
	)

	return e, nil
}

// SpawnPlayer creates a new player entity at the specified position
func (i *Instance) SpawnPlayer(x, y float64) (ecs.EntityID, error) {
	sheet, err := asset.LoadImageFromPath("assets/wizard.png")

	if err != nil {
		return 0, errors.Wrap(err, "could not load sprite sheet")
	}

	frames, err := asset.LoadFramesFromPath("assets/wizard.idle.frames")

	if err != nil {
		return 0, errors.Wrap(err, "could not load frames")
	}

	e := i.world.NewEntity()

	opts := components.SpriteAnimationOptions{
		FPS:   10,
		Loops: true,
	}

	i.world.AddComponent(e, i.componentSprite, components.NewSpriteAnimated(sheet, frames, opts, i.layerObjects))
	i.world.AddComponent(e, i.componentPosition, &components.Position{X: x, Y: y})
	i.world.AddComponent(e, i.componentInputLocal, &components.InputLocal{Speed: 10})

	return e, nil
}

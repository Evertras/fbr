package game

import (
	"time"

	"github.com/Evertras/fbr/lib/ecs"
	"github.com/Evertras/fbr/lib/game/systems"
	"github.com/hajimehoshi/ebiten"
)

// Instance is a runnable game instance
type Instance struct {
	world *ecs.World

	componentSprite   ecs.ComponentType
	componentPosition ecs.ComponentType
}

// Step steps the game forward by the given delta
func (i *Instance) Step(delta time.Duration) {
	i.world.Step(delta)
}

// Draw draws the world to the given target
func (i *Instance) Draw(target *ebiten.Image) {
	i.world.Draw(target)
}

// NewClient creates a new Instance made for Clients
func NewClient() *Instance {
	i := &Instance{}

	i.world = ecs.NewWorld()

	i.initComponentTypes()

	// Updates
	i.world.RegisterSystem(systems.NewSpriteAnimations(i.componentSprite))
	i.world.RegisterSystem(systems.NewSpriteReaper(i.componentSprite))

	// Draws
	i.world.RegisterSystemDraw(systems.NewSpriteDraw(i.componentSprite, i.componentPosition))

	return i
}

// NewServer creates a new Instance made for Servers
func NewServer() *Instance {
	panic("not implemented yet")
}

// NumEntities returns the current number of entities in the world
func (i *Instance) NumEntities() uint32 {
	return i.world.NumEntities()
}

func (i *Instance) initComponentTypes() {
	i.componentSprite = i.world.NewComponent()
	i.componentPosition = i.world.NewComponent()
}

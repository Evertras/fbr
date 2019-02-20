package ecs

import (
	"time"

	"github.com/hajimehoshi/ebiten"
)

// System is something that knows how to act on a World
type System interface {
	ActOn(world *World, delta time.Duration)
}

// SystemDraw is something that knows how to draw to an Ebiten image target
type SystemDraw interface {
	Draw(world *World, target *ebiten.Image)
}

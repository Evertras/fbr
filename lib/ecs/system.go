package ecs

import "time"

// System is something that knows how to act on a World
type System interface {
	ActOn(world *World, delta time.Duration)
}

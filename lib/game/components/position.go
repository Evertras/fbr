package components

import "github.com/Evertras/fbr/lib/ecs"

// Position defines a place in the world that this entity exists
type Position struct {
	ecs.BaseComponent
	X float64
	Y float64
}

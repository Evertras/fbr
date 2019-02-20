package components

import "github.com/Evertras/fbr/lib/ecs"

// InputLocal indicates this entity takes input directly from the environment
type InputLocal struct {
	ecs.BaseComponent

	Speed float64
}

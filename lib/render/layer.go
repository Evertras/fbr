package render

import (
	"github.com/hajimehoshi/ebiten"
)

// DrawRequest contains information to draw something to Ebiten
type DrawRequest struct {
	Image   *ebiten.Image
	Options *ebiten.DrawImageOptions

	// An arbitrary sort order, almost always will be the Y coordinate of the sprite in the world.
	// Greater values will draw on top of lower values.
	Sort float64
}

// Layer represents a layer to draw to that psuedo-Z Buffers
type Layer struct {
	queue []DrawRequest
}

// NewLayer returns a new rendering Layer
func NewLayer() *Layer {
	return &Layer{
		// Give ourselves 100 capacity to start, append will automagically grow this later if needed
		// but we'd like to avoid growing when possible
		queue: make([]DrawRequest, 0, 100),
	}
}

// Queue adds a draw request to the pending draw queue
func (l *Layer) Queue(req DrawRequest) {
	// Keep the queue sorted
	for i := 0; i < len(l.queue); i++ {
		if l.queue[i].Sort > req.Sort {
			// Basically just a fancy insert
			l.queue = append(l.queue, DrawRequest{})
			copy(l.queue[i+1:], l.queue[i:])
			l.queue[i] = req
			return
		}
	}

	l.queue = append(l.queue, req)
}

// Draw goes through the queue and draws all requests, emptying the queue
func (l *Layer) Draw(target *ebiten.Image) {
	for _, req := range l.queue {
		// Eat error for now?  Feels bad
		target.DrawImage(req.Image, req.Options)
	}

	// Reuse storage, append will overwrite the old values
	l.queue = l.queue[:0]
}

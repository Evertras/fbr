package ecs

import (
	"sync"
	"sync/atomic"
	"time"
)

// EntityID identifies a unique Entity in the World
type EntityID uint32

// ComponentType is a simple tag to identify which component a system wants
type ComponentType uint32

// World represents a world filled with Entities/Components that can be acted on by Systems
type World struct {
	components         map[ComponentType][]Component
	componentsByEntity map[EntityID]map[ComponentType]Component
	systems            []System

	entityIDCounter      uint32
	componentTypeCounter uint32

	mu sync.RWMutex
}

// NewWorld will create a blank World ready to be populated by entities, components, and systems
func NewWorld() *World {
	return &World{
		components:         make(map[ComponentType][]Component),
		componentsByEntity: make(map[EntityID]map[ComponentType]Component),
		systems:            make([]System, 0),
	}
}

// NewEntity generates a new Entity that's added to the world
func (w *World) NewEntity() EntityID {
	id := EntityID(atomic.AddUint32(&w.entityIDCounter, 1))

	w.mu.Lock()
	w.componentsByEntity[id] = make(map[ComponentType]Component)
	w.mu.Unlock()

	return id
}

// NewComponent generates a new ComponentType for reference later
func (w *World) NewComponent() ComponentType {
	id := ComponentType(atomic.AddUint32(&w.componentTypeCounter, 1))

	w.mu.Lock()
	w.components[id] = make([]Component, 0)
	w.mu.Unlock()

	return id
}

// RegisterSystem adds the given system to the world.  Systems are
// run in the same order they're registered.
func (w *World) RegisterSystem(s System) {
	w.systems = append(w.systems, s)
}

// Step moves the world forward by the given time step
func (w *World) Step(delta time.Duration) {
	for _, s := range w.systems {
		s.ActOn(w, delta)
	}
}

// AddComponent adds a component to a given entity
func (w *World) AddComponent(e EntityID, c ComponentType, data Component) {
	w.mu.Lock()
	data.SetOwner(e)
	w.components[c] = append(w.components[c], data)
	w.componentsByEntity[e][c] = data
	w.mu.Unlock()
}

// GetComponent gets the component data for an entity
func (w *World) GetComponent(e EntityID, c ComponentType) (data Component, found bool) {
	w.mu.RLock()
	entity, ok := w.componentsByEntity[e]
	if !ok {
		found = false
		w.mu.RUnlock()
		return
	}

	data, found = entity[c]
	w.mu.RUnlock()

	return
}

// GetComponents gets all components of a given type
func (w *World) GetComponents(c ComponentType) []Component {
	return w.components[c]
}

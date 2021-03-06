package ecs

import (
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

	pendingDelete []EntityID

	entityIDCounter      uint32
	componentTypeCounter uint32

	numEntities uint32
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

	w.componentsByEntity[id] = make(map[ComponentType]Component)

	atomic.AddUint32(&w.numEntities, 1)

	return id
}

// NumEntities returns the current number of active entities in the world
func (w *World) NumEntities() uint32 {
	return w.numEntities
}

// MarkEntityDeleted marks the entity for a pending delete during the next cleanup
func (w *World) MarkEntityDeleted(id EntityID) {
	w.pendingDelete = append(w.pendingDelete, id)
}

// NewComponent generates a new ComponentType for reference later
func (w *World) NewComponent() ComponentType {
	id := ComponentType(atomic.AddUint32(&w.componentTypeCounter, 1))

	w.components[id] = make([]Component, 10000)[:0]

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

	// This is bad, optimize later if it turns ou to be demonstrably slow
	for _, e := range w.pendingDelete {
		for cType := range w.componentsByEntity[e] {
			for i, c := range w.components[cType] {
				if c.GetOwner() == e {
					w.components[cType][i] = w.components[cType][len(w.components[cType])-1]
					w.components[cType] = w.components[cType][:len(w.components[cType])-1]
					break
				}
			}
		}

		delete(w.componentsByEntity, e)
		w.numEntities--
	}

	w.pendingDelete = w.pendingDelete[:0]
}

// AddComponent adds a component to a given entity; this will override an existing component of the same type
func (w *World) AddComponent(e EntityID, c ComponentType, data Component) {
	data.SetOwner(e)
	w.components[c] = append(w.components[c], data)
	w.componentsByEntity[e][c] = data
}

// GetComponent gets the component data for an entity
func (w *World) GetComponent(e EntityID, c ComponentType) (data Component, found bool) {
	entity, ok := w.componentsByEntity[e]
	if !ok {
		found = false
		return
	}

	data, found = entity[c]

	return
}

// GetComponents gets all components of a given type
func (w *World) GetComponents(c ComponentType) []Component {
	return w.components[c]
}

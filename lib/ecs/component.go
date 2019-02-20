package ecs

// Component is some blob of data that knows its owner
type Component interface {
	// GetOwner returns the owner of the Component
	GetOwner() EntityID

	// SetOwner sets the owner of the component
	SetOwner(id EntityID)
}

// BaseComponent provides an easy way to create components
// that can be used in this ECS system.
//
// Add BaseComponent to the struct like so:
// type MyComponent struct {
//     BaseComponent
//
//     MyData int
// }
type BaseComponent struct {
	owner EntityID
}

// GetOwner gets the owner of this component's data
func (c *BaseComponent) GetOwner() EntityID {
	return c.owner
}

// SetOwner sets the owner of this component's data
func (c *BaseComponent) SetOwner(id EntityID) {
	c.owner = id
}

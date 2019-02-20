package ecs

import (
	"testing"
	"time"
)

func TestCreatedEntitiesHaveUniqueIDs(t *testing.T) {
	w := NewWorld()

	total := 1000

	entities := make([]EntityID, total)

	for i := 0; i < total; i++ {
		entities[i] = w.NewEntity()
	}

	for i := 0; i < total-1; i++ {
		for j := i + 1; j < total; j++ {
			if entities[i] == entities[j] {
				t.Fatal("Found duplicate entity IDs")
			}
		}
	}
}

func TestGetComponentsCorrectlyReturnsMissing(t *testing.T) {
	w := NewWorld()

	e := w.NewEntity()
	c := w.NewComponent()

	_, ok := w.GetComponent(e, c)

	if ok {
		t.Fatal("Got some unexpected data back")
	}
}

func TestGetComponent(t *testing.T) {
	w := NewWorld()

	e := w.NewEntity()
	c := w.NewComponent()

	w.AddComponent(e, c, &SampleComponent{
		T: time.Second,
	})

	raw, ok := w.GetComponent(e, c)

	if !ok {
		t.Fatal("Did not get component back")
	}

	data := raw.(*SampleComponent)

	// Quick sanity check to make sure it's the same thing
	if data.T != time.Second {
		t.Fatalf("Expected T to be %v, but was %v", time.Second, data.T)
	}
}

func TestGetComponents(t *testing.T) {
	w := NewWorld()

	numEntities := 10

	c := w.NewComponent()

	for i := 0; i < numEntities; i++ {
		e := w.NewEntity()

		w.AddComponent(e, c, &SampleComponent{
			T: time.Second,
		})
	}

	components := w.GetComponents(c)

	if len(components) != numEntities {
		t.Fatalf("Expected %d components, but got %d", numEntities, len(components))
	}
}

func TestSimplestPossibleWorks(t *testing.T) {
	w := NewWorld()

	e := w.NewEntity()
	c := w.NewComponent()

	w.RegisterSystem(NewSampleSystem(c))

	w.AddComponent(e, c, &SampleComponent{})

	w.Step(time.Second)

	raw, ok := w.GetComponent(e, c)

	if !ok {
		t.Fatal("Expected to find component data, but found nothing")
	}

	updated := raw.(*SampleComponent)

	if updated.T != time.Second {
		t.Fatalf("Expected to have duration of %v, but got %v", time.Second, updated.T)
	}
}

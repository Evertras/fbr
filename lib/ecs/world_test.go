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

func TestSimplestPossibleWorks(t *testing.T) {
	w := NewWorld()

	e := w.NewEntity()
	c := w.NewComponent()

	w.RegisterSystem(&SampleSystem{})

	w.AddComponent(e, c, &SampleComponent{0, 0, 0})

	w.Step(time.Second)

	raw, ok := w.GetComponent(e, c)

	if !ok {
		t.Fatal("Expected to find component data, but found nothing")
	}

	updated := raw.(*SampleComponent)

	if updated.T != time.Second {
		t.Fatalf("Expected to have duration of 1 second, but got %v", updated.T)
	}
}

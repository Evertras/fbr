package ecs

import (
	"fmt"
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

func TestDeleteEntity(t *testing.T) {
	w := NewWorld()

	e1 := w.NewEntity()
	e2 := w.NewEntity()
	c := w.NewComponent()

	w.AddComponent(e1, c, &SampleComponent{})
	w.AddComponent(e2, c, &SampleComponent{})

	// Quick sanity check
	if len(w.GetComponents(c)) != 2 {
		t.Fatal("Did not actually add the component in the first place")
	}

	w.MarkEntityDeleted(e1)

	// Shouldn't be deleted yet
	if len(w.GetComponents(c)) != 2 {
		t.Fatal("Incorrectly deleted entity, which could cause nasty race conditions")
	}

	// Do a full step, even with no systems
	w.Step(time.Second)

	// Now it should be gone
	after := w.GetComponents(c)
	if len(after) != 1 {
		t.Fatalf("Expected to have 1 entity but got %d", len(after))
	}

	// The only remaining ID should be e2's
	if after[0].GetOwner() != e2 {
		t.Fatal("Deleted the wrong entity!")
	}
}

func Benchmark10kEntitiesSingleSimpleSystem(b *testing.B) {
	w := NewWorld()

	numEntities := 10000

	c := w.NewComponent()

	for i := 0; i < numEntities; i++ {
		e := w.NewEntity()

		w.AddComponent(e, c, &SampleComponent{})
	}

	w.RegisterSystem(NewSampleSystem(c))

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		w.Step(time.Nanosecond)
	}
}

func BenchmarkCreateAndDeleteSingleEntity(b *testing.B) {
	benchmarks := []struct {
		NumEntities   int
		NumComponents int
	}{
		{
			NumEntities:   10,
			NumComponents: 1,
		},
		{
			NumEntities:   1,
			NumComponents: 10,
		},
		{
			NumEntities:   100,
			NumComponents: 10,
		},
		{
			NumEntities:   10,
			NumComponents: 100,
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(fmt.Sprintf("%d entities %d components", benchmark.NumEntities, benchmark.NumComponents), func(b *testing.B) {
			w := NewWorld()

			components := make([]ComponentType, benchmark.NumComponents)

			for i := 0; i < benchmark.NumComponents; i++ {
				components[i] = w.NewComponent()
			}

			for i := 0; i < benchmark.NumEntities; i++ {
				e := w.NewEntity()

				for j := 0; j < benchmark.NumComponents; j++ {
					w.AddComponent(e, components[j], &SampleComponent{})
				}
			}

			fakeData := &SampleComponent{}

			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				e := w.NewEntity()

				for j := 0; j < benchmark.NumComponents; j++ {
					w.AddComponent(e, components[j], fakeData)
				}

				w.MarkEntityDeleted(e)
				w.Step(time.Nanosecond)
			}
		})
	}
}

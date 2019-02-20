package ecs

import (
	"time"
)

type SampleComponent struct {
	BaseComponent

	T time.Duration
}

type SampleSystem struct {
	scType ComponentType
}

func NewSampleSystem(scType ComponentType) *SampleSystem {
	return &SampleSystem{
		scType: scType,
	}
}

func (s *SampleSystem) ActOn(w *World, delta time.Duration) {
	for _, raw := range w.GetComponents(s.scType) {
		sc := raw.(*SampleComponent)

		sc.T += delta
	}
}

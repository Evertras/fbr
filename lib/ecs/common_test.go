package ecs

import "time"

type SampleComponent struct {
	X int
	Y int
	T time.Duration
}

type SampleSystem struct{}

func (s *SampleSystem) ActOn(w *World, delta time.Duration) {
}

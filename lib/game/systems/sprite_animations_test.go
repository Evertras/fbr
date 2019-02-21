package systems

import (
	"image"
	"math"
	"testing"
	"time"

	"github.com/Evertras/fbr/lib/ecs"

	"github.com/Evertras/fbr/lib/game/components"
)

func TestSpriteAnimatedLoops(t *testing.T) {
	frames := []image.Rectangle{
		image.Rect(0, 0, 1, 1),
		image.Rect(0, 0, 1, 1),
		image.Rect(0, 0, 1, 1),
		image.Rect(0, 0, 1, 1),
	}

	opts := components.SpriteAnimationOptions{
		FPS:   1,
		Loops: true,
	}

	sprite := components.NewSpriteAnimated(nil, frames, opts, nil)

	if sprite.CurrentFrame != 0 {
		t.Fatal("Expected to start at frame 0")
	}

	if sprite.NumFrames != 4 {
		t.Fatalf("Expected 4 frames, but got %f.1", sprite.NumFrames)
	}

	w := ecs.NewWorld()

	e := w.NewEntity()
	spriteType := w.NewComponent()
	w.AddComponent(e, spriteType, sprite)

	animationSystem := NewSpriteAnimations(spriteType)

	w.RegisterSystem(animationSystem)

	w.Step(time.Second)

	if math.Abs(sprite.CurrentFrame-1) > epsilon {
		t.Fatalf("Expected frame 1, but got %.1f", sprite.CurrentFrame)
	}

	w.Step(time.Second * 3)

	if math.Abs(sprite.CurrentFrame) > epsilon {
		t.Fatalf("Expected frame to wrap back around to 0, but at frame %.1f", sprite.CurrentFrame)
	}
}

func TestSpriteAnimatedComplets(t *testing.T) {
	frames := []image.Rectangle{
		image.Rect(0, 0, 1, 1),
		image.Rect(0, 0, 1, 1),
		image.Rect(0, 0, 1, 1),
		image.Rect(0, 0, 1, 1),
	}

	opts := components.SpriteAnimationOptions{
		FPS:   1,
		Loops: false,
	}

	sprite := components.NewSpriteAnimated(nil, frames, opts, nil)

	if sprite.CurrentFrame != 0 {
		t.Fatal("Expected to start at frame 0")
	}

	if sprite.NumFrames != 4 {
		t.Fatalf("Expected 4 frames, but got %f.1", sprite.NumFrames)
	}

	w := ecs.NewWorld()

	e := w.NewEntity()
	spriteType := w.NewComponent()
	w.AddComponent(e, spriteType, sprite)

	animationSystem := NewSpriteAnimations(spriteType)

	w.RegisterSystem(animationSystem)

	w.Step(time.Second)

	if math.Abs(sprite.CurrentFrame-1) > epsilon {
		t.Fatalf("Expected frame 1, but got %.1f", sprite.CurrentFrame)
	}

	w.Step(time.Second * 3)

	if !sprite.Completed {
		t.Fatal("Expected animation to complete, but it didn't")
	}
}

package sprite

import (
	"image"
	"math"
	"testing"
	"time"
)

func TestSpriteAnimatedUpdates(t *testing.T) {
	frames := []image.Rectangle{
		image.Rect(0, 0, 1, 1),
		image.Rect(0, 0, 1, 1),
		image.Rect(0, 0, 1, 1),
		image.Rect(0, 0, 1, 1),
	}

	opts := AnimationOptions{
		FPS: 1,
	}

	sprite := NewAnimated(nil, frames, opts).(*animated)

	if sprite.currentFrame != 0 {
		t.Fatal("Expected to start at frame 0")
	}

	if sprite.numFrames != 4 {
		t.Fatalf("Expected 4 frames, but got %f.1", sprite.numFrames)
	}

	sprite.Update(time.Second)

	if math.Abs(sprite.currentFrame-1) > epsilon {
		t.Fatalf("Expected frame 1, but got %.1f", sprite.currentFrame)
	}

	sprite.Update(time.Second * 3)

	if math.Abs(sprite.currentFrame) > epsilon {
		t.Fatalf("Expected frame to wrap back around to 0, but at frame %.1f", sprite.currentFrame)
	}
}

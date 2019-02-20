package main

import (
	"fmt"
	"math"
	"time"

	"github.com/Evertras/fbr/lib/game"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var lastFrame time.Time
var instance *game.Instance

func update(screen *ebiten.Image) error {
	// Frame timing
	now := time.Now()
	delta := now.Sub(lastFrame)

	defer func() { lastFrame = now }()

	// Input processing
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()

		instance.SpawnFire(float64(x), float64(y))
	}

	// Regular updates
	instance.Step(delta)

	// Drawing... if we should
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	instance.Draw(screen)

	ebitenutil.DebugPrint(
		screen,
		fmt.Sprintf("TPS: %.0f FPS: %.0f Entity Count: %d Delta: %.1f ms",
			ebiten.CurrentTPS(),
			ebiten.CurrentFPS(),
			instance.NumEntities(),
			delta.Seconds()*1000.0))

	return nil
}

func main() {
	var err error

	instance = game.NewClient()

	// <sandbox>
	instance.SpawnFire(0, 0)
	_, err = instance.SpawnPlayer(50, 50)
	if err != nil {
		panic(err)
	}
	// </sandbox

	w, h := ebiten.ScreenSizeInFullscreen()

	ebiten.SetFullscreen(true)

	lastFrame = time.Now()

	// Make it square
	size := int(math.Min(float64(w), float64(h)))

	if err = ebiten.Run(update, size, size, 1.0, "海老天 Playground"); err != nil {
		panic(err)
	}
}

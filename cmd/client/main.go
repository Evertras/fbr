// +build js,wasm

package main

import (
	"fmt"
	"time"

	"github.com/Evertras/fbr/lib/game"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var lastFrame time.Time
var instance *game.Instance

func update(screen *ebiten.Image) error {
	now := time.Now()
	defer func() { lastFrame = now }()

	delta := now.Sub(lastFrame)

	instance.Step(delta)

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	instance.Draw(screen)

	ebitenutil.DebugPrint(
		screen,
		fmt.Sprintf("TPS: %.0f FPS: %.0f Delta: %.1f ms",
			ebiten.CurrentTPS(),
			ebiten.CurrentFPS(),
			delta.Seconds()*1000.0))

	return nil
}

func main() {
	var err error

	instance = game.NewClient()

	// <sandbox>
	instance.SpawnFire(50, 50)
	// </sandbox

	w, h := ebiten.ScreenSizeInFullscreen()

	ebiten.SetFullscreen(true)

	lastFrame = time.Now()

	// TODO: handle resize ratio changes
	if err = ebiten.Run(update, w, h, 1.0, "Fantasy Battle Royale (Powered by 海老天)"); err != nil {
		panic(err)
	}
}

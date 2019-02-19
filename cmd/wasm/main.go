// +build js,wasm

package main

import (
	"fmt"
	"time"

	"github.com/Evertras/fbr/lib/sprite"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var wizSprite sprite.Sprite
var fireSprite sprite.Sprite
var lastFrame time.Time

// Temporary sandbox
func loadImages() {
	var err error

	wizSprite, err = sprite.AnimatedFromPath("assets/wizard.png", "assets/wizard.idle.frames", sprite.AnimationOptions{
		FPS:   10,
		Loops: true,
	})

	if err != nil {
		panic(err)
	}

	fireSprite, err = sprite.AnimatedFromPath("assets/fire.png", "assets/fire.frames", sprite.AnimationOptions{
		FPS:   60,
		Loops: true,
	})

	if err != nil {
		panic(err)
	}
}

func update(screen *ebiten.Image) error {
	now := time.Now()
	defer func() { lastFrame = now }()

	delta := now.Sub(lastFrame)
	wizSprite.Update(delta)
	fireSprite.Update(delta)

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	viewMatrix := ebiten.GeoM{}

	viewMatrix.Translate(50, 50)
	//viewMatrix.Invert()

	wizSprite.Draw(screen, viewMatrix)

	for x := 0; x < 50; x++ {
		for y := 0; y < 50; y++ {
			fireMatrix := ebiten.GeoM{}

			fireMatrix.Translate(float64(x*32), float64(y*32+32))

			fireSprite.Draw(screen, fireMatrix)
		}
	}

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

	loadImages()

	w, h := ebiten.ScreenSizeInFullscreen()

	ebiten.SetFullscreen(true)

	lastFrame = time.Now()

	// TODO: handle resize ratio changes
	if err = ebiten.Run(update, w, h, 1.0, "海老天 - Fantasy Battle Royale"); err != nil {
		panic(err)
	}
}

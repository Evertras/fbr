// +build js,wasm

package main

import (
	"fmt"

	"github.com/Evertras/fbr/lib/sprite"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var wizImage sprite.Sprite

func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	viewMatrix := ebiten.GeoM{}

	viewMatrix.Translate(10, 10)

	wizImage.Draw(screen, viewMatrix)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %.0f", ebiten.CurrentFPS()))

	return nil
}

func main() {
	var err error

	ebiten.SetFullscreen(true)

	wizImage, err = sprite.StaticFromPath("assets/wizard.png")

	if err != nil {
		panic(err)
	}

	if err := ebiten.Run(update, 500, 500, 1.0, "海老天"); err != nil {
		panic(err)
	}
}

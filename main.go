package main

import (
	_ "embed"
	"github.com/elgopher/pi"
)

//go:embed sprite-sheet.png
var spriteSheetPNG []byte

func main() {
	pi.SetScreenSize(128, 128)
	pi.Palette = pi.DecodePalette(spriteSheetPNG)
	spriteSheet := pi.DecodeCanvas(spriteSheetPNG)
	piSprite := pi.SpriteFrom(spriteSheet, 0, 0, 8, 8)

	pi.Update = func() {
		// Called every frame.
		//
		// This function should handle user input, performs calculations, and updates
		// the game state. Typically, it does not draw anything on the screen.
	}
	pi.Draw = func() {
		// Called (almost) every frame.
		//
		// Write here the code drawing on the screen.
		pi.Cls()
		pi.Spr(piSprite, 0, 0)
	}

	run()
}

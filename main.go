package main

import (
	"embed"

	"github.com/elgopher/pi"

	"github.com/elgopher/pi-template/game"
)

//go:embed sprite-sheet.png custom-font.png
var resources embed.FS

func main() {
	pi.Resources = resources
	pi.Update = game.Update
	pi.Draw = game.Draw

	run()
}

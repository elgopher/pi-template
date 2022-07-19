package main

import (
	"embed"
	"fmt"

	"github.com/elgopher/pi"
	"github.com/elgopher/pi-template/game"
)

//go:embed sprite-sheet.png
var resources embed.FS

func main() {
	pi.Resources = resources
	pi.Update = game.Update
	pi.Draw = game.Draw

	if err := pi.Run(); err != nil {
		fmt.Printf("I'm sorry. Pi game cannot be run :( The error is: %s\n", err.Error())
	}
}

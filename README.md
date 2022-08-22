# Pi game template 

This is a template for writing new games powered by [Pi](https://github.com/elgopher/pi).

## What's in the box? <img src="sprite-sheet.png" align="right" />

* sprite-sheet file with 16 colors palette inspired by [Pico-8](https://www.lexaloffle.com/pico-8.php).
* the sprite-sheet file has one sprite with number 0
* project is a Go module importing all the necessary dependencies
* main.go file contains code running the game
* game code is inside game package

## How to use it?

* [Go 1.18+](https://go.dev/dl/) is required
* Install [Ebitengine dependencies](https://ebiten.org/documents/install.html)
* ```go run main.go```

## Howtos

### How to rename the module?

* this template has a module named `github.com/elgopher/pi-template`
* edit the go.mod and replace the module name with your own, e.g. `module github.com/you/name`
* do the same with import in main.go

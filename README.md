# Pi Game Template

This is a template for creating new games powered by [Pi](https://github.com/elgopher/pi).

## What's in the box? <img src="sprite-sheet.png" align="right" />

* A sprite sheet with a 32-color palette from [Picotron][picotron-faq].
* The sprite sheet includes one 8×8 sprite.
* A Go module importing all necessary dependencies.
* A `main.go` file containing code to run the game.
* Two variants of code: development (with dev tools) and release (without dev tools).

## How to use it

* [Go 1.24+](https://go.dev/dl/) is required.
* Use any Go editor, such as [Visual Studio Code](https://code.visualstudio.com/) 
  or [GoLand](https://www.jetbrains.com/go/).
* On Linux or macOS, install additional dependencies:
  * [Linux instructions](https://github.com/elgopher/pi/blob/master/docs/install-linux.md)
  * [macOS instructions](https://github.com/elgopher/pi/blob/master/docs/install-macos.md)
* Run the game with:

  go run .

  > Do not run `main.go` directly, because Go will not compile the other `main_*` files.

## How-to guides

### How to rename the module

* The template module is named `github.com/elgopher/pi-template`.
* Edit `go.mod` and replace the module name with your own, e.g. `module github.com/you/name`.
* Update the import path in `main.go` accordingly.

### How to edit PNG files

* Use a pixel art editor that supports indexed color mode, such as [Aseprite][aseprite] or [LibreSprite][libresprite].
* The palette for the game is stored inside `sprite-sheet.png`.

### How to create a release build (without dev tools)

```
go build -tags release .
```

## Attributions

* The palette used in `sprite-sheet.png` is the [original Picotron palette][picotron-faq] created by Zep.

[aseprite]: https://www.aseprite.org/
[libresprite]: https://libresprite.github.io/
[picotron-faq]: https://www.lexaloffle.com/picotron.php?page=faq

---

Skopiuj od samego początku do końca – **to nie ma żadnych bloków kodowych** ani znaczników backtickowych, które mogłyby przeszkodzić w kopiowaniu.

Jakby nadal coś nie działało – powiedz, dam Ci go jeszcze prościej (np. całkiem surowy tekst bez formatowania nagłówków).

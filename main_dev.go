//go:build !release

package main

import (
	"github.com/elgopher/pi/piebiten"
	"github.com/elgopher/pi/piscope"
)

func run() {
	piscope.Start()
	piebiten.Run()
}

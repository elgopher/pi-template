//go:build !release

package main

import (
	"github.com/elgopher/pi/devtools"
	"github.com/elgopher/pi/ebitengine"
)

func run() {
	devtools.MustRun(ebitengine.Run)
}

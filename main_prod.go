//go:build release

package main

import (
	"fmt"

	"github.com/elgopher/pi/ebitengine"
)

// this function is executed only for release build
func run() {
	if err := ebitengine.Run(); err != nil {
		fmt.Printf("I'm sorry. Pi game cannot be run :( The error is: %s\n", err.Error())
	}
}

//go:build release

package main

import (
	"fmt"
	"github.com/elgopher/pi/piebiten"
)

// this function is executed only for release build
func run() {
	if err := piebiten.RunOrErr(); err != nil {
		fmt.Printf("I'm sorry. Pi game cannot be run :( The error is: %s\n", err.Error())
	}
}

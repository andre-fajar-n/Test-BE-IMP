// Package Main
package main

import (
	"test-be-IMP/cmd/server/cmd"
)

// main function
func main() {
	if err := cmd.Start(); nil != err {
		panic(err)
	}
}

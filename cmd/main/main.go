// Package main is the entry point for the application.
// Users typically do not need to modify this file.
package main

import (
	"fmt"
	"os"

	"github.com/mrrizkin/boot/system"
)

// main is the entry point of the application.
// It initializes and runs the system, handling any errors that occur.
//
// Note to users: This file serves as the standard entry point for the application.
// In most cases, you should not need to modify this file. Customizations and
// additional functionality should be implemented in other packages, particularly
// within the 'system' package or your own application-specific packages.
func main() {
	// Run the system and handle any errors
	if err := system.Run(); err != nil {
		// If an error occurs, print it to stderr and exit with a non-zero status
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	// If system.Run() completes without error, the program will exit normally
}

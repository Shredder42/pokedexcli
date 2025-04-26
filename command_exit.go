package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *config, args ...string) error {
	defer os.Exit(0)
	fmt.Printf("Closing the Pokedex... Goodbye!\n")
	return nil
}

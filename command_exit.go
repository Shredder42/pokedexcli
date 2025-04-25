package main

import (
	"fmt"
	"os"
)

func commandExit(config *Config) error {
	defer os.Exit(0)
	fmt.Printf("Closing the Pokedex... Goodbye!\n")
	return nil
}

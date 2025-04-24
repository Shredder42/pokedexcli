package main

import (
	"fmt"
	"os"
)

func commandExit() error {
	defer os.Exit(0)
	fmt.Printf("Closing the Pokedex... Goodbye!\n")
	return nil
}

package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.Pokedex) == 0 {
		return errors.New("you haven't caught any Pokemon")
	}

	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.Pokedex {
		fmt.Printf("  - %s\n", pokemon.Name)
	}

	return nil
}

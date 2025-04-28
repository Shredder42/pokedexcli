package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must specify pokemon to inspect")
	}

	pokemonName := args[0]

	value, exists := cfg.Pokedex[pokemonName]
	if exists {
		fmt.Printf("Name: %s\n", value.Name)
		fmt.Printf("Height: %d\n", value.Height)
		fmt.Printf("Weight: %d\n", value.Weight)
		fmt.Println("Stats:")
		for _, stat := range value.Stats {
			fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, typeInfo := range value.Types {
			fmt.Printf("  -%s\n", typeInfo.Type.Name)
		}
	} else {
		fmt.Printf("you haven't caught %s!\n", pokemonName)
	}

	return nil
}

package main

import (
	"errors"
	// "encoding/json"
	"fmt"
	// "io"
	// "net/http"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	area := args[0]

	location, err := cfg.pokeapiClient.GetLocation(area)
	if err != nil {
		return err
	}

	for _, encounter := range location.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}

	return nil
}

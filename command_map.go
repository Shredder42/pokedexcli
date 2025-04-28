package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config, args ...string) error {
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.NextLocationAreaURL)
	if err != nil {
		return err
	}

	if locationsResp.Next != nil {
		cfg.NextLocationAreaURL = locationsResp.Next
	}

	if locationsResp.Previous != nil {
		cfg.PrevLocationAreaURL = locationsResp.Previous
	}

	for _, location := range locationsResp.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.PrevLocationAreaURL == nil {
		return errors.New("you're on the first page")
	}

	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.PrevLocationAreaURL)
	if err != nil {
		return err
	}

	cfg.NextLocationAreaURL = locationsResp.Next
	cfg.PrevLocationAreaURL = locationsResp.Previous

	for _, location := range locationsResp.Results {
		fmt.Println(location.Name)
	}

	return nil
}

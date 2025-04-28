package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	// catch ranges from 20 to 185 or so
	pokemon_name := args[0]

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon_name)

	throwSkill := rand.Intn(186)

	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemon_name)
	if err != nil {
		return err
	}

	pokemonExperience := pokemon.BaseExperience

	if throwSkill >= pokemonExperience {
		fmt.Printf("%s was caught!\n", pokemon_name)
		cfg.Pokedex[pokemon_name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemon_name)
	}

	return nil
}

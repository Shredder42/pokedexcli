package main

import (
	// "errors"
	// "encoding/json"
	"fmt"
	// "io"
	// "net/http"
)

func commandExplore(cfg *config, args ...string) error {
	fmt.Println("Ran commandExplore")
	// url := "https://pokeapi.co/api/v2/location-area/canalave-city-area"

	if len(args) != 1 {
		return fmt.Errorf(("you can only explore one area"))
	}

	area := args[0]

	location, err := cfg.pokeapiClient.GetLocation(area)
	if err != nil {
		return err
	}

	for _, encounter := range location.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}

	// this will take the command city (so need to update the signature)
	// - logic to make sure only one this is entered
	// - run a request to get the raw data in another function (try only looking at the location list and types locations and command mapstuff)
	// 	- call it location_get.go
	// - print out the list of pokemon
	// - i think you have the pieces to follow through on this one

	// if pageURL != nil {
	// 	url = *pageURL
	// }

	// if val, ok := c.cache.Get(url); ok {
	// locationsResp := RespShallowLocations{}
	// 	err := json.Unmarshal(val, &locationsResp)
	// 	if err != nil {
	// 		return RespShallowLocations{}, nil
	// 	}
	// 	return locationsResp, nil
	// }

	// req, err := http.NewRequest("GET", url, nil)
	// if err != nil {
	// 	return nil, err
	// }

	// resp, err := http.Get(url)
	// if err != nil {
	// 	return fmt.Errorf("error creating request: %w", err)
	// }

	// defer resp.Body.Close()

	// dat, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	return fmt.Errorf("error reading response: %w", err)
	// }

	// if err := json.Unmarshal(dat, &areaResp); err != nil {
	// 	return fmt.Errorf("error unmarshaling data: %w", err)
	// }
	// // fmt.Println(string(dat))

	// for _, encounter := range areaResp.PokemonEncounters {
	// 	fmt.Println(encounter.Pokemon.Name)
	// }

	// c.cache.Add(url, dat)

	// }

	return nil
}

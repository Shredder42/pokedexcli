package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMap(config *Config) error {

	url := "https://pokeapi.co/api/v2/location-area/"

	if config.NextLocationAreaURL != "" {
		url = config.NextLocationAreaURL
	}

	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	defer res.Body.Close()

	var locationsResp struct {
		Count    int     `json:"count"`
		Next     *string `json:"next"`
		Previous *string `json:"previous"`
		Results  []struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"results"`
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error reading response: %w", err)
	}

	if err := json.Unmarshal(data, &locationsResp); err != nil {
		return fmt.Errorf("error unmarshaling data: %w", err)
	}

	config.NextLocationAreaURL = ""
	if locationsResp.Next != nil {
		config.NextLocationAreaURL = *locationsResp.Next
	}

	config.PrevLocationAreaURL = ""
	if locationsResp.Previous != nil {
		config.PrevLocationAreaURL = *locationsResp.Previous
	}

	for _, location := range locationsResp.Results {
		fmt.Println(location.Name)
	}

	return nil
}

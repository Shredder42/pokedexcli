package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/shredder42/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient       pokeapi.Client
	NextLocationAreaURL *string
	PrevLocationAreaURL *string
	Pokedex             map[string]pokeapi.Pokemon
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		cleanedText := cleanInput(text)
		if len(cleanedText) == 0 {
			continue
		}

		commandName := cleanedText[0]
		args := []string{}
		if len(cleanedText) > 1 {
			args = cleanedText[1:]
		}

		command, ok := getCommands()[commandName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		} else {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		}
	}
}

func cleanInput(text string) []string {
	editedWords := []string{}
	fieldedText := strings.Fields(text)

	for _, word := range fieldedText {
		editedWords = append(editedWords, strings.ToLower(word))
	}
	return editedWords
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help mesage",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous 20 locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Displays pokemon in area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Throw Pokeball to catch Pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Displays information about caught Pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Displays all the Pokemon you have caught",
			callback:    commandPokedex,
		},
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		cleanedText := cleanInput(text)
		if len(cleanedText) == 0 {
			continue
		}

		command, ok := getCommands()[cleanedText[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		} else {
			err := command.callback()
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
	callback    func() error
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
	}
}

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

		fmt.Printf("Your command was: %s\n", cleanedText[0])
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

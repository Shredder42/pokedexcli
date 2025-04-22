package main

import (
	"strings"
)

func cleanInput(text string) []string {
	editedWords := []string{}
	fieldedText := strings.Fields(text)

	for _, word := range fieldedText {
		editedWords = append(editedWords, strings.ToLower(word))
	}
	return editedWords
}

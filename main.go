package main

import (
	"fmt"
	"strings"
)

func main() {

	text := "Hello, World!"
	fmt.Print(cleanInput(text))

}

func cleanInput(text string) []string {
	lwrString := strings.ToLower(text)
	words := strings.Fields(strings.ToLower(lwrString))

	return words

}

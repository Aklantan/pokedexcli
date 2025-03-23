package main

import (
	"fmt"
	"os"
)

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	err := fmt.Errorf("program exited")
	os.Exit(0)
	return err
}

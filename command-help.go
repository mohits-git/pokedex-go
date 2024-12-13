package main

import "fmt"

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the PokedexGo!")
	fmt.Println("Usage: ")
	fmt.Println("help: Show the list of available commands")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

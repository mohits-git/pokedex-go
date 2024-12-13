package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	cleanedInput := []string{}
	for _, word := range strings.Fields(text) {
		cleanedInput = append(cleanedInput, strings.ToLower(word))
	}
	return cleanedInput
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commands map[string]cliCommand

func commandExit() error {
	fmt.Println("Closing the PokedexGo... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the PokedexGo!")
	fmt.Println("Usage: ")
	fmt.Println("help: Show the list of available commands")
	fmt.Println()
	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func main() {
	commands = map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
      callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the PokedexGo",
			callback:    commandExit,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("PokedexGo > ")
		scanner.Scan()
		input := scanner.Text()
		words := cleanInput(input)
		command := ""
		if len(words) != 0 {
			command = words[0]
		}

		if cmd, ok := commands[command]; ok {
			err := cmd.callback()
			if err != nil {
				fmt.Println("Error while running the command: ", err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}

}

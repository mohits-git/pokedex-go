package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type config struct {
	nextLocationUrl *string
	prevLocationUrl *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func cleanInput(text string) []string {
	cleanedInput := []string{}
	for _, word := range strings.Fields(text) {
		cleanedInput = append(cleanedInput, strings.ToLower(word))
	}
	return cleanedInput
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func startRepl(cfg *config) {
	commands := getCommands()
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
			err := cmd.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

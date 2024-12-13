package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
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

type config struct {
	Next     *string
	Previous *string
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
	config      *config
}

type LocationArea struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type LocationAreasResponse struct {
	Count    int            `json:"count"`
	Next     *string        `json:"next"`
	Previous *string        `json:"previous"`
	Results  []LocationArea `json:"results"`
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

func processLocationAreasCommand(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	result := LocationAreasResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return err
	}
	commands["map"].config.Next = result.Next
	commands["map"].config.Previous = result.Previous
	for _, locationArea := range result.Results {
		fmt.Println(locationArea.Name)
	}
	return nil
}

func commandMap() error {
	url := commands["map"].config.Next
	if url == nil {
		fmt.Println("you're on the last page")
		return nil
	}
	return processLocationAreasCommand(*url)
}

func commandMapb() error {
	url := commands["mapb"].config.Previous
	if url == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	return processLocationAreasCommand(*url)
}

func main() {
	locationAreaApiUrl := "https://pokeapi.co/api/v2/location-area"
	mapConfig := &config{
		Next:     &locationAreaApiUrl,
		Previous: nil,
	}

	commands = map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
			config:      nil,
		},
		"exit": {
			name:        "exit",
			description: "Exit the PokedexGo",
			callback:    commandExit,
			config:      nil,
		},
		"map": {
			name:        "map",
			description: "Displays the names of next 20 location areas in the Pokemon world",
			callback:    commandMap,
			config:      mapConfig,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of previous 20 location areas in the Pokemon world",
			callback:    commandMapb,
			config:      mapConfig,
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

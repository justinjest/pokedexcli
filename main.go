package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	commands := getCommands()
	for {
		var input string
		fmt.Printf("Pokedex > ")
		fmt.Scan(&input)
		command, exists := commands[input]
		if exists {
			err := command.callback()
			if err != nil {
				fmt.Printf("Error running command.callback %v\n", err)
			}
		} else {
			fmt.Printf("Sorry, that's not a valid input \n")
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

type location struct {
	name string `json:"name"`
	url  string `json:"url"`
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Provides the next 20 locations in the map",
			callback:    mapNext,
		},
		"mapb": {
			name:        "mapb",
			description: "Provides the previous 20 locations in the map",
			callback:    mapPrevious,
		},
	}
}
func mapPrevious() error {
	url := "https://pokeapi.co/api/v2/location-area/"
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var locations []location
	if err := json.Unmarshal(data, &locations); err != nil {
		return err
	}

	fmt.Printf("%v\n", locations)
	return nil
}
func mapNext() error {
	url := "https://pokeapi.co/api/v2/location-area/"
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var locations []location
	if err := json.Unmarshal(data, &locations); err != nil {
		return err
	}

	fmt.Printf("%v\n", locations)
	return nil
}

func commandHelp() error {
	commands := getCommands()
	fmt.Printf("Welcome to the Pokedex!\n")
	fmt.Printf("Usage:\n\n")
	for _, cmd := range commands {
		fmt.Printf("%v: %v \n", cmd.name, cmd.description)
	}
	fmt.Printf("\n")
	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}

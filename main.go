package main

import (
	"fmt"
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
	}
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

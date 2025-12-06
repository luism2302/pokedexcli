package main

import (
	"fmt"
	"github.com/luism2302/pokedexcli/pokeapi"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*pokeapi.Config) error
}

func getCommands() map[string]cliCommand {
	supportedCommands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 location-areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 location-areas",
			callback:    commandMapb,
		},
	}
	return supportedCommands
}

func commandExit(config *pokeapi.Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *pokeapi.Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Print("\n")
	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func commandMap(config *pokeapi.Config) error {
	locations, err := pokeapi.GetLocationAreas(config.Next, config)
	if err != nil {
		return err
	}
	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapb(config *pokeapi.Config) error {
	locations, err := pokeapi.GetLocationAreas(config.Previous, config)
	if err != nil {
		return err
	}
	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	return nil
}

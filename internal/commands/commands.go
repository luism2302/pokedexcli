package commands

import (
	"fmt"
	"github.com/luism2302/pokedexcli/internal/pokeapi"
	"os"
)

type CliCommand struct {
	Name        string
	Description string
	Callback    func(*pokeapi.Config) error
}

func GetCommands() map[string]CliCommand {
	supportedCommands := map[string]CliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    commandHelp,
		},
		"map": {
			Name:        "map",
			Description: "Displays the next 20 location-areas",
			Callback:    commandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Displays the previous 20 location-areas",
			Callback:    commandMapb,
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
	for _, command := range GetCommands() {
		fmt.Printf("%s: %s\n", command.Name, command.Description)
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

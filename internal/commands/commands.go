package commands

import (
	"fmt"
	"github.com/luism2302/pokedexcli/internal/pokeapi"
	"os"
)

type CliCommand struct {
	Name        string
	Description string
	Callback    func(*pokeapi.Config, ...string) error
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
		"explore": {
			Name:        "explore",
			Description: "Lists the pokemon encounters in <area-name>",
			Callback:    commandExplore,
		},
	}
	return supportedCommands
}

func commandExit(config *pokeapi.Config, parameters ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *pokeapi.Config, parameters ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Print("\n")
	for _, command := range GetCommands() {
		fmt.Printf("%s: %s\n", command.Name, command.Description)
	}
	return nil
}

func commandMap(config *pokeapi.Config, parameters ...string) error {
	locations, err := config.PokeClient.GetLocationAreas(config.Next, config)
	if err != nil {
		return err
	}
	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	config.Next = locations.Next
	if previous, ok := locations.Previous.(string); !ok {
		config.Previous = ""
	} else {
		config.Previous = previous
	}
	return nil
}

func commandMapb(config *pokeapi.Config, parameters ...string) error {
	locations, err := config.PokeClient.GetLocationAreas(config.Previous, config)
	if err != nil {
		return err
	}
	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	config.Next = locations.Next
	if previous, ok := locations.Previous.(string); !ok {
		config.Previous = ""
	} else {
		config.Previous = previous
	}
	return nil
}

func commandExplore(config *pokeapi.Config, parameters ...string) error {
	if len(parameters) < 1 {
		fmt.Println("Usage: explore <area-name>")
		return nil
	}
	area := parameters[0]
	locationData, err := config.PokeClient.GetPokemonEncounters(area)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", area)
	fmt.Println("Found Pokemon:")
	for _, pokemon := range locationData.PokemonEncounters {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}
	return nil
}

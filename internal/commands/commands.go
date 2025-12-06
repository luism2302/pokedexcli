package commands

import (
	"fmt"
	"github.com/luism2302/pokedexcli/internal/pokeapi"
	"math/rand"
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
		"catch": {
			Name:        "catch",
			Description: "Attempts to catch <pokemon-name>",
			Callback:    commandCatch,
		},
		"inspect": {
			Name:        "inspect",
			Description: "Inspects <pokemon-name>",
			Callback:    commandInspect,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "Display the names of caught pokemon",
			Callback:    commandPokedex,
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

func commandCatch(config *pokeapi.Config, parameters ...string) error {
	if len(parameters) < 1 {
		fmt.Println("Usage: catch <pokemon-name>")
		return nil
	}
	pokemonName := parameters[0]
	pokemon, err := config.PokeClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	catchDifficulty := 100.0 / (100.0 + float64(pokemon.BaseExperience))
	prob := rand.Float64()

	if prob < catchDifficulty {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}
	config.Pokedex[pokemon.Name] = pokemon
	fmt.Printf("%s was caught!\n", pokemon.Name)
	return nil
}

func commandInspect(config *pokeapi.Config, parameters ...string) error {
	if len(parameters) < 1 {
		fmt.Println("Usage: inspect <pokemon-name>")
		return nil
	}
	name := parameters[0]
	pokemon, ok := config.Pokedex[name]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("	-%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, pokemonType := range pokemon.Types {
		fmt.Printf("	- %s\n", pokemonType.Type.Name)
	}
	return nil
}

func commandPokedex(config *pokeapi.Config, parameters ...string) error {
	fmt.Println("Your Pokedex:")
	for name, _ := range config.Pokedex {
		fmt.Printf("	-%s\n", name)
	}
	return nil
}

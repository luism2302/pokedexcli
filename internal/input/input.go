package input

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/luism2302/pokedexcli/internal/pokeapi"
	"github.com/luism2302/pokedexcli/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(s *pokeapi.State, c *pokecache.Cache, parameters ...string) error
}

func getSupportedCommands() map[string]cliCommand {
	supportedCommands := map[string]cliCommand{
		"exit":    {name: "exit", description: "Exit the Pokedex", callback: commandExit},
		"help":    {name: "help", description: "Displays a help message", callback: commandHelp},
		"map":     {name: "map", description: "Displays 20 locations areas in the Pokemon World", callback: commandMap},
		"mapb":    {name: "mapb", description: "Displays the previous 20 locations areas in the Pokemon World", callback: commandMapB},
		"explore": {name: "explore", description: "Displays the pokemon in the <area_name> typed", callback: commandExplore},
	}
	return supportedCommands
}

func commandExit(s *pokeapi.State, c *pokecache.Cache, parameters ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(s *pokeapi.State, c *pokecache.Cache, parameters ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, command := range getSupportedCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func commandMap(s *pokeapi.State, c *pokecache.Cache, parameters ...string) error {
	locations, err := pokeapi.GetLocationAreas(s.CurrentUrl, c)
	if err != nil {
		return err
	}
	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	s.PreviousUrl = locations.Previous
	s.CurrentUrl = locations.Next
	return nil
}

func commandMapB(s *pokeapi.State, c *pokecache.Cache, parameters ...string) error {
	if s.PreviousUrl == "" {
		fmt.Println("You are on the first page")
		return nil
	}
	locations, err := pokeapi.GetLocationAreas(s.PreviousUrl, c)
	if err != nil {
		return err
	}
	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	s.CurrentUrl = s.PreviousUrl
	if locations.Previous == "null" {
		s.PreviousUrl = ""
		return nil
	}
	s.PreviousUrl = locations.Previous
	return nil
}

func commandExplore(s *pokeapi.State, c *pokecache.Cache, parameters ...string) error {
	if parameters == nil {
		fmt.Println("Empty arguments")
		return fmt.Errorf("empty <area_name>")
	}
	locationData, err := pokeapi.GetPokemonInArea(parameters[0], c)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Printf("Exploring %s...\n", parameters[0])
	fmt.Println("Found Pokemon:")
	for _, v := range locationData.PokemonEncounters {
		fmt.Printf("- %s\n", v.Pokemon.Name)
	}
	return nil

}

func commandCatch(s *pokeapi.State, c *pokecache.Cache, parameters ...string) error {
	if parameters == nil {
		fmt.Println("Empty Pokemon name")
		return fmt.Errorf("empty <pokemon_name>")
	}
}
func CleanInput(text string) []string {
	if text == "" {
		return []string{}
	}

	lower := strings.ToLower(text)
	trimmed := strings.Fields(lower)

	return trimmed
}

func Repl(s *pokeapi.State, c *pokecache.Cache) {
	scanner := bufio.NewScanner(os.Stdin)
	flag := true

	for flag {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := CleanInput(scanner.Text())

		if len(words) <= 0 {
			fmt.Println("Unknown command")
			continue
		}
		supportedCommands := getSupportedCommands()
		if _, ok := supportedCommands[words[0]]; !ok {
			fmt.Println("Unkown command")
			continue
		}
		if len(words) == 1 {
			supportedCommands[words[0]].callback(s, c)
		} else {
			supportedCommands[words[0]].callback(s, c, words[1])
		}
	}
}

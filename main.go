package main

import (
	"bufio"
	"fmt"
	"github.com/luism2302/pokedexcli/pokeapi"
	"github.com/luism2302/pokedexcli/text"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	config := &pokeapi.Config{
		Previous: "",
		Next:     "https://pokeapi.co/api/v2/location-area/",
	}
	for {
		fmt.Print("Pokedex > ")
		hasInput := scanner.Scan()

		if !hasInput {
			continue
		}

		cleanInput := text.CleanInput(scanner.Text())[0]
		supportedCommands := getCommands()

		if command, ok := supportedCommands[cleanInput]; ok {
			err := command.callback(config)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}

	}
}

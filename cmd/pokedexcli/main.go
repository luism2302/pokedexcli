package main

import (
	"bufio"
	"fmt"
	"github.com/luism2302/pokedexcli/internal/commands"
	"github.com/luism2302/pokedexcli/internal/pokeapi"
	"github.com/luism2302/pokedexcli/internal/text"
	"os"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	config := &pokeapi.Config{
		PokeClient: pokeapi.NewClient(10*time.Second, 5*time.Second),
		Pokedex:    make(map[string]pokeapi.Pokemon),
		Previous:   "",
		Next:       "https://pokeapi.co/api/v2/location-area/",
	}

	for {
		fmt.Print("Pokedex > ")
		hasInput := scanner.Scan()

		if !hasInput {
			continue
		}
		cleanInput := text.CleanInput(scanner.Text())
		command := cleanInput[0]
		parameters := []string{}
		if len(cleanInput) > 1 {
			parameters = cleanInput[1:]
		}
		supportedCommands := commands.GetCommands()

		if command, ok := supportedCommands[command]; ok {
			err := command.Callback(config, parameters...)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}

	}
}

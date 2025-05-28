package main

import (
	"github.com/luism2302/pokedexcli/internal/input"
	"github.com/luism2302/pokedexcli/internal/pokeapi"
)

func main() {
	currentState := pokeapi.State{CurrentUrl: "https://pokeapi.co/api/v2/location-area/", PreviousUrl: ""}
	input.Repl(&currentState)
}

package main

import (
	"time"

	"github.com/luism2302/pokedexcli/internal/input"
	"github.com/luism2302/pokedexcli/internal/pokeapi"
	"github.com/luism2302/pokedexcli/internal/pokecache"
)

func main() {
	currentState := pokeapi.State{CurrentUrl: "https://pokeapi.co/api/v2/location-area/", PreviousUrl: ""}
	cache := pokecache.NewCache(5 * time.Second)
	input.Repl(&currentState, cache)
}

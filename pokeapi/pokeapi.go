package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Config struct {
	Previous string
	Next     string
}

type LocationAreas struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocationAreas(url string, config *Config) (LocationAreas, error) {
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area/"
	}
	response, err := http.Get(url)
	if err != nil {
		return LocationAreas{}, fmt.Errorf("error receiving response: %w", err)
	}

	defer response.Body.Close()

	var locations LocationAreas
	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&locations); err != nil {
		return LocationAreas{}, fmt.Errorf("error decoding response: %w", err)
	}
	config.Next = locations.Next
	previous, ok := locations.Previous.(string)
	if ok {
		config.Previous = previous
	} else {
		config.Previous = ""
	}

	return locations, nil
}

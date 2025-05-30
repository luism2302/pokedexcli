package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/luism2302/pokedexcli/internal/pokecache"
)

type LocationArea struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type State struct {
	CurrentUrl  string
	PreviousUrl string
}

func GetLocationAreas(url string, cache *pokecache.Cache) (*LocationArea, error) {
	var locationAreas LocationArea
	//check if its cached
	if value, ok := cache.Get(url); ok {
		data := value
		err := json.Unmarshal(data, &locationAreas)
		if err != nil {
			return nil, fmt.Errorf("error unmarshaling into struct: %w", err)
		}
		return &locationAreas, nil
	} else {
		//new client
		client := http.Client{}
		//create request
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, fmt.Errorf("couldnt create the request: %w", err)
		}
		//make request
		res, err := client.Do(req)
		if err != nil {
			return nil, fmt.Errorf("couldnt make the request: %w", err)
		}
		defer res.Body.Close()
		data, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, fmt.Errorf("couldnt read response: %w", err)
		}
		err = json.Unmarshal(data, &locationAreas)
		if err != nil {
			return nil, fmt.Errorf("error unmarshaling into struct: %w", err)
		}

		cache.Add(url, data)
		return &locationAreas, nil
	}
}

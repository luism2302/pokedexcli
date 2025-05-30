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

type LocationAreaData struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
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

func GetPokemonInArea(area_name string, cache *pokecache.Cache) (*LocationAreaData, error) {
	var LocationData LocationAreaData
	url := "https://pokeapi.co/api/v2/location-area/" + area_name
	//check if its cached
	if value, ok := cache.Get(url); ok {
		data := value
		err := json.Unmarshal(data, &LocationData)
		if err != nil {
			return nil, fmt.Errorf("error unmarshaling into struct: %w", err)
		}
		return &LocationData, nil
	} else {
		//new client
		client := http.Client{}
		//request
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, fmt.Errorf("error couldnt create the request: %w", err)
		}
		//make the request
		res, err := client.Do(req)
		if err != nil {
			return nil, fmt.Errorf("couldnt make the request: %w", err)
		}
		defer res.Body.Close()
		data, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, fmt.Errorf("couldnt read response: %w", err)
		}
		if string(data) == "Not Found" {
			return nil, fmt.Errorf("area doesnt exist")
		}
		err = json.Unmarshal(data, &LocationData)
		if err != nil {
			return nil, fmt.Errorf("error unmarshaling into struct: %w", err)
		}
		cache.Add(url, data)
		return &LocationData, nil
	}
}

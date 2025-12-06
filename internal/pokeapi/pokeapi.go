package pokeapi

import (
	"encoding/json"
	"fmt"
	"github.com/luism2302/pokedexcli/internal/pokecache"
	"io"
	"net/http"
	"time"
)

const baseUrl = "https://pokeapi.co/api/v2"

type Config struct {
	PokeClient Client
	Pokedex    map[string]Pokemon
	Previous   string
	Next       string
}

type Client struct {
	cache  pokecache.Cache
	client http.Client
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: *pokecache.NewCache(cacheInterval),
		client: http.Client{
			Timeout: timeout,
		},
	}
}

func (client *Client) GetLocationAreas(url string, config *Config) (LocationAreas, error) {
	if url == "" {
		url = fmt.Sprintf("%s/location-area", baseUrl)
	}
	if value, ok := client.cache.Get(url); ok {
		var locations LocationAreas
		err := json.Unmarshal(value, &locations)
		if err != nil {
			return LocationAreas{}, fmt.Errorf("error unmarshaling json: %w", err)
		}
		return locations, nil
	}
	response, err := makeGetRequest(client, url)
	if err != nil {
		return LocationAreas{}, err
	}

	defer response.Body.Close()
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return LocationAreas{}, fmt.Errorf("error reading response body: %w", err)
	}
	var locations LocationAreas
	err = json.Unmarshal(data, &locations)
	if err != nil {
		return LocationAreas{}, fmt.Errorf("error unmarshaling data: %w", err)
	}
	client.cache.Add(url, data)
	return locations, nil
}

func (client *Client) GetPokemonEncounters(areaName string) (ParticularLocation, error) {
	url := fmt.Sprintf("%s/location-area/%s", baseUrl, areaName)
	if value, ok := client.cache.Get(url); ok {
		var locations ParticularLocation
		err := json.Unmarshal(value, &locations)
		if err != nil {
			return ParticularLocation{}, fmt.Errorf("error unmarshaling json: %w", err)
		}
		return locations, nil
	}

	response, err := makeGetRequest(client, url)
	if err != nil {
		return ParticularLocation{}, err
	}
	if response.StatusCode != http.StatusOK {
		return ParticularLocation{}, fmt.Errorf("bad status code: %d", response.StatusCode)
	}
	defer response.Body.Close()
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return ParticularLocation{}, fmt.Errorf("error reading response body: %w", err)
	}
	var locations ParticularLocation
	err = json.Unmarshal(data, &locations)
	if err != nil {
		return ParticularLocation{}, fmt.Errorf("error unmarshaling data: %w", err)
	}
	client.cache.Add(url, data)
	return locations, nil
}

func (client *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	url := fmt.Sprintf("%s/pokemon/%s", baseUrl, pokemonName)
	if value, ok := client.cache.Get(url); ok {
		var pokemon Pokemon
		err := json.Unmarshal(value, &pokemon)
		if err != nil {
			return Pokemon{}, fmt.Errorf("error unmarshaling json: %w", err)
		}
		return pokemon, nil
	}
	response, err := makeGetRequest(client, url)
	if err != nil {
		return Pokemon{}, err
	}
	if response.StatusCode != http.StatusOK {
		return Pokemon{}, fmt.Errorf("bad status code: %d", response.StatusCode)
	}
	defer response.Body.Close()
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return Pokemon{}, fmt.Errorf("error reading response: %w", err)
	}
	var pokemon Pokemon
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, fmt.Errorf("error unmarshaling data: %w", err)
	}
	client.cache.Add(url, data)
	return pokemon, nil
}

func makeGetRequest(client *Client, url string) (*http.Response, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error making request")
	}
	response, err := client.client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("error receiving response")
	}
	return response, nil
}

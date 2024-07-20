package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(name *string) (Pokemon, error) {
	url := baseURL + "/pokemon/"
	if name == nil {
		return Pokemon{}, errors.New("no pokemon indicated")
	} else {
		url += *name
	}

	if val, ok := c.cache.Get(url); ok {
		pokemonResp := Pokemon{}
		err := json.Unmarshal(val, &pokemonResp)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemonResp := Pokemon{}
	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		if _, ok := err.(*json.SyntaxError); ok {
			return Pokemon{}, errors.New("pokemon unknown")
		}
		return Pokemon{}, err
	}

	c.cache.Add(url, dat)
	return pokemonResp, nil
}

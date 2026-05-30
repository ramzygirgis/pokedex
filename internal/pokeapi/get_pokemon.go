package pokeapi

import (
	"encoding/json"
	"net/http"
	"fmt"
)

func (client *Client) PokemonCall(name string) (Pokemon, error) {
		baseUrl := "https://pokeapi.co/api/v2/pokemon/"
		url := fmt.Sprintf("%s%s%s", baseUrl, name, "/")

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return Pokemon{}, err
		}

		resp, err := client.httpClient.Do(req)
		if err != nil {
			return Pokemon{}, err
		}

		var data Pokemon
		if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			return Pokemon{}, err
		}

		return data, nil
	}



func CatchSuccessProbability(pokemon Pokemon) float64 {
	BaseExp := pokemon.BaseExperience
	if BaseExp < 75 {
		return 0.8
	}
	return 55.0/float64(BaseExp)
}


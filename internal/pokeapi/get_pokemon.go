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


func (p Pokemon) Print() {
	fmt.Printf("Name: %s\n", p.Name)
	fmt.Printf("Height: %d\n", p.Height)
	fmt.Printf("Weight: %d\n", p.Weight)
	fmt.Println("Stats:")
	for _, s := range p.Stats {
		fmt.Printf("	-%s: %d\n", s.Stat.Name, s.Effort)
	}
	fmt.Println("Types:")
	for _, t := range p.Types {
		fmt.Printf("	- %s\n", t.Type.Name)
	}
}


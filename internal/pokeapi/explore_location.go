package pokeapi

import (
	"encoding/json"
	"net/http"
	"fmt"
)

func (client *Client) ExploreCall(name string) ([]string, error) {
		baseUrl := "https://pokeapi.co/api/v2/location-area/"
		url := fmt.Sprintf("%s%s%s", baseUrl, name, "/")

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return []string{}, err
		}

		resp, err := client.httpClient.Do(req)
		if err != nil {
			return []string{}, err
		}

		var data ExploreResult
		if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			return []string{}, err
		}
		
		names := make([]string, len(data.PokemonEncounters))

		for i, p := range data.PokemonEncounters {
			names[i] = p.Pokemon.Name
		}

		return names, nil
	}

func ListEncounters(name string, pokemonList []string) error {
		fmt.Printf("Exploring %s...\n", name)
		fmt.Println("Found Pokemon:")
		for i := 0; i < len(pokemonList); i++ {
			fmt.Printf("- %s\n", pokemonList[i])
		}
		return nil
}

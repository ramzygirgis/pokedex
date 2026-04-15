package pokeapi

import (
	"encoding/json"
	"net/http"
	"fmt"
	"github.com/ramzygirgis/pokedex/internal/pokecache"
	"github.com/ramzygirgis/pokedex/internal/shared_types"
)


func (client *Client) PokeapiCall(url string) (shared_types.locationArea, err) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return shared_types.locationArea{}, err
	}

	resp, err := client.httpClient.Do(req)
	if err != nil {
		return shared_types.locationArea{}, err
	}
	defer resp.Body.Close()

	var data shared_types.locationArea
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return shared_types.locationArea{}, err
	}

	return data, nil
	
}

func ListLocations(forwards bool, data shared_types.locationArea) (string, error) {
	
	for i := 0; i < len(data.Results); i++ {
		fmt.Println(data.Results[i].Name)
	}

	if forwards {
		return data.Next, nil
	}
	return data.Previous, nil
}

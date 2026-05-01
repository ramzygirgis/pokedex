package pokeapi

import (
	"encoding/json"
	"net/http"
	"fmt"
)


func (client *Client) LocationsCall(url string) (LocationArea, error) {
	
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationArea{}, err
	}

	resp, err := client.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer resp.Body.Close()

	var data LocationArea
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return LocationArea{}, err
	}

	return data, nil
	
}

func ListLocations(forwards bool, data LocationArea) (string, error) {
	
	for i := 0; i < len(data.Results); i++ {
		fmt.Println(data.Results[i].Name)
	}

	if forwards {
		return data.Next, nil
	}
	return data.Previous, nil
}

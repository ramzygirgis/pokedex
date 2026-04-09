package pokeapi

import (
	"encoding/json"
	"net/http"
	"fmt"
)

func (client *Client) ListLocations(forwards bool, url string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	resp, err := client.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var data locationArea
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", err
	}

	for i := 0; i < len(data.Results); i++ {
		fmt.Println(data.Results[i].Name)
	}

	if forwards {
		return data.Next, nil
	}
	return data.Previous, nil
}

package pokeapi

func (client *Client) ListLocations(forwards bool, c *config) error {
	if forwards {
		link := c.next
	} else {
			if c.previous == "" {
			fmt.Println("No previous locations\n")
			return nil
		}
		link := c.previous
	}

	req, err := http.NewRequest("GET", link, nil)
	if err != nil {
		return err
	}

	resp, err := client.httpClient.Do(req)
	if err != nil {
		return err
	}
	refer resp.Body.Close()

	var data locationArea
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return err
	}

	for i := 0; i < len(data.Results); i++ {
		fmt.Println(data.Results[i].Name)
	}

	if forwards {
		c.previous = c.next
		c.next = data.Next
	} else {
		c.next = c.previous
		c.previous = data.Previous
	}
	return nil
}

package pokeapi

import (
	"encoding/json"
	// "fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocation(area string) (Location, error) {
	url := baseURL + "/location-area/" + area

	if val, ok := c.cache.Get(url); ok {
		location := Location{}
		err := json.Unmarshal(val, &location)
		if err != nil {
			return Location{}, nil
		}
		return location, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}

	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Location{}, err
	}
	// fmt.Println(string(dat))
	location := Location{}
	err = json.Unmarshal(dat, &location)
	if err != nil {
		return Location{}, err
	}

	c.cache.Add(url, dat)

	return location, nil
}

package pokeapi

import (
	"encoding/json"
	"net/http"
)

// Returns a list of 20 location-areas names and urls as per the given page url or the first page
func (c *Client) ListLocations(pageUrl *string) (LocationAreasResponse, error) {
	url := baseUrl + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}

	resp, err := http.Get(url)
	if err != nil {
		return LocationAreasResponse{}, err
	}
	defer resp.Body.Close()

	locationsResp := LocationAreasResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&locationsResp); err != nil {
		return LocationAreasResponse{}, err
	}

	return locationsResp, nil
}

package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// Returns a list of 20 location-areas names and urls as per the given page url or the first page
func (c *Client) ListLocations(pageUrl *string) (LocationAreasResponse, error) {
	url := baseUrl + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}

  cachedResp, cached := c.cache.Get(url)
  if cached {
    locationsResp := LocationAreasResponse{}
    if err := json.Unmarshal(cachedResp, &locationsResp); err != nil {
      return LocationAreasResponse{}, err
    }
    return locationsResp, nil 
  }

	resp, err := http.Get(url)
	if err != nil {
		return LocationAreasResponse{}, err
	}
	defer resp.Body.Close()

  data, err := io.ReadAll(resp.Body)
  if err != nil {
    return LocationAreasResponse{}, err
  }

	locationsResp := LocationAreasResponse{}
  err = json.Unmarshal(data, &locationsResp)
  if err != nil {
    return LocationAreasResponse{}, err
  }

  c.cache.Add(url, data)

	return locationsResp, nil
}

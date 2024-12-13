package pokeapi

import (
	"net/http"
	"time"
)

// PokeAPI Client
type Client struct {
	httpClient http.Client
}

// Creates a new PokeAPI client
func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}

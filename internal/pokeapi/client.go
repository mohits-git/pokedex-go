package pokeapi

import (
	"net/http"
	"time"

	"github.com/mohits-git/pokedex-go/internal/pokecache"
)

// PokeAPI Client
type Client struct {
	httpClient http.Client
	cache      *pokecache.Cache
}

// Creates a new PokeAPI client
func NewClient(timeout time.Duration, maxCacheTime ...time.Duration) *Client {
	if len(maxCacheTime) == 0 {
		maxCacheTime = append(maxCacheTime, 1*time.Minute)
	}
	return &Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: pokecache.NewCache(maxCacheTime[0]),
	}
}

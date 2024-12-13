package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *config) error {
	fmt.Println("Closing the PokedexGo... Goodbye!")
	os.Exit(0)
	return nil
}

package main

import (
	"fmt"
	"strings"
)

func cleanInput(text string) []string {
  cleanedInput := []string{}
  for _, word := range strings.Fields(text) {
    cleanedInput = append(cleanedInput, strings.ToLower(word))
  }
  return cleanedInput
}

func main() {
    fmt.Println("Hello, World!")
}

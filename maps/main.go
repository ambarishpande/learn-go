package main

import (
	"fmt"
)

func main() {
	colors := make(map[string]string)
	colors["red"] = "#FF0000"
	colors["green"] = "#00FD00"
	colors["blue"] = "#0000FF"
	printMap(colors)
}

func printMap(m map[string]string) {
	for color, code := range m {
		fmt.Printf("%v - %v\n", color, code)
	}
}

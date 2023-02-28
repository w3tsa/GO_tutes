package main

import "fmt"

func main() {
	colors := map[string] string {
		"red": "#FF0000",
		"green" : "#FFA500",
		"yellow" : "#FFFF00",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println("Hex code for", key, "is", value)
	}
}
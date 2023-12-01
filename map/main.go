package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf8lk",
	}

	print(colors)
}

func print(m map[string]string) {
	for key, value := range m {
		fmt.Println(key, ": ", value)
	}
}

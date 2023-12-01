package main

import "fmt"

func main() {
	number := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, n := range number {
		if n%2 == 0 {
			fmt.Println(i, " is even")
		} else {
			fmt.Println(i, " is odd")
		}
	}
}

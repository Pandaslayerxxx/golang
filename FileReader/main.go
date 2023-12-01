package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	arg := os.Args[1]
	f, err := os.Open(arg)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)
}

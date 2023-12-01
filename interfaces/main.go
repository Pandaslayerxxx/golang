package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}
type integer int

type bot interface {
	getGreeting() string
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	var x integer
	printGreeting(x)
	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (integer) getGreeting() string {
	return "number"
}

func (englishBot) getGreeting() string {
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	return "Hola amigo!"
}

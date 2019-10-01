package main

import "fmt"

func main() {
	var locale string = "fr"
	var greeting string

	if locale == "en" {
		greeting = "Hello"
	} else if locale == "es" {
		greeting = "Hola"
	} else if locale == "de" {
		greeting = "Guten Tag"
	} else {
		greeting = "Yo"
	}

	fmt.Printf(greeting + ", Go!\n")
}

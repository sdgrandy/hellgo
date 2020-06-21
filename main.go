package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

var HELLGO_API_URL string

func main() {
	//url := os.Getenv("HELLGO_API_URL")
	file, _ := ioutil.ReadFile("vars.sh")
	s := string(file)
	fmt.Println("file contents:", s)
	port := os.Getenv("HELLGO_API_PORT")
	fmt.Println("port: ", port)

	url := os.Getenv("HELLGO_API_URL")
	fmt.Println("url: ", url)

	user := os.Getenv("HELLGO_API_USER")
	fmt.Println("user: ", user)

	pw := os.Getenv("HELLGO_API_PASSWORD")
	fmt.Println("password: ", pw)

	var locale, greeting string
	var languages = [4]string{"en", "es", "de", "fr"}
	locale = languages[1]

	/*
		if locale == "en" {
			greeting = "Hello"
		} else if locale == "es" {
			greeting = "Hola"
		} else if locale == "de" {
			greeting = "Guten Tag"
		} else {
			greeting = "Yo"
		}
	*/

	switch locale {
	case "en":
		greeting = "Hello"
	case "es":
		greeting = "Hola"
	case "de":
		greeting = "Guten Tag"
	case "fr":
		greeting = "Bonjour"
	default:
		greeting = "Yo"
	}

	fmt.Printf(greeting + ", Go!\n")
}

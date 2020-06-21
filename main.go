package main

import (
	"fmt"
	"os"
	"strings"
)

var HELLGO_API_URL string

func main() {
	//url := os.Getenv("HELLGO_API_URL")
	// file, _ := ioutil.ReadFile("vars.sh")
	// s := string(file)
	// fmt.Println("file contents:", s)
	port := os.Getenv("HELLGO_API_PORT")
	fmt.Println("port: ", port)

	url := os.Getenv("HELLGO_API_URL")
	fmt.Println("url: ", url)

	user := os.Getenv("HELLGO_API_USER")
	user = replace(user)
	fmt.Println("user: ", user)

	pw := os.Getenv("HELLGO_API_PASSWORD")
	pw = replace(pw)
	fmt.Println("password: ", pw)
	admin := "[ a, d, m, i, n ]"
	admin = replace(admin)
	fmt.Println("admin:", admin)

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

func replace(s string) string {
	s = s[1 : len(s)-2]
	s = strings.ReplaceAll(s, ",", "")
	s = strings.TrimSpace(s)
	var r []rune
	for _, v := range s {
		if v != 32 {
			r = append(r, v)
		}
	}
	s = string(r)
	return s
}

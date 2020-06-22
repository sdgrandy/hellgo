package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

var HELLGO_API_URL string

func main() {
	file, _ := ioutil.ReadFile("vars.env")
	s := string(file)
	fmt.Println("file contents:", s)

	if fileExists("vars.env") {
		setVars("vars.env")
	} else {
		fmt.Println("file does not exist")
	}
	env := os.Getenv("ENVIRONMENT")
	fmt.Println("environment: ", env)

	url := os.Getenv("HELLGO_API_URL")
	fmt.Println("url: ", url)

	port := os.Getenv("HELLGO_API_PORT")
	fmt.Println("port: ", port)

	user := os.Getenv("HELLGO_API_USER")
	fmt.Println("user: ", user)

	pw := os.Getenv("HELLGO_API_PASSWORD")
	fmt.Println("password: ", pw)

	if env == "master" && user == "admin1" && pw == "bimil1" {
		fmt.Println("YES it's MASTER")
	} else if env == "qa" && user == "admin2" && pw == "bimil2" {
		fmt.Println("YES it is QA")
	} else if env == "dev" && user == "admin3" && pw == "bimil3" {
		fmt.Println("YES it is DEV")
	} else {
		fmt.Println("WRONG USERNAME AND/OR PASSWORD")
	}
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

func setVars(f string) error {
	i := 0
	vars := []string{"HELLGO_API_URL", "HELLGO_API_PORT", "HELLGO_API_USER", "HELLGO_API_PASSWORD"}
	file, err := os.Open(f)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		os.Setenv(vars[i], line)
	}
	return scanner.Err()
}

func fileExists(file string) bool {
	info, err := os.Stat(file)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

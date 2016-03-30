package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

const aURL = "http://artii.herokuapp.com"

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Printf("Usage:\n")
		return
	}

	// first argument should be a string
	// so split on spaces to separate words
	s := strings.Split(args[0], " ")

	// if only one word and it is fonts
	// print the available fonts
	if len(s) == 1 && s[0] == "fonts" {
		fmt.Printf("%v", fontList())
		return
	}

	// pass in slice of words to draw function
	fmt.Println(draw(s))
}

func fontList() string {
	url := aURL + "/fonts_list"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	s := string(body) + "\n"
	return s
}

func draw(s []string) string {
	js := strings.Join(s, "+")
	url := fmt.Sprintf("%s/make?text=%s", aURL, js)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	a := string(body)
	return a
}

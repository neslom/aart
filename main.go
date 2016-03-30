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

	switch args[0] {
	case "fonts":
		fmt.Printf("%v", fontList())
		return
	}

	fmt.Println(draw(args))
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
	f := strings.Split(s[0], " ")
	js := strings.Join(f, "+")
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

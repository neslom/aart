package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const aURL = "http://artii.herokuapp.com"
const defaultFont = "mini"

func main() {
	var font string
	flag.StringVar(&font, "f", "", "select a font. to see available fonts, run 'aart fonts'")
	flag.Usage = func() {
		fmt.Println("Usage of aart:\n")
		fmt.Println("\trun 'aart fonts' to see available fonts")
		fmt.Println("\trun 'aart -f yourfont \"and some text\"' to print your text as ascii art!")
	}
	flag.Parse()

	args := flag.Args()

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

	// pass in slice of words and font to draw function
	fmt.Println(draw(s, font))
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

func draw(s []string, font string) string {
	if font == "" {
		font = defaultFont
	}
	js := strings.Join(s, "+")
	f := fmt.Sprintf("&font=%s", font)
	url := fmt.Sprintf("%s/make?text=%s%s", aURL, js, f)

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

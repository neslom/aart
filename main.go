package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

const aURL = "http://artii.herokuapp.com"
const defaultFont = "alligator"

func main() {
	var font string
	var randFont bool
	var printFont bool
	flag.StringVar(&font, "f", "", "select a font")
	flag.BoolVar(&randFont, "r", false, "selects a random font for you")
	flag.BoolVar(&printFont, "p", false, "print font used")
	flag.Usage = func() {
		fmt.Printf("%v\n", draw([]string{"Welcome", "to", "aart!"}, "mini"))
		fmt.Println("run 'aart fonts' to see available fonts")
		fmt.Println("run 'aart -f yourfont \"and some text\"' to print your text as ascii art!")
		flag.PrintDefaults()
	}
	flag.Parse()

	args := flag.Args()

	if len(args) == 0 {
		flag.Usage()
		return
	}

	// first argument should be a string
	// so split on spaces to separate words
	s := strings.Split(args[0], " ")

	// if only one word and it is fonts
	// print the available fonts
	if len(s) == 1 && s[0] == "fonts" {
		fmt.Printf("%v%s", string(fontList()), "\n")
		return
	}

	if randFont {
		font = randomFont()
	}
	// pass in slice of words and font to draw function
	fmt.Println(draw(s, font))

	if printFont {
		fmt.Printf("font used: %s\n", font)
	}
}

func fontList() []byte {
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

	return body
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

func randomFont() string {
	rand.Seed(time.Now().UnixNano())
	fl := string(fontList())
	f := strings.Split(fl, "\n")
	rf := f[rand.Intn(len(f))]
	return rf
}

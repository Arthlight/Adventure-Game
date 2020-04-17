package main

import (
	"flag"
	"fmt"
	"log"
	"html/template"
	"encoding/json"
	"io/ioutil"
)

type Story map[string]Chapter

type Chapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Options `json:"options"`
}

type Options struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

func main() {
	filename := flag.String("file", "adventure.json", "the file to choose for your adventure")
	flag.Parse()

	file, err := ioutil.ReadFile(*filename)
	if err != nil {
		log.Println(err)
	}
	var myStory Chapter
	err = json.Unmarshal(file, &myStory)
	if err != nil {
		log.Println(err)
	}
}

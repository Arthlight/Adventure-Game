package stories

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Story map[string]Chapter

type Chapter struct {
	Title   string    `json:"title"`
	Story   []string  `json:"story"`
	Options []Options `json:"options"`
}

type Options struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

func JsonStory(s string) (Story, error) {
	file, err := ioutil.ReadFile(s)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var myStory Story
	err = json.Unmarshal(file, &myStory)
	if err != nil {
		return nil, err
	}

	return myStory, nil
}
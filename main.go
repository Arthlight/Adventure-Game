package main

import (
	"Adventure-Game/stories"
	"flag"
	"fmt"
	"log"
)

func main() {
	filename := flag.String("file", "./adventure.json", "the file to choose for your adventure")
	flag.Parse()

	myStory, err := stories.JsonStory(*filename)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("%+v", myStory)
}

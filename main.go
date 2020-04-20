package main

import (
	"Adventure-Game/stories"
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := flag.Int("port", 3000, "the port to use for the webserver")
	filename := flag.String("file", "./adventure.json", "the file to choose for your adventure")
	flag.Parse()

	myStory, err := stories.JsonStory(*filename)
	if err != nil {
		log.Println(err)
	}

	handler := stories.NewHandler(myStory)
	fmt.Printf("Starting the webserver on port: %d", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d",*port), handler))
}

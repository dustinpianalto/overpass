package main

import (
	"log"

	"github.com/dustinpianalto/overpass/internal/twitter"
)

func main() {
	tClient := twitter.Connect()
	_, err := twitter.GetTimeline(tClient, 20)
	if err != nil {
		log.Fatal(err)
	}
}

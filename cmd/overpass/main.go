package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/dustinpianalto/overpass/internal/twitter"
)

func main() {
	tClient := twitter.Connect()
	tweetChan, stream := twitter.StartUserScanner(tClient, "3347145623")
	for tweet := range tweetChan {
		log.Printf("%s\n%s - %s\n%s", tweet.IDStr, tweet.User.Name, tweet.User.IDStr, tweet.FullText)
	}
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	log.Println(<-ch)

	fmt.Println("Stopping Stream...")
	stream.Stop()
}

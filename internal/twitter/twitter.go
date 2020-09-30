package twitter

import (
	"log"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dustinpianalto/overpass/internal/oauth1"
)

type Tweet struct {
	twitter.Tweet
}

func Connect() *twitter.Client {
	httpClient := oauth1.GetClient()
	return twitter.NewClient(httpClient)
}

func StartUserScanner(client *twitter.Client, userID string, tweetChan chan *Tweet) *twitter.Stream {
	demux := twitter.NewSwitchDemux()
	demux.Tweet = func(tweet *twitter.Tweet) {
		if tweet.User.IDStr == userID {
			tweetChan <- &Tweet{*tweet}
		}
	}
	demux.StatusDeletion = func(deletion *twitter.StatusDeletion) {
		log.Printf("%#v\n", deletion)
	}
	demux.StreamLimit = func(limit *twitter.StreamLimit) {
		log.Printf("%#v\n", limit)
	}
	demux.StreamDisconnect = func(disconnect *twitter.StreamDisconnect) {
		log.Printf("%#v\n", disconnect)
	}
	demux.Event = func(event *twitter.Event) {
		log.Printf("%#v\n", event)
	}
	log.Printf("Starting Stream for %s\n", userID)
	filterParams := &twitter.StreamFilterParams{
		Follow:        []string{userID},
		StallWarnings: twitter.Bool(true),
	}
	stream, err := client.Streams.Filter(filterParams)
	if err != nil {
		log.Println(err)
	}
	go demux.HandleChan(stream.Messages)
	return stream
}

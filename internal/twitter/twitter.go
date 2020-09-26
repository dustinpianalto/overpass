package twitter

import (
	"fmt"

	"github.com/dghubble/go-twitter/twitter"
	oauth "github.com/dustinpianalto/overpass/internal/oauth2"
)

func Connect() *twitter.Client {
	httpClient := oauth.GetClient()
	return twitter.NewClient(httpClient)
}

func GetTimeline(client *twitter.Client, count int) ([]twitter.Tweet, error) {
	tweets, resp, err := client.Timelines.HomeTimeline(&twitter.HomeTimelineParams{
		Count: count,
	})
	if err != nil {
		return nil, err
	}
	fmt.Println(resp)
	fmt.Printf("%#v", tweets)
	return tweets, nil
}

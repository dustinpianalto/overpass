package tweets

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/dustinpianalto/overpass/internal/twitter"
)

func TweetHandler(session *discordgo.Session, tweetChan <-chan *twitter.Tweet) {

	for tweet := range tweetChan {
		_, err := session.ChannelMessageSend("404569276012560386", fmt.Sprintf("%s\n%s - %s\n%s\n", tweet.IDStr, tweet.User.Name, tweet.User.IDStr, tweet.Text))
		if err != nil {
			log.Println(err)
		}
		log.Printf("%s\n%s - %s\n%s\n", tweet.IDStr, tweet.User.Name, tweet.User.IDStr, tweet.Text)
	}
}

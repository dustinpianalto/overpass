package oauth1

import (
	"net/http"
	"os"

	"github.com/dghubble/oauth1"
)

func GetClient() *http.Client {
	apiKey := os.Getenv("TWITTER_API_KEY")
	apiSecret := os.Getenv("TWITTER_API_SECRET")
	accessKey := os.Getenv("TWITTER_ACCESS_KEY")
	accessSecret := os.Getenv("TWITTER_ACCESS_SECRET")
	config := oauth1.NewConfig(apiKey, apiSecret)
	token := oauth1.NewToken(accessKey, accessSecret)
	// OAuth1 http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)

	return httpClient
}

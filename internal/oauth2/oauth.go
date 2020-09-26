package oauth

import (
	"net/http"
	"os"

	"golang.org/x/oauth2"

	"golang.org/x/oauth2/clientcredentials"
)

func GetClient() *http.Client {
	config := &clientcredentials.Config{
		ClientID:     os.Getenv("TWITTER_API_KEY"),
		ClientSecret: os.Getenv("TWITTER_API_SECRET_KEY"),
		TokenURL:     "https://api.twitter.com/oauth2/token",
	}

	return config.Client(oauth2.NoContext)
}

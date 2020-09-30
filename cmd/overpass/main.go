package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/dustinpianalto/overpass/internal/exts/tweets"
	"github.com/dustinpianalto/overpass/internal/postgres"
	"github.com/dustinpianalto/overpass/internal/twitter"
)

func main() {
	token := os.Getenv("DISCORD_TOKEN")
	if token == "" {
		log.Fatal("Discord token not found in env")
	}
	dbString := os.Getenv("DATABASE_URL")
	if dbString == "" {
		log.Fatal("DB connection string not found in env")
	}
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal("There was an error creating the Discord session")
	}
	session.StateEnabled = true

	_ = postgres.ConnectDatabase(dbString)

	err = session.Open()
	if err != nil {
		log.Fatal("There was an error opening the connection to Discord, ", err)
	}
	log.Println("Bot is running")

	tClient := twitter.Connect()
	tweetChan := make(chan *twitter.Tweet, 100)
	stream := twitter.StartUserScanner(tClient, "34743251", tweetChan)
	stream1 := twitter.StartUserScanner(tClient, "25073877", tweetChan)
	stream2 := twitter.StartUserScanner(tClient, "3347145623", tweetChan)
	go tweets.TweetHandler(session, tweetChan)

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, os.Kill, os.Interrupt)
	log.Println(<-ch)

	fmt.Println("Stopping Stream...")
	stream.Stop()
	stream1.Stop()
	stream2.Stop()
}

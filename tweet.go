package overpass

type Tweet struct {
	ID        string
	MessageID string
	ChannelID string
	GuildID   string
	AuthorID  string
}

type TweetService interface {
	Tweet(id string) (*Tweet, error)
	AddTweet(t *Tweet) error
	DeleteTweet(t *Tweet) error
}

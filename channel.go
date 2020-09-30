package overpass

type Channel struct {
	ID       string
	embed    bool
	retweets bool
	guild    *Guild
}

type ChannelService interface {
	Channel(id string) (*Channel, error)
	AddChannel(c *Channel) error
	UpdateChannel(c *Channel) error
	DeleteChannel(c *Channel) error
}

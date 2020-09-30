package overpass

type Guild struct {
	ID     string
	prefix string
}

type GuildService interface {
	Guild(id string) (*Guild, error)
	AddGuild(g *Guild) error
	DeleteGuild(g *Guild) error
	UpdateGuild(g *Guild) error
	GuildChannels(g *Guild) ([]*Channel, error)
}

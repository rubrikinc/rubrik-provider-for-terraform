package rubrikcdm

import "github.com/rubrikinc/rubrik-sdk-for-go/rubrikcdm"

// Config is per-provider, specifies where to connect to Rubrik CDM
type Config struct {
	NodeIP   string
	Username string
	Password string
}

// Client returns a *rubrik.Client to interact with the configured Rubrik CDM instance
func (c *Config) Client() (*rubrikcdm.Credentials, error) {

	rubrik := rubrikcdm.Connect(c.NodeIP, c.Username, c.Password)

	return rubrik, nil
}

package bca

// Config is config to access BCA API
type Config struct {
	ClientID     string
	ClientSecret string
	APIKey       string
	APISecret    string
	URL          string
	CorporateID  string
	OriginHost   string

	ChannelID    string
	CredentialID string

	LogLevel int

	LogPath string
}

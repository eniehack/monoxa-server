package config

type DatabaseConfig struct {
	Url    string `toml:"url"`
	DBType string `toml:"type"`
}

type Config struct {
	FirebaseCredential string         `toml:"firebase"`
	FrontendURL        []string       `toml:"frontend"`
	Database           DatabaseConfig `toml:"database"`
	FileBucket         string         `toml:"file_backet"`
}

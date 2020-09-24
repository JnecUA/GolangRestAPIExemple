package config

//Config type for server settings
type Config struct {
	DBUrl string            `toml:"dbUrl"`
	SMTP  map[string]string `toml:"smtp"`
}

//DefaultConfig ... Initialize a config type variable
func DefaultConfig() *Config {
	return &Config{
		DBUrl: "",
		SMTP:  map[string]string{},
	}
}

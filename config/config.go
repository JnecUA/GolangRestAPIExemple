package config

//Config type for server settings
type Config struct {
	DBUrl string `toml:"dbUrl"`
}

//DefaultConfig ... Initialize a config type variable
func DefaultConfig() *Config {
	return &Config{
		DBUrl: "",
	}
}

package apiserver

//Config type for server settings
type Config struct {
	IP       string `toml:"ip_addr"`
	Port     int    `toml:"port"`
	LogLevel string `toml:"log_level"`
}

//DefaultConfig ... Initialize a config type variable
func DefaultConfig() *Config {
	return &Config{
		IP:       "127.0.0.1",
		Port:     8080,
		LogLevel: "debug",
	}
}

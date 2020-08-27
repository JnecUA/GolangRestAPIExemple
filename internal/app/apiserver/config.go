package apiserver

//Config type for server settings
type Config struct {
	Ip       string `toml:"ip_addr"`
	Port     int    `toml:"port"`
	LogLevel string `toml:"log_level"`
}

func DefaultConfig() *Config {
	return &Config{
		Ip:       "127.0.0.1",
		Port:     8080,
		LogLevel: "debug",
	}
}

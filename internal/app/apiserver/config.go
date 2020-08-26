package apiserver

//Config type for server settings
type Config struct {
	ip   string `toml:ip_addr`
	port int    `toml:port`
}

func DefaultConfig() *Config {
	return &Config{
		ip:   "127.0.0.1",
		port: 8080,
	}
}

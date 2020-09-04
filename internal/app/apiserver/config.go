package apiserver

//Config type for server settings
type Config struct {
	IP       string `toml:"ip_addr"`
	Port     string `toml:"port"`
	LogLevel string `toml:"log_level"`
}

//DefaultConfig ... Initialize a config type variable
func DefaultConfig() *Config {
	return &Config{
		IP:       "127.0.0.1",
		Port:     "8080",
		LogLevel: "debug",
	}
}

//MapConfiguring ... Change config var, from toml file
func (c *Config) MapConfiguring(m map[string]string) {
	c.IP = m["ip_addr"]
	c.Port = m["port"]
	c.LogLevel = m["log_level"]
}

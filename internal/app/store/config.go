package store

//Config ... Config var with DB settings
type Config struct {
	dbURL string `toml:"dbURL"`
}

//DefaultConfig ... Initialize config var
func DefaultConfig() *Config {
	return &Config{
		dbURL: "",
	}
}

//MapConfiguring ... Change config var, from toml file
func (c *Config) MapConfiguring(m map[string]string) {
	c.dbURL = m["dbURL"]
}

package apiserver

type APIserver struct {
	config *Config
}

func Init(config *Config) *APIserver {
	return &APIserver{
		config: config,
	}
}

func (server *APIserver) Start() error {
	return nil
}

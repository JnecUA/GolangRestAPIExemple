package apiserver

import (
	"github.com/sirupsen/logrus"
)

type APIserver struct {
	config *Config
	logger *logrus.Logger
}

func Init(config *Config) *APIserver {
	return &APIserver{
		config: config,
		logger: logrus.New(),
	}
}

func (s *APIserver) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}
	s.logger.Info("Logger start work, you can change logger log level in configure file")

	return nil
}

func (s *APIserver) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)

	return nil
}

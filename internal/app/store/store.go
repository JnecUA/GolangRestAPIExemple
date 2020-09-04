package store

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

//Store ... Database
type Store struct {
	config *Config
	dbpool *pgxpool.Pool
}

//Init ... Initialize DB variable
func Init(config *Config) *Store {
	return &Store{
		config: config,
	}
}

//OpenConnection ... Start DB connection with REST API
func (s *Store) OpenConnection() error {
	//БЛЯЯЯЯЯ ЧО ДЕЛАТЬ НИХЕРА НЕ ПОНИМАЮ ААААААААААААААААААААААААААААА, ДОДИК ЮЗАЕТ НЕ ПОСТГРЕС, А НА НУЖНУЮ БИБЛЯХУ НЕТ ГАЙДАААААААААА
	dbpool, err := pgxpool.Connect(context.Background(), os.Getenv(s.config.dbURL))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	s.dbpool = dbpool
	return nil
}

//CloseConnection ... Stop DB connection with REST API
func (s *Store) CloseConnection() {
	s.dbpool.Close()
}

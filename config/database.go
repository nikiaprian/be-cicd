package config

import (
	"database/sql"
	"fmt"
	"codein/services"

	_ "github.com/lib/pq"
)

func NewPostgreSQLDB() (*services.PostgreSQL, error) {
	connectionString := fmt.Sprintf(
		"host=10.10.5.173 port=5432 user=postgres password=nikiskripsi dbname=postgres sslmode=disable",
	)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to PostgreSQL!")

	return &services.PostgreSQL{
		DB: db,
	}, nil
}

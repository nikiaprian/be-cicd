package config

import (
	"database/sql"
	"fmt"
	"codein/services"

	_ "github.com/lib/pq"
)

func NewPostgreSQLDB() (*services.PostgreSQL, error) {
	connectionString := fmt.Sprintf(
<<<<<<< HEAD
		"host=10.10.5.53 port=5432 user=postgres password=nikiskripsi dbname=postgres sslmode=disable",
=======
		"host=10.10.5.87 port=5432 user=postgres password=nikiskripsi dbname=postgres sslmode=disable",
>>>>>>> e37d2f6713d3e567389c9ffe3accd79e59cef8b8
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

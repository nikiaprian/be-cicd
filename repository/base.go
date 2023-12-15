package repository

import "kel15/services"

type Repository struct {
	db *services.PostgreSQL
}

func NewRepository(db *services.PostgreSQL) Repository {
	return Repository{db: db}
}

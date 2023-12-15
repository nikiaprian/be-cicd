package repository

import "codein/services"

type Repository struct {
	db *services.PostgreSQL
}

func NewRepository(db *services.PostgreSQL) Repository {
	return Repository{db: db}
}

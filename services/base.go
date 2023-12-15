package services

import "database/sql"

type PostgreSQL struct {
	*sql.DB
}

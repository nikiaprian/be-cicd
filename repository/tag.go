package repository

import (
	"database/sql"
	"fmt"
	"codein/models"
	"time"

	"github.com/gin-gonic/gin"
)

func (repository *Repository) GetTagByID(c *gin.Context, id int64) (*models.Tag, error) {
	var tag models.Tag

	query := "SELECT id, tag FROM Tags WHERE id = $1"
	row := repository.db.QueryRow(query, id)

	err := row.Scan(&tag.ID, &tag.Tag)
	if err != nil {
		return nil, err
	}

	return &tag, nil
}

func (repository *Repository) GetTagByName(c *gin.Context, tag string) (*models.Tag, error) {
	var Tag models.Tag

	query := "SELECT id, tag FROM Tags WHERE tag = $1"
	row := repository.db.QueryRow(query, tag)

	err := row.Scan(&Tag.ID, &Tag.Tag)
	if err == sql.ErrNoRows {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return &Tag, nil
}

func (repository *Repository) CreateTag(c *gin.Context, tag string) (*models.Tag, error) {
	Tag, err := repository.GetTagByName(c, tag)
	// fmt.Println(Tag, err, tag)
	if Tag != nil {
		return Tag, nil
	}

	query := "INSERT INTO Tags (tag, created_at, updated_at) VALUES ($1, $2, $3)"
	result, err := repository.db.Exec(query, tag, time.Now(), time.Now())
	fmt.Println(result, err, "tagerror")

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return repository.GetTagByID(c, id)
}


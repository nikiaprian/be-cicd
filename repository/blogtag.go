package repository

import (
	"codein/models"

	"github.com/gin-gonic/gin"
)

func (repository *Repository) CreateBlogTag(c *gin.Context, blogID int64, tagID int64) (*models.Tag, error) {
	query := "INSERT INTO BlogTags (blog_id, tag_id) VALUES ($1, $2)"
	_, err := repository.db.Exec(query, blogID, tagID)

	if err != nil {
		return nil, err
	}

	return repository.GetTagByID(c, tagID)
}

func (repository *Repository) GetBlogTagByBlogID(c *gin.Context, id int64) (*[]models.BlogTag, error) {
	var blog_tags []models.BlogTag

	query := "SELECT id, blog_id, tag_id FROM BlogTags WHERE blog_id = $1"
	rows, err := repository.db.Query(query, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var blog_tag models.BlogTag
		
		err := rows.Scan(&blog_tag.ID, &blog_tag.BlogID, &blog_tag.TagID)
		if err != nil {
			return nil, err
		}

		blog_tags = append(blog_tags, blog_tag)
	}

	return &blog_tags, nil
}
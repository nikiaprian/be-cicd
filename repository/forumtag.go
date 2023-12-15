package repository

import (
	"codein/models"

	"github.com/gin-gonic/gin"
)

func (repository *Repository) CreateForumTag(c *gin.Context, forumID int64, tagID int64) (*models.Tag, error) {
	query := "INSERT INTO ForumTags (forum_id, tag_id) VALUES ($1, $2)"
	_, err := repository.db.Exec(query, forumID, tagID)

	if err != nil {
		return nil, err
	}

	return repository.GetTagByID(c, tagID)
}

func (repository *Repository) GetForumTagByForumID(c *gin.Context, id int64) (*[]models.Forumtag, error) {
	var forum_tags []models.Forumtag

	query := "SELECT id, forum_id, tag_id FROM ForumTags WHERE forum_id = $1"
	rows, err := repository.db.Query(query, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var forum_tag models.Forumtag

		err := rows.Scan(&forum_tag.ID, &forum_tag.ForumID, &forum_tag.TagID)
		if err != nil {
			return nil, err
		}

		forum_tags = append(forum_tags, forum_tag)
	}

	return &forum_tags, nil
}

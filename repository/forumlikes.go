package repository

import (
	"database/sql"
	"errors"
	"codein/models"
	"time"

	"github.com/gin-gonic/gin"
)

func (repository *Repository) GetAllLikeByForumID(c *gin.Context, id int) (*[]models.ForumsLikesResponse, error) {
	var forumLikes = make([]models.ForumsLikesResponse, 0)

	query := `
        SELECT 
            ForumsLikes.id, ForumsLikes.forum_id, ForumsLikes.created_at, ForumsLikes.updated_at,
            Users.id, Users.username, Users.email, Users.photo
        FROM ForumsLikes 
        JOIN Users ON ForumsLikes.user_id = Users.id
        WHERE ForumsLikes.forum_id = $1
        ORDER BY ForumsLikes.created_at DESC
    `
	rows, err := repository.db.Query(query, id)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var forumLike models.ForumsLikesResponse
		var user models.User

		err := rows.Scan(&forumLike.ID, &forumLike.ForumID, &forumLike.CreatedAt, &forumLike.UpdatedAt, &user.ID, &user.Username, &user.Email, &user.Photo)
		if err != nil {
			return nil, err
		}
		forumLike.User = user

		forumLikes = append(forumLikes, forumLike)
	}

	return &forumLikes, nil
}

func (repository *Repository) GetLikeByUserIDAndForumID(c *gin.Context, user_id, forum_id int) (*models.ForumsLikesResponse, error) {
	var forumLike models.ForumsLikesResponse

	query := `
        SELECT 
            ForumsLikes.id, ForumsLikes.forum_id, ForumsLikes.created_at, ForumsLikes.updated_at,
            Users.id, Users.username, Users.email, Users.photo
        FROM ForumsLikes 
        JOIN Users ON ForumsLikes.user_id = Users.id
        WHERE user_id = $1 AND forum_id = $2
    `
	row := repository.db.QueryRow(query, user_id, forum_id)
	err := row.Scan(&forumLike.ID, &forumLike.ForumID, &forumLike.CreatedAt, &forumLike.UpdatedAt, &forumLike.User.ID, &forumLike.User.Username, &forumLike.User.Email, &forumLike.User.Photo)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &forumLike, nil
}

func (repository *Repository) CreateLikeByForumId(c *gin.Context, user_id, forum_id int) (*models.ForumsLikesResponse, error) {
	data, err := repository.GetForumById(c, forum_id)
	if data == nil && err == nil {
		return nil, errors.New("Forum not found")
	}

	forumLike, err := repository.GetLikeByUserIDAndForumID(c, user_id, forum_id)

	if err == nil && forumLike != nil {
		return forumLike, nil
	}

	query := `
        INSERT INTO ForumsLikes (user_id, forum_id, created_at, updated_at)
        VALUES ($1, $2, $3, $4)
    `
	result, err := repository.db.Exec(query, user_id, forum_id, time.Now(), time.Now())
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	forumLike = &models.ForumsLikesResponse{
		ID:      int(id),
		ForumID: forum_id,
		User: models.User{
			ID: user_id,
		},
	}
	return forumLike, nil
}

func (repository *Repository) DeleteLikeByForumId(c *gin.Context, user_id, forum_id int) (*models.ForumsLikesResponse, error) {
	forumLike, err := repository.GetLikeByUserIDAndForumID(c, user_id, forum_id)

	if err == nil && forumLike == nil {
		return nil, errors.New("Like not found")
	}

	if err != nil {
		return nil, err
	}

	query := `
        DELETE FROM ForumsLikes WHERE id = $1
    `
	_, err = repository.db.Exec(query, forumLike.ID)
	if err != nil {
		return nil, err
	}

	return forumLike, nil
}

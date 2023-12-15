package repository

import (
	"database/sql"
	"errors"
	"codein/models"
	"time"

	"github.com/gin-gonic/gin"
)

func (repository *Repository) GetAllLikeByBlogID(c *gin.Context, id int) (*[]models.BlogsLikesResponse, error) {
	var blogLikes = make([]models.BlogsLikesResponse, 0)

	query := `
        SELECT 
            BlogsLikes.id, BlogsLikes.blog_id, BlogsLikes.created_at, BlogsLikes.updated_at,
            Users.id, Users.username, Users.email, Users.photo
        FROM BlogsLikes 
        JOIN Users ON BlogsLikes.user_id = Users.id
        WHERE BlogsLikes.blog_id = $1
        ORDER BY BlogsLikes.created_at DESC
    `
	rows, err := repository.db.Query(query, id)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var blogLike models.BlogsLikesResponse
		var user models.User

		err := rows.Scan(&blogLike.ID, &blogLike.BlogID, &blogLike.CreatedAt, &blogLike.UpdatedAt, &user.ID, &user.Username, &user.Email, &user.Photo)
		if err != nil {
			return nil, err
		}
		blogLike.User = user

		blogLikes = append(blogLikes, blogLike)
	}

	return &blogLikes, nil
}

func (repository *Repository) GetLikeByUserIDAndBlogID(c *gin.Context, user_id, blog_id int) (*models.BlogsLikesResponse, error) {
	var blogLike models.BlogsLikesResponse

	query := `
        SELECT 
            BlogsLikes.id, BlogsLikes.blog_id, BlogsLikes.created_at, BlogsLikes.updated_at,
            Users.id, Users.username, Users.email, Users.photo
        FROM BlogsLikes 
        JOIN Users ON BlogsLikes.user_id = Users.id
        WHERE user_id = $1 AND blog_id = $2
    `
	row := repository.db.QueryRow(query, user_id, blog_id)
	err := row.Scan(&blogLike.ID, &blogLike.BlogID, &blogLike.CreatedAt, &blogLike.UpdatedAt, &blogLike.User.ID, &blogLike.User.Username, &blogLike.User.Email, &blogLike.User.Photo)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &blogLike, nil
}

func (repository *Repository) CreateLikeByBlogId(c *gin.Context, user_id, blog_id int) (*models.BlogsLikesResponse, error) {
	data, err := repository.GetBlogByID(c, blog_id)
	if data == nil && err == nil {
		return nil, errors.New("Blog not found")
	}

	blogLike, err := repository.GetLikeByUserIDAndBlogID(c, user_id, blog_id)

	if err == nil && blogLike != nil {
		return blogLike, nil
	}

	query := `
        INSERT INTO BlogsLikes (user_id, blog_id, created_at, updated_at)
        VALUES ($1, $2, $3, $4)
    `
	result, err := repository.db.Exec(query, user_id, blog_id, time.Now(), time.Now())
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	blogLike = &models.BlogsLikesResponse{
		ID:     int(id),
		BlogID: blog_id,
		User: models.User{
			ID: user_id,
		},
	}
	return blogLike, nil
}

func (repository *Repository) DeleteLikeByBlogId(c *gin.Context, user_id, blog_id int) (*models.BlogsLikesResponse, error) {
	blogLike, err := repository.GetLikeByUserIDAndBlogID(c, user_id, blog_id)

	if err == nil && blogLike == nil {
		return nil, errors.New("Like not found")
	}

	if err != nil {
		return nil, err
	}

	query := `
        DELETE FROM BlogsLikes WHERE id = $1
    `
	_, err = repository.db.Exec(query, blogLike.ID)
	if err != nil {
		return nil, err
	}

	return blogLike, nil
}

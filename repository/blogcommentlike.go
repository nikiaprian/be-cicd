package repository

import (
	"database/sql"
	"errors"
	"kel15/models"
	"time"

	"github.com/gin-gonic/gin"
)

func (repository *Repository) GetAllLikeByBlogCommentID(c *gin.Context, id int) (*[]models.BlogCommentLikesResponse, error) {
	var BlogCommentLikes = make([]models.BlogCommentLikesResponse, 0)

	query := `
        SELECT 
            BlogCommentLikes.id, BlogCommentLikes.blog_comment_id, BlogCommentLikes.created_at, BlogCommentLikes.updated_at,
            Users.id, Users.username, Users.email, Users.photo
        FROM BlogCommentLikes 
        JOIN Users ON BlogCommentLikes.user_id = Users.id
        WHERE BlogCommentLikes.blog_comment_id = $1
        ORDER BY BlogCommentLikes.created_at DESC
    `
	rows, err := repository.db.Query(query, id)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var blogCommentLike models.BlogCommentLikesResponse
		var user models.User

		err := rows.Scan(&blogCommentLike.ID, &blogCommentLike.BlogCommentID, &blogCommentLike.CreatedAt, &blogCommentLike.UpdatedAt, &user.ID, &user.Username, &user.Email, &user.Photo)
		if err != nil {
			return nil, err
		}

		blogCommentLike.User = user

		BlogCommentLikes = append(BlogCommentLikes, blogCommentLike)
	}

	return &BlogCommentLikes, nil
}

func (repository *Repository) GetLikeByUserIDAndBlogCommentID(c *gin.Context, user_id, blog_comment_id int) (*models.BlogCommentLikesResponse, error) {
	var blogCommentLike models.BlogCommentLikesResponse

	query := `
        SELECT 
            BlogCommentLikes.id, BlogCommentLikes.blog_comment_id, BlogCommentLikes.created_at, BlogCommentLikes.updated_at,
            Users.id, Users.username, Users.email, Users.photo
        FROM BlogCommentLikes 
        JOIN Users ON BlogCommentLikes.user_id = Users.id
        WHERE user_id = $1 AND blog_comment_id = $2
    `
	row := repository.db.QueryRow(query, user_id, blog_comment_id)
	err := row.Scan(&blogCommentLike.ID, &blogCommentLike.BlogCommentID, &blogCommentLike.CreatedAt, &blogCommentLike.UpdatedAt, &blogCommentLike.User.ID, &blogCommentLike.User.Username, &blogCommentLike.User.Email, &blogCommentLike.User.Photo)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &blogCommentLike, nil
}

func (repository *Repository) CreateLikeByBlogCommentId(c *gin.Context, user_id, blog_comment_id int) (*models.BlogCommentLikesResponse, error) {
	data, err := repository.GetCommentById(c, blog_comment_id)
	if data == nil && err == nil {
		return nil, errors.New("Comment not found")
	}

	blogCommentLike, err := repository.GetLikeByUserIDAndBlogCommentID(c, user_id, blog_comment_id)

	if err == nil && blogCommentLike != nil {
		return blogCommentLike, nil
	}

	query := `
        INSERT INTO BlogCommentLikes (user_id, blog_comment_id, created_at, updated_at)
        VALUES ($1, $2, $3, $4)
    `
	result, err := repository.db.Exec(query, user_id, blog_comment_id, time.Now(), time.Now())
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	blogCommentLike = &models.BlogCommentLikesResponse{
		ID:            int(id),
		BlogCommentID: blog_comment_id,
		User: models.User{
			ID: user_id,
		},
	}
	return blogCommentLike, nil
}

func (repository *Repository) DeleteLikeByBlogCommentId(c *gin.Context, user_id, blog_comment_id int) (*models.BlogCommentLikesResponse, error) {
	blogCommentLike, err := repository.GetLikeByUserIDAndBlogCommentID(c, user_id, blog_comment_id)

	if err == nil && blogCommentLike == nil {
		return nil, errors.New("Like not found")
	}

	if err != nil {
		return nil, err
	}

	query := `
        DELETE FROM BlogCommentLikes WHERE id = $1
    `
	_, err = repository.db.Exec(query, blogCommentLike.ID)
	if err != nil {
		return nil, err
	}

	return blogCommentLike, nil
}

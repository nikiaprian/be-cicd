package repository

import (
	"database/sql"
	"codein/models"
	"time"

	"github.com/gin-gonic/gin"
)

func (repository *Repository) GetCommentBlogById(c *gin.Context, id int) (*models.CommentForum, error) {

	query := `SELECT Comments.id, Comments.comment, Comments.forum_id, Comments.created_at, Comments.updated_at,
			 Users.id, Users.username, Users.email, Users.role, Users.created_at, Users.updated_at, Users.photo
		 	 FROM ComentBlog as Comments 
			JOIN Users ON Comments.user_id = Users.id
			WHERE Comments.id = $1;`

	row := repository.db.QueryRow(query, id)

	var comment models.CommentForum
	var User models.User

	err := row.Scan(&comment.ID, &comment.Comment, &comment.ForumId, &comment.CreatedAt, &comment.UpdatedAt,
		&User.ID, &User.Username, &User.Email, &User.Role, &User.CreatedAt, &User.UpdatedAt, &User.Photo)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	comment.User = User

	return &comment, nil
}

func (repository *Repository) CreateCommentBlog(c *gin.Context, comment string, blog_id, user_id int) (*models.CommentBlog, error) {

	query := `INSERT INTO CommentBlog (comment, blog_id, user_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5);`

	result, err := repository.db.Exec(query, comment, blog_id, user_id, time.Now(), time.Now())
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	user, err := repository.GetUserById(c, int64(user_id))
	if err != nil {
		return nil, err
	}

	return &models.CommentBlog{
		ID:        int(id),
		Comment:   comment,
		User:      *user,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (repository *Repository) GetAllCommentByBlogID(c *gin.Context, id int) ([]models.CommentBlog, error) {
	// var comments []models.CommentBlog

	query := `SELECT Comments.id, Comments.comment, Comments.created_at, Comments.updated_at,
			 Users.id, Users.username, Users.email, Users.role, Users.created_at, Users.updated_at, Users.photo
		 	 FROM CommentBlog as Comments 
			 JOIN Users ON Comments.user_id = Users.id 
			 WHERE Comments.blog_id = $1
			 ORDER BY Comments.id desc;`

	rows, err := repository.db.Query(query, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var comments []models.CommentBlog
	var UserDataLogin models.User
	userLogin, isUser := c.Get("user")

	if isUser == true {
		UserDataLogin = *userLogin.(*models.User)
	}

	for rows.Next() {
		var comment models.CommentBlog
		var User models.User

		err := rows.Scan(&comment.ID, &comment.Comment, &comment.CreatedAt, &comment.UpdatedAt,
			&User.ID, &User.Username, &User.Email, &User.Role, &User.CreatedAt, &User.UpdatedAt, &User.Photo)

		if err != nil {
			return nil, err
		}
		blog_comment_likes, err := repository.GetAllLikeByBlogCommentID(c, comment.ID)
		comment.TotalLikes = len(*blog_comment_likes)

		for _, blog_like := range *blog_comment_likes {
			if isUser == true {
				if blog_like.User.ID == UserDataLogin.ID {
					comment.IsYouLike = true
					break
				}
			}
		}

		comment.BlogCommentLikes = *blog_comment_likes

		comment.User = User
		comments = append(comments, comment)
	}

	return comments, nil
}

func (repository *Repository) DeleteCommentByID(c *gin.Context, id int) error {
	query := `DELETE FROM CommentBlog WHERE id = $1;`

	_, err := repository.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

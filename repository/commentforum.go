package repository

import (
	"database/sql"
	"errors"
	"codein/models"
	"time"

	"github.com/gin-gonic/gin"
)

func (repository *Repository) GetCommentById(c *gin.Context, id int) (*models.CommentForum, error) {

	query := `SELECT Comments.id, Comments.comment, Comments.forum_id, Comments.created_at, Comments.updated_at,
			 Users.id, Users.username, Users.email, Users.role, Users.created_at, Users.updated_at, Users.photo
		 	 FROM CommentForum as Comments 
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

func (repository *Repository) CreateCommentForum(c *gin.Context, comment string, forum_id, user_id int) (*models.CommentForum, error) {

	query := `INSERT INTO CommentForum (comment, forum_id, user_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5);`

	result, err := repository.db.Exec(query, comment, forum_id, user_id, time.Now(), time.Now())
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

	return &models.CommentForum{
		ID:        int(id),
		Comment:   comment,
		User:      *user,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (repository *Repository) GetAllCommentByForumID(c *gin.Context, id int) ([]models.CommentForum, error) {
	var comments = make([]models.CommentForum, 0)
	query := `SELECT Comments.id, Comments.comment, Comments.is_answer, Comments.created_at, Comments.updated_at,
			 Users.id, Users.username, Users.email, Users.role, Users.created_at, Users.updated_at, Users.photo
		 	 FROM CommentForum as Comments 
			JOIN Users ON Comments.user_id = Users.id
			WHERE Comments.forum_id = $1
			ORDER BY Comments.is_answer desc, Comments.created_at DESC;`

	rows, err := repository.db.Query(query, id)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	defer rows.Close()

	var UserDataLogin models.User
	userLogin, isUser := c.Get("user")

	if isUser == true {
		UserDataLogin = *userLogin.(*models.User)
	}

	for rows.Next() {
		var User models.User
		var comment models.CommentForum

		err := rows.Scan(&comment.ID, &comment.Comment, &comment.IsAnswer, &comment.CreatedAt, &comment.UpdatedAt,
			&User.ID, &User.Username, &User.Email, &User.Role, &User.CreatedAt, &User.UpdatedAt, &User.Photo)

		if err != nil {
			return nil, err
		}

		forum_comment_likes, err := repository.GetAllLikeByForumCommentID(c, comment.ID)
		comment.TotalLikes = len(*forum_comment_likes)

		for _, forum_like := range *forum_comment_likes {
			if isUser == true {
				if forum_like.User.ID == UserDataLogin.ID {
					comment.IsYouLike = true
					break
				}
			}
		}

		comment.ForumCommentLikes = *forum_comment_likes

		comment.User = User
		comments = append(comments, comment)
	}

	return comments, nil
}

func (repository *Repository) DeleteCommentForum(c *gin.Context, id int) error {

	query := `DELETE FROM CommentForum WHERE id = $1;`

	_, err := repository.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (repository *Repository) SelectedCommentAnswer(c *gin.Context, id int, is_answer bool) (*models.CommentForum, error) {
	comment, err := repository.GetCommentById(c, id)
	if comment == nil && err == nil {
		return nil, errors.New("Comment not found")
	}
	if err != nil {
		return nil, err
	}

	forum, err := repository.GetForumById(c, *comment.ForumId)

	if forum == nil && err == nil {
		return nil, errors.New("Forum not found")
	}

	if err != nil {
		return nil, err
	}

	user_id := c.MustGet("user").(*models.User).ID

	if forum.User.ID != user_id {
		return nil, errors.New("You are not allowed to do this")
	}

	query := `UPDATE CommentForum SET is_answer = $1 WHERE id = $2;`

	_, err = repository.db.Exec(query, is_answer, id)
	if err != nil {
		return nil, err
	}

	comment.IsAnswer = is_answer
	return comment, nil
}

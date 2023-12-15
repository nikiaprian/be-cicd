package usecase

import (
	"encoding/json"
	"errors"
	"codein/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (usecase *Usecase) CreateCommentForum(c *gin.Context) (*models.CommentForum, error) {
	user, _ := c.Get("user")
	userData := user.(*models.User)

	if user == nil {
		return nil, errors.New("user not found")
	}

	forum_id := c.Param("id")
	Convid, err := strconv.Atoi(forum_id)

	var commentForum models.CommentForumRequest
	err = c.ShouldBindJSON(&commentForum)
	if err != nil {
		return nil, err
	}

	var comment string = commentForum.Comment

	commentForumResponse, err := usecase.repository.CreateCommentForum(c, comment, int(Convid), userData.ID)

	if err != nil {
		return nil, err
	}

	return commentForumResponse, nil
}

func (usecase *Usecase) GetAllCommentByForumID(c *gin.Context) ([]models.CommentForum, error) {
	id := c.Param("id")
	Convid, err := strconv.Atoi(id)

	if err != nil {
		return nil, err
	}

	comments, err := usecase.repository.GetAllCommentByForumID(c, Convid)
	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (usecase *Usecase) DeleteCommentForum(c *gin.Context) error {
	id := c.Param("id")
	Convid, err := strconv.Atoi(id)

	if err != nil {
		return err
	}

	err = usecase.repository.DeleteCommentForum(c, Convid)
	if err != nil {
		return err
	}

	return nil
}

func (usecase *Usecase) SelectedCommentAnswer(c *gin.Context) (*models.CommentForum, error) {
	id := c.Param("id")
	Convid, err := strconv.Atoi(id)
	var body models.SelectedCommentAnswerRequest

	err = json.NewDecoder(c.Request.Body).Decode(&body)

	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	comment, err := usecase.repository.SelectedCommentAnswer(c, Convid, body.IsAnswer)
	if err != nil {
		return nil, err
	}

	return comment, nil
}

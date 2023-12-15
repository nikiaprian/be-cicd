package usecase

import (
	"errors"
	"codein/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (usecase *Usecase) CreateCommentBlog(c *gin.Context) (*models.CommentBlog, error) {
	// var User *models.User
	// user := c.Get("user").(User)
	user, _ := c.Get("user")
	userData := user.(*models.User)

	if user == nil {
		return nil, errors.New("user not found")
	}

	blog_id := c.Param("id")
	Convid, err := strconv.Atoi(blog_id)

	var commentBlog models.CommentBlogRequest
	err = c.ShouldBindJSON(&commentBlog)
	if err != nil {
		return nil, err
	}

	var comment string = commentBlog.Comment

	commentBlogResponse, err := usecase.repository.CreateCommentBlog(c, comment, int(Convid), userData.ID)
	

	if err != nil {
		return nil, err
	}

	return commentBlogResponse, nil
}

func (usecase *Usecase) GetAllCommentByBlogID(c *gin.Context) ([]models.CommentBlog, error) {
	id := c.Param("id")
	Convid, err := strconv.Atoi(id) 
	if err != nil {
		return nil, err
	}

	comments, err := usecase.repository.GetAllCommentByBlogID(c, Convid)
	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (usecase *Usecase) DeleteCommentByID(c *gin.Context) error {
	id := c.Param("id")
	Convid, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	err = usecase.repository.DeleteCommentByID(c, Convid)
	if err != nil {
		return err
	}

	return nil
}
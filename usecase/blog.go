package usecase

import (
	"errors"
	"fmt"
	"codein/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (usecase *Usecase) GetAllBlog(c *gin.Context) ([]models.Blog, error) {
	blogs, err := usecase.repository.GetAllBlog(c)
	if err != nil {
		return nil, err
	}

	return blogs, nil
}
func (usecase *Usecase) CreateBlog(c *gin.Context) (*models.Blog, error) {
	user, _ := c.Get("user")
	userData := user.(*models.User)
	if user == nil {
		return nil, errors.New("user not found")
	}

	fileName, _ := c.Get("file")
	c.Request.ParseForm()
	tags, _ := c.Request.PostForm["tags"]


	var payload models.BlogRequest
	payload.Content = c.Request.FormValue("content")
	payload.Title = c.Request.FormValue("title")
	// payload.Tags = a

	blog, err := usecase.repository.CreateBlog(c, payload, fileName.(string), userData.ID)
	if err != nil {
		return nil, err
	}
fmt.Println(len(tags), tags)
	for _, tag := range tags {
		Tag, _ := usecase.repository.CreateTag(c, tag)
		if Tag != nil {
			blogtag, err := usecase.repository.CreateBlogTag(c, int64(blog.ID), int64(Tag.ID))
			fmt.Println(err!=nil)
			if err != nil {
				continue
			}
			blog.Tags = append(blog.Tags, *blogtag)
		}
	}

	return blog, nil
}

// func (usecase *Usecase) UpdateBlog(c *gin.Context) (*models.Blog, error) {
// 	user, _ := c.Get("user")
// 	userData := user.(*models.User)
// 	if user == nil {
// 		return nil, errors.New("user not found")
// 	}

// 	i := c.Param("id")
// 	id, err := strconv.Atoi(i)
// 	if err != nil {
// 		return nil, err
// 	}

// 	fileName, _ := c.Get("file")

// 	var payload models.BlogRequest
// 	payload.Content = c.Request.FormValue("content")
// 	payload.Title = c.Request.FormValue("title")

// 	blog, err := usecase.repository.UpdateBlog(c, payload, id, fileName.(string), userData.ID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return blog, nil
	// var payload models.BlogRequest
	// err := c.BindJSON(&payload)
	// i := c.Param("id")
	// id, _ := strconv.Atoi(i)
	// if err != nil {
	// 	return nil, err
	// }

	// blog, err := usecase.repository.UpdateBlog(c, payload, id, "")
	// if err != nil {
	// 	return nil, err
	// }

	// return blog, nil
// }

func (usecase *Usecase) DeleteBlog(c *gin.Context) error {
	i := c.Param("id")
	id, err := strconv.Atoi(i)
	if err != nil {
		return err
	}

	err = usecase.repository.DeleteBlog(c, id)
	if err != nil {
		return err
	}

	return nil
}

func (usecase *Usecase) GetBlogByID(c *gin.Context) (*models.Blog, error) {
	i := c.Param("id")
	id, _ := strconv.Atoi(i)
	blog, err := usecase.repository.GetBlogByID(c, id)
	if err != nil {
		return nil, err
	}

	return blog, nil
}

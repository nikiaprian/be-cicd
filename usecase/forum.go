package usecase

import (
	"errors"
	"kel15/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (usecase *Usecase) GetAllForum(c *gin.Context) ([]models.Forum, error) {
	forums, err := usecase.repository.GetAllForum(c)
	if err != nil {
		return nil, err
	}

	return forums, nil
}

func (usecase *Usecase) CreateForum(c *gin.Context) (*models.Forum, error) {
	user, _ := c.Get("user")
	userData := user.(*models.User)

	if user == nil {
		return nil, errors.New("user not found")
	}

	// tags, _ := c.Request.PostForm["tags"]

	var payload models.ForumRequest
	payload.Title = c.Request.FormValue("title")
	payload.Content = c.Request.FormValue("content")
	payload.Tags = c.Request.PostForm["tags"]
	
	forum, err := usecase.repository.CreateForum(c, payload.Title, payload.Content, userData.ID)

	if err != nil {
		return nil, err
	}

	for _, tag := range payload.Tags {
		Tag, _ := usecase.repository.CreateTag(c, tag)
		if Tag != nil {
			forumtag, err := usecase.repository.CreateForumTag(c, int64(forum.ID), int64(Tag.ID))
			if err != nil {
				continue
			}
			forum.Tags = append(forum.Tags, *forumtag)
		}
	}
	
	return forum, nil

}

// func (usecase *Usecase) UpdateForum(c *gin.Context) (*models.Forum, error) {
// 	user := c.MustGet("user").(*models.User)

// 	if user == nil {
// 		return nil, errors.New("user not found")
// 	}

// 	id := c.Param("id")
// 	Convid, err := strconv.Atoi(id)

// 	if err != nil {
// 		return nil, err
// 	}

// 	var forum models.ForumRequest

// 	err = c.BindJSON(&forum)

// 	if err != nil {
// 		return nil, err
// 	}

// 	forumResponse, err := usecase.repository.UpdateForum(c, Convid, forum.Title, forum.Content)

// 	if err != nil {
// 		return nil, err
// 	}

// 	for _, tag := range forum.Tags {
// 		Tag, _ := usecase.repository.CreateTag(c, tag)
// 		if Tag != nil {
// 			usecase.repository.CreateForumTag(c, int64(forumResponse.ID), int64(Tag.ID))
// 		}
// 	}

// 	return forumResponse, nil
// }

func (usecase *Usecase) DeleteForum(c *gin.Context) (*models.Forum, error) {
	user := c.MustGet("user").(*models.User)
	if user == nil {
		return nil, errors.New("user not found")
	}

	id := c.Param("id")
	Convid, err := strconv.Atoi(id)

	if err != nil {
		return nil, err
	}

	err = usecase.repository.DeleteForum(c, Convid)

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (usecase *Usecase) GetForumById(c *gin.Context) (*models.Forum, error) {
	id := c.Param("id")
	Convid, _ := strconv.Atoi(id)
	forum, err := usecase.repository.GetForumById(c, Convid)
	if err != nil {
		return nil, err
	}

	return forum, nil
}

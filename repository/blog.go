package repository

import (
	"database/sql"
	"codein/models"
	"time"

	"github.com/gin-gonic/gin"
)

func (repository *Repository) GetAllBlog(c *gin.Context) ([]models.Blog, error) {

	query := `SELECT Blogs.id, Blogs.photo, Blogs.title, Blogs.content, Blogs.created_at, Blogs.updated_at,
			  Users.id, Users.username, Users.email, Users.role, Users.created_at, Users.updated_at, Users.photo
			  FROM Blogs 
			  JOIN Users ON Blogs.user_id = Users.id
			  ORDER BY Blogs.id DESC`

	rows, err := repository.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var blogs []models.Blog
	var User models.User

	var UserDataLogin models.User
	userLogin, isUser := c.Get("user")

	if isUser == true {
		UserDataLogin = *userLogin.(*models.User)
	}

	for rows.Next() {
		var blog models.Blog
		err := rows.Scan(&blog.ID, &blog.Photo, &blog.Title, &blog.Content, &blog.CreatedAt, &blog.UpdatedAt,
			&User.ID, &User.Username, &User.Email, &User.Role, &User.CreatedAt, &User.UpdatedAt, &User.Photo)
		if err != nil {
			return nil, err
		}
		blog.User = User

		blog_tags, err := repository.GetBlogTagByBlogID(c, int64(blog.ID))
		if err != nil {
			continue
		}

		for _, blog_tag := range *blog_tags {
			tag, err := repository.GetTagByID(c, int64(blog_tag.TagID))
			if err != nil {
				continue
			}
			blog.Tags = append(blog.Tags, *tag)
		}

		blog_likes, _ := repository.GetAllLikeByBlogID(c, int(blog.ID))
		blog.TotalLikes = len(*blog_likes)

		for _, forum_like := range *blog_likes {
			if isUser == true {
				if forum_like.User.ID == UserDataLogin.ID {
					blog.IsYouLike = true
					break
				}
			}
		}

		blog.BlogsLikes = *blog_likes

		blog_comments, _ := repository.GetAllCommentByBlogID(c, int(blog.ID))
		blog.TotalComment = len(blog_comments)

		blogs = append(blogs, blog)
	}

	return blogs, nil
}

func (repository *Repository) CreateBlog(c *gin.Context, req models.BlogRequest, photo string, user_id int) (*models.Blog, error) {
	title := req.Title
	content := req.Content

	query := "INSERT INTO Blogs (photo, user_id, title, content, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)"
	row1, err := repository.db.Exec(query, photo, user_id, title, content, time.Now(), time.Now())
	if err != nil {
		return nil, err
	}
	id, err := row1.LastInsertId()
	if err != nil {
		return nil, err
	}

	user, err := repository.GetUserById(c, int64(user_id))
	if err != nil {
		return nil, err
	}

	return &models.Blog{
		ID:        uint(id),
		User:      *user,
		Photo:     photo,
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

// func (repository *Repository) UpdateBlog(c *gin.Context, req models.BlogRequest, id int, photo string, user_id int) (*models.Blog, error) {
// 	title := req.Title
// 	content := req.Content

// 	query := `UPDATE Blogs INNER JOIN Users ON Blogs.user_id = Users.id
// 			  SET photo = ?, user_id = ?, title = ?, content = ?, created_at = ?, updated_at = ?
// 			  WHERE Blogs.id = ?`

// 	_, err := repository.db.Exec(query, photo, user_id, title, content, time.Now(), time.Now(), id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	user, err := repository.GetUserById(c, int64(id))
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &models.Blog{
// 		ID:        uint(id),
// 		User:      *user,
// 		Photo:     photo,
// 		Title:     title,
// 		Content:   content,
// 		CreatedAt: time.Now(),
// 		UpdatedAt: time.Now(),
// 	}, nil
// }

func (repository *Repository) DeleteBlog(c *gin.Context, id int) error {
	query := "DELETE from Blogs where id=?"
	_, err := repository.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (repository *Repository) GetBlogByID(c *gin.Context, id int) (*models.Blog, error) {
	var blog models.Blog
	var User models.User

	query := `SELECT Blogs.id, Blogs.photo, Blogs.title, Blogs.content, Blogs.created_at, Blogs.updated_at,
			  Users.id, Users.username, Users.email, Users.role, Users.created_at, Users.updated_at, Users.photo
			  FROM Blogs
			  JOIN Users ON Blogs.user_id = Users.id
			  WHERE Blogs.id = $1`

	row := repository.db.QueryRow(query, id)
	err := row.Scan(&blog.ID, &blog.Photo, &blog.Title, &blog.Content, &blog.CreatedAt, &blog.UpdatedAt,
		&User.ID, &User.Username, &User.Email, &User.Role, &User.CreatedAt, &User.UpdatedAt, &User.Photo)

	if err != nil && err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	blog_tags, _ := repository.GetBlogTagByBlogID(c, int64(blog.ID))
	for _, blog_tag := range *blog_tags {
		tag, err := repository.GetTagByID(c, int64(blog_tag.TagID))
		if err != nil {
			continue
		}
		blog.Tags = append(blog.Tags, *tag)
	}

	var UserDataLogin models.User
	userLogin, isUser := c.Get("user")

	if isUser == true {
		UserDataLogin = *userLogin.(*models.User)
	}

	blog_likes, _ := repository.GetAllLikeByBlogID(c, int(blog.ID))
	blog.TotalLikes = len(*blog_likes)
	blog.BlogsLikes = *blog_likes

	for _, forum_like := range *blog_likes {
		if isUser == true {
			if forum_like.User.ID == UserDataLogin.ID {
				blog.IsYouLike = true
				break
			}
		}
	}

	blog.User = User

	blog_comments, _ := repository.GetAllCommentByBlogID(c, int(blog.ID))
	blog.TotalComment = len(blog_comments)

	return &blog, nil
}

package repository

import (
	"database/sql"
	"kel15/models"
	"time"

	"github.com/gin-gonic/gin"
)

func (repository *Repository) GetAllForum(c *gin.Context) ([]models.Forum, error) {
	query := `SELECT Forums.id, Forums.title, Forums.content, Forums.created_at, Forums.updated_at,
			  Users.id, Users.username, Users.email, Users.role, Users.created_at, Users.updated_at, Users.photo
			  FROM Forums 
			  JOIN Users ON Forums.user_id = Users.id
			  ORDER BY Forums.id DESC
			  `

	rows, err := repository.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var forums []models.Forum
	var UserForum models.User

	var UserDataLogin models.User
	userLogin, isUser := c.Get("user")

	if isUser == true {
		UserDataLogin = *userLogin.(*models.User)
	}

	for rows.Next() {
		var forum models.Forum
		forum.IsYouLike = false

		err := rows.Scan(&forum.ID, &forum.Title, &forum.Content, &forum.CreatedAt, &forum.UpdatedAt,
			&UserForum.ID, &UserForum.Username, &UserForum.Email, &UserForum.Role, &UserForum.CreatedAt, &UserForum.UpdatedAt, &UserForum.Photo)

		if err != nil {
			return nil, err
		}

		forum.User = UserForum

		forum_tags, _ := repository.GetForumTagByForumID(c, int64(forum.ID))

		for _, forum_tag := range *forum_tags {
			tag, err := repository.GetTagByID(c, int64(forum_tag.TagID))
			if err != nil {
				continue
			}

			forum.Tags = append(forum.Tags, *tag)
		}

		forum_likes, _ := repository.GetAllLikeByForumID(c, forum.ID)
		forum.TotalLikes = len(*forum_likes)

		for _, forum_like := range *forum_likes {
			if isUser == true {
				if forum_like.User.ID == UserDataLogin.ID {
					forum.IsYouLike = true
					break
				}
			}
		}

		forum.ForumsLikes = *forum_likes
                forum_comments, _ := repository.GetAllCommentByForumID(c, forum.ID)
                forum.TotalComment = len(forum_comments)
		forums = append(forums, forum)
	}

	return forums, nil
}

func (repository *Repository) CreateForum(c *gin.Context, title, content string, user_id int) (*models.Forum, error) {
	query := `INSERT INTO Forums (title, content, user_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)`

	result, err := repository.db.Exec(query, title, content, user_id, time.Now(), time.Now())
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

	return &models.Forum{
		ID:        int(id),
		User:      *user,
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

// func (repository *Repository) UpdateForum(c *gin.Context, id int, title, contents string) (*models.Forum, error) {
// 	query := `UPDATE Forums SET title = ?, content = ?, updated_at = ? WHERE id = ?`

// 	_, err := repository.db.Exec(query, title, contents, time.Now(), id)

// 	if err != nil {
// 		return nil, err
// 	}

// 	user, err := repository.GetUserById(c, int64(id))
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &models.Forum{
// 		ID:        id,
// 		User:      *user,
// 		Title:     title,
// 		Content:   contents,
// 		CreatedAt: time.Now(),
// 		UpdatedAt: time.Now(),
// 	}, nil
// }

func (repository *Repository) DeleteForum(c *gin.Context, id int) error {
	query := `DELETE FROM Forums WHERE id = $1`

	_, err := repository.db.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}

func (repository *Repository) GetForumById(c *gin.Context, id int) (*models.Forum, error) {
	var forum models.Forum
	var User models.User

	query := `SELECT Forums.id, Forums.title, Forums.content, Forums.created_at, Forums.updated_at,
			  Users.id, Users.username, Users.email, Users.role, Users.created_at, Users.updated_at, Users.photo
			  FROM Forums
			  JOIN Users ON Forums.user_id = Users.id
			  WHERE Forums.id = $1`

	row := repository.db.QueryRow(query, id)
	err := row.Scan(&forum.ID, &forum.Title, &forum.Content, &forum.CreatedAt, &forum.UpdatedAt,
		&User.ID, &User.Username, &User.Email, &User.Role, &User.CreatedAt, &User.UpdatedAt, &User.Photo)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	forum_likes, _ := repository.GetAllLikeByForumID(c, forum.ID)
	forum.TotalLikes = len(*forum_likes)
	forum.ForumsLikes = *forum_likes
	var UserDataLogin models.User
	userLogin, isUser := c.Get("user")

	if isUser == true {
		UserDataLogin = *userLogin.(*models.User)
	}

	for _, forum_like := range *forum_likes {
		if forum_like.User.ID == UserDataLogin.ID {
			forum.IsYouLike = true
			break
		}
	}

	forum_tags, _ := repository.GetForumTagByForumID(c, int64(forum.ID))
	for _, forum_tag := range *forum_tags {
		tag, err := repository.GetTagByID(c, int64(forum_tag.TagID))
		if err != nil {
			continue
		}

		forum.Tags = append(forum.Tags, *tag)
	}

	forum.User = User

	forum_comments, _ := repository.GetAllCommentByForumID(c, forum.ID)
	forum.TotalComment = len(forum_comments)

	return &forum, nil
}

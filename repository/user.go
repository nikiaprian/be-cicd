package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"codein/models"
	"codein/utils"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func (repository *Repository) UserTotal(c *gin.Context, pagination models.Pagination) int {
	var total int
	search := pagination.Search
	query := `
		SELECT COUNT(*) as total FROM Users 
		WHERE email LIKE $1 OR username LIKE $2
	`
	err := repository.db.QueryRow(query, search, search).Scan(&total)
	fmt.Println(err, total)

	return total
}

func (repository *Repository) UserList(c *gin.Context, pagination models.Pagination) ([]models.User, error) {
	var users []models.User

	search := pagination.Search
	pagination.Count = repository.UserTotal(c, pagination)

	pagination = utils.SetTotalPagePagination(c, pagination)

	utils.SetPaginationNew(c, pagination)

	offset := (pagination.PageActive - 1) * pagination.Limit
	query := `
		SELECT id, email, username, photo, created_at, updated_at FROM Users 
		WHERE email LIKE $1 OR username LIKE $2
		ORDER BY created_at DESC
		LIMIT $3
		OFFSET $4
	`
	rows, err := repository.db.Query(query, search, search, pagination.Limit, offset)
	// defer rows.Close()

	if err != nil && err != sql.ErrNoRows {
		return users, err
	}

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Email, &user.Username, &user.Photo, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (repository *Repository) GetUserById(c *gin.Context, id int64) (*models.User, error) {
	query := `SELECT id, email, username, photo, created_at, updated_at FROM Users WHERE id = $1 limit 1`
	row := repository.db.QueryRow(query, id)

	if row.Err() != nil {
		return nil, row.Err()
	}

	var user models.User
	err := row.Scan(&user.ID, &user.Email, &user.Username, &user.Photo, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return &user, errors.New("User not found")
		}
		return nil, err
	}

	return &user, nil
}

func (repository *Repository) GetUserByGoogleId(c *gin.Context, google_id string) (*models.User, error) {
	query := `SELECT Users.id, email, username, photo, Users.created_at, Users.updated_at FROM Users JOIN GoogleAccounts ON Users.id = GoogleAccounts.user_id WHERE GoogleAccounts.google_id = $1 limit 1`
	row := repository.db.QueryRow(query, google_id)

	if row.Err() != nil {
		return nil, row.Err()
	}

	var user models.User
	err := row.Scan(&user.ID, &user.Email, &user.Username, &user.Photo, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("User not found")
		}
		return nil, err
	}

	return &user, nil
}

func (repository *Repository) GetUserByEmail(c *gin.Context, email string) (*models.User, error) {
	query := `SELECT id, email, username, password, created_at, updated_at FROM Users WHERE email = $1 limit 1`
	row := repository.db.QueryRow(query, email)

	if row.Err() != nil {
		return nil, row.Err()
	}

	var user models.User
	err := row.Scan(&user.ID, &user.Email, &user.Username, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return &user, errors.New("User not found")
		}
		return nil, err
	}

	return &user, nil
}

func (repository *Repository) GetUserByUsername(c *gin.Context, username string) (*models.User, error) {
	query := `SELECT id, email, password, created_at, updated_at FROM Users WHERE username = $1 limit 1`
	row := repository.db.QueryRow(query, username)

	var user models.User
	err := row.Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repository *Repository) GetUserByEmailOrUsername(c *gin.Context, req models.UserRegisterRequest) (*models.User, error) {
	email := req.Email
	username := req.Username
	query := `SELECT id, email, password, created_at, updated_at FROM Users WHERE email = $1 or username = $2 limit 1`
	row := repository.db.QueryRow(query, email, username)

	var user models.User
	err := row.Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repository *Repository) CreateUser(c *gin.Context, req models.UserRegisterRequest) error {
	email := req.Email
	username := req.Username
	password, err := utils.GeneratePassword(req.Password)
	if err != nil {
		return err
	}

	query := `INSERT INTO Users (email, username, password, provider, created_at, updated_at) VALUES ($1, $2, $3, 'local', $4, $5)`
	_, err = repository.db.Exec(query, email, username, password, time.Now(), time.Now())

	if err != nil {
		return err
	}

	return nil
}

func (repository *Repository) CreateUserGoogle(c *gin.Context, req models.UserGoogleResponse) (*models.User, error) {
	email := req.Email
	googleId := req.GoogleId
	username := strings.Split(email, "@")[0]

	user, err := repository.GetUserByGoogleId(c, googleId)

	if user != nil {
		return user, nil
	}

	if err != nil && user != nil {
		return nil, err
	}

	query := `INSERT INTO Users (email, username, provider, created_at, updated_at) VALUES ($1, $2, 'google', $3, $4)`

	data, err := repository.db.Exec(query, email, username, time.Now(), time.Now())

	if err != nil {
		return nil, err
	}

	id, err := data.LastInsertId()

	query = `INSERT INTO GoogleAccounts (user_id, google_id, created_at, updated_at) VALUES ($1, $2, $3, $4)`

	data, err = repository.db.Exec(query, id, googleId, time.Now(), time.Now())

	if err != nil {
		return nil, err
	}

	user, err = repository.GetUserById(c, id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repository *Repository) UserProfileUpdate(c *gin.Context, req models.UserUpdateProfileRequest) (*models.User, error) {
	fileName := req.FileName
	user := c.MustGet("user").(*models.User)

	if fileName == nil || *fileName == "" {
		req.FileName = user.Photo
	}

	query := `UPDATE Users SET  username = $1, photo = $2, updated_at = $3 WHERE id = $4`
	_, err := repository.db.Exec(query, req.Username, req.FileName, time.Now(), user.ID)

	if err != nil {
		return nil, err
	}

	user.Username = req.Username
	user.Photo = req.FileName

	return user, nil
}

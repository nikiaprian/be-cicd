package models

import "time"

type ForumsLikes struct {
	ID        int        `json:"id"`
	User      User       `json:"user"`
	Forum     Forum      `json:"forum"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type ForumsLikesResponse struct {
	ID        int        `json:"id"`
	ForumID   int        `json:"forum_id"`
	User      User       `json:"user"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

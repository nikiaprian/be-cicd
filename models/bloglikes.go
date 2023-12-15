package models

import "time"

type BlogsLikes struct {
	ID        int        `json:"id"`
	User      User       `json:"user"`
	Blog      Blog       `json:"blog"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type BlogsLikesResponse struct {
	ID        int        `json:"id"`
	BlogID    int        `json:"blog_id"`
	User      User       `json:"user"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

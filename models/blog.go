package models

import "time"

type Blog struct {
	ID           uint                 `json:"id"`
	User         User                 `json:"user"`
	Photo        string               `json:"photo"`
	Title        string               `json:"title"`
	Content      string               `json:"content"`
	Tags         []Tag                `json:"tag"`
	BlogsLikes   []BlogsLikesResponse `json:"blog_likes"`
	TotalLikes   int                  `json:"total_likes"`
	IsYouLike    bool                 `json:"is_you_like"`
	TotalComment int                  `json:"total_comment"`
	CreatedAt    time.Time            `json:"created_at"`
	UpdatedAt    time.Time            `json:"updated_at"`
}

type BlogRequest struct {
	Title   string `json:"title"`
	Tags    []Tag  `json:"tag"`
	Content string `json:"content"`
}

type BlogResponse struct {
	Blog    []Blog `json:"blog"`
	Message string `json:"message"`
}

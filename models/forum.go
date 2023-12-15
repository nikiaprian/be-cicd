package models

import "time"

type Forum struct {
	ID           int                   `json:"id"`
	User         User                  `json:"user"`
	Title        string                `json:"title"`
	Content      string                `json:"content"`
	Tags         []Tag                 `json:"tag"`
	ForumsLikes  []ForumsLikesResponse `json:"forum_likes"`
	TotalLikes   int                   `json:"total_likes"`
	TotalComment int                   `json:"total_comment"`
	IsYouLike    bool                  `json:"is_you_like"`
	CreatedAt    time.Time             `json:"created_at"`
	UpdatedAt    time.Time             `json:"updated_at"`
}

type ForumRequest struct {
	Title   string   `json:"title"`
	Tags    []string `json:"tags"`
	Content string   `json:"content"`
}

type ForumResponse struct {
	Forum   []Forum `json:"forum"`
	Message string  `json:"message"`
}

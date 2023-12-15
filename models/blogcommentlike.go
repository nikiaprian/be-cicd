package models

import "time"

type BlogCommentLikes struct {
	ID          int         `json:"id"`
	User        User        `json:"user"`
	CommentBlog CommentBlog `json:"comment"`
	CreatedAt   *time.Time  `json:"created_at"`
	UpdatedAt   *time.Time  `json:"updated_at"`
}

type BlogCommentLikesResponse struct {
	ID            int        `json:"id"`
	BlogCommentID int        `json:"blog_comment_id"`
	User          User       `json:"user"`
	CreatedAt     *time.Time `json:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at"`
}

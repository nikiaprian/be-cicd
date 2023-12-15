package models

import "time"

type ForumCommentLikes struct {
	ID           int          `json:"id"`
	User         User         `json:"user"`
	CommentForum CommentForum `json:"comment"`
	CreatedAt    *time.Time   `json:"created_at"`
	UpdatedAt    *time.Time   `json:"updated_at"`
}

type ForumCommentLikesResponse struct {
	ID             int        `json:"id"`
	ForumCommentID int        `json:"forum_comment_id"`
	User           User       `json:"user"`
	CreatedAt      *time.Time `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at"`
}

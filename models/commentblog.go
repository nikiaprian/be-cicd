package models

import "time"

type CommentBlog struct {
	ID               int                        `json:"id"`
	User             User                       `json:"user"`
	Comment          string                     `json:"comment" `
	BlogId           *int                       `json:"blog_id,omitempty"`
	BlogCommentLikes []BlogCommentLikesResponse `json:"blog__comment_likes"`
	TotalLikes       int                        `json:"total_likes"`
	IsYouLike        bool                       `json:"is_you_like"`
	CreatedAt        time.Time                  `json:"created_at" `
	UpdatedAt        time.Time                  `json:"updated_at"`
}

type CommentBlogRequest struct {
	// BlogID    int 		`json:"blog_id"`
	Comment string `json:"comment"`
}

type CommentBlogResponse struct {
	CommentBlog []CommentBlog `json:"comment_blog"`
	Message     string        `json:"message"`
}

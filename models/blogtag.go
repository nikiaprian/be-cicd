package models

type BlogTag struct {
	ID     int   `json:"id"`
	BlogID int64 `json:"blog_id"`
	TagID  int64 `json:"tag_id"`
}
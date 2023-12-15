package models

type Forumtag struct {
	ID      int   `json:"id"`
	ForumID int64 `json:"forum_id"`
	TagID   int64 `json:"tag_id"`
}

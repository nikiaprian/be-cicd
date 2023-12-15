package models

import "time"

type Tag struct {
	ID        int       `json:"id"`
	Tag       string    `json:"tag"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type TagRequest struct {
	Tag string `json:"tag"`
}
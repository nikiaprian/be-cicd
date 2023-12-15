package models

type Pagination struct {
	PageActive int    `json:"page"`
	Limit      int    `json:"limit"`
	Count      int    `json:"count"`
	TotalPage  int    `json:"total_page"`
	Search     string `json:"search"`
}

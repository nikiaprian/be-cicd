package handler

import (
	"kel15/models"
	"kel15/project"
)

type Handler struct {
	Project *project.Project
}

func NewHandler(project *project.Project) *Handler {
	return &Handler{
		Project: project,
	}
}

type sendResponseError struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type sendResponseSuccess struct {
	Success    bool               `json:"success"`
	Code       int                `json:"code"`
	Message    string             `json:"message,omitempty"`
	Data       interface{}        `json:"data"`
	Pagination *models.Pagination `json:"pagination,omitempty"`
}

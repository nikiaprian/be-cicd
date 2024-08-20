package project

import (
	"codein/storage"
	"codein/usecase"
)

type Project struct {
	Usecase *usecase.Usecase
	Storage *storage.StorageS3Stuct
}

type ProjectInterface interface {
    Usecase() UsecaseInterface // Pastikan ini sesuai dengan struktur Anda
    Storage() StorageInterface   // Jika ada Storage
}

func NewProject(usecase *usecase.Usecase, sto storage.StorageS3Stuct) *Project {
	return &Project{Usecase: usecase, Storage: &sto}
}

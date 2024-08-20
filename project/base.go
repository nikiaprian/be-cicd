package project

import (
	"codein/storage"
	"codein/usecase"
)

type Project struct {
	Usecase *usecase.Usecase
	Storage *storage.StorageS3Stuct
}

func NewProject(usecase *usecase.Usecase, sto storage.StorageS3Stuct) *Project {
	return &Project{Usecase: usecase, Storage: &sto}
}

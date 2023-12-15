package project

import (
	"kel15/storage"
	"kel15/usecase"
)

type Project struct {
	Usecase *usecase.Usecase
	Storage *storage.StorageS3Stuct
}

func NewProject(usecase *usecase.Usecase, sto storage.StorageS3Stuct) *Project {
	return &Project{Usecase: usecase, Storage: &sto}
}

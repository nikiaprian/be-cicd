package usecase

import "kel15/repository"

type Usecase struct {
	repository *repository.Repository
}

func NewUsecase(repository *repository.Repository) *Usecase {
	return &Usecase{repository: repository}
}

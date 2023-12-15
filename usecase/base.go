package usecase

import "codein/repository"

type Usecase struct {
	repository *repository.Repository
}

func NewUsecase(repository *repository.Repository) *Usecase {
	return &Usecase{repository: repository}
}

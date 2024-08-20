// mock_project.go

package handler

import (
	"codein/project"
	"codein/usecase"
	"codein/storage"
)

type MockProject struct {
	Usecase *MockUsecase
	Storage *MockStorage
}

// Implementasi method yang sesuai dengan interface project.Project
func (m *MockProject) Usecase() *usecase.Usecase {
	return m.Usecase
}

func (m *MockProject) Storage() *storage.StorageS3Stuct {
	return m.Storage
}

package desenvolvedor

import (
	"backend/internal/repository"
	"backend/internal/usecase"
)

type desenvolvedorUseCase struct {
	desenvolvedorRepository repository.DesenvolvedorRepository
}

const (
	// Error Messages
	errorSavingData           string = "error saving data"
	errorListingData          string = "error listing data"
	errorDeletingData         string = "error deleting data"
	invalidParameter          string = "invalid parameter"
	invalidRequestBody        string = "invalid request body"
	invalidDataShouldAdequate string = "invalid data should be adequate"
	developerNotFound         string = "developer not found"
	nothingToList             string = "nothing to list"
)

func NewUseCase(desenvolvedorRepository repository.DesenvolvedorRepository) usecase.DesenvolvedorUseCase {
	return &desenvolvedorUseCase{
		desenvolvedorRepository: desenvolvedorRepository,
	}
}

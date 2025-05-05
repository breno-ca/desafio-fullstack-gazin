package nivel

import (
	"backend/internal/repository"
	"backend/internal/usecase"
)

type nivelUseCase struct {
	nivelRepository repository.NivelRepository
}

const (
	// Error Messages
	errorSavingData           string = "error saving data"
	errorListingData          string = "error listing data"
	errorDeletingData         string = "error deleting data"
	invalidParameter          string = "invalid parameter"
	invalidRequestBody        string = "invalid request body"
	invalidDataShouldAdequate string = "invalid data should be adequate"
	nivelCannotBeDeleted      string = "nivel cannot be deleted. it is associated with a developer"
	nothingToList             string = "nothing to list"
)

func NewUseCase(nivelRepository repository.NivelRepository) usecase.NivelUseCase {
	return &nivelUseCase{
		nivelRepository: nivelRepository,
	}
}

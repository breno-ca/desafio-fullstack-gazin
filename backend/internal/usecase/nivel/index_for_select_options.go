package nivel

import (
	"errors"
	"net/http"

	"backend/pkg/response"
)

// @Summary      Listar opções de níveis
// @Description  Retorna opções simplificadas de níveis para uso em selects
// @Tags         Níveis
// @Router       /niveis/select-options [get]
// @Produce      json
// @Success      200  {array}   presenter.NivelSelectOptions
// @Failure      404  {object}  response.ErrorResponse
// @Failure      500  {object}  response.ErrorResponse
func (nu *nivelUseCase) IndexForSelectOptions(w http.ResponseWriter, r *http.Request) {
	niveis, total, err := nu.nivelRepository.IndexForSelectOptions()
	if err != nil {
		response.WriteError(w, http.StatusInternalServerError, errorListingData, err)
		return
	}
	if total == 0 {
		response.WriteError(w, http.StatusNotFound, errorListingData, errors.New(nothingToList))
		return
	}

	response.Write(w, http.StatusOK, niveis)
}

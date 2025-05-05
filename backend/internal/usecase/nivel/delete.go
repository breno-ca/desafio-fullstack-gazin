package nivel

import (
	"errors"
	"net/http"

	"backend/pkg/response"

	"github.com/google/uuid"
)

// @Summary      Remover nível
// @Description  Remove um nível pelo ID
// @Tags         Níveis
// @Router       /niveis/{id} [delete]
// @Produce      json
// @Param        id   path      int  true  "ID do nível"
// @Success      204  {string}  string  "Removido com sucesso"
// @Failure      404  {object}  response.ErrorResponse
// @Failure      500  {object}  response.ErrorResponse
func (nu *nivelUseCase) Delete(w http.ResponseWriter, r *http.Request) {
	paramID := r.PathValue("id")
	_, err := uuid.Parse(paramID)
	if err != nil {
		response.WriteError(w, http.StatusUnprocessableEntity, invalidParameter, err)
		return
	}

	hasNivelAssociado, err := nu.nivelRepository.HasNivelAssociado(paramID)
	if err != nil {
		response.WriteError(w, http.StatusUnprocessableEntity, errorDeletingData, err)
		return

	}
	if hasNivelAssociado {
		response.WriteError(w, http.StatusBadRequest, errorDeletingData, errors.New(nivelCannotBeDeleted))
		return
	}

	err = nu.nivelRepository.Delete(paramID)
	if err != nil {
		response.WriteError(w, http.StatusInternalServerError, errorDeletingData, err)
		return
	}

	response.Write(w, http.StatusNoContent, nil)
}

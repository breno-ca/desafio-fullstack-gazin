package desenvolvedor

import (
	"errors"
	"net/http"

	"backend/pkg/response"

	"github.com/google/uuid"
)

// @Summary      Remover desenvolvedor
// @Description  Remove um desenvolvedor pelo ID
// @Tags         Desenvolvedores
// @Router       /desenvolvedores/{id} [delete]
// @Produce      json
// @Param        id   path      int  true  "ID do desenvolvedor"
// @Success      204  {string}  string     "Removido com sucesso"
// @Failure      404  {object}  response.ErrorResponse
// @Failure      500  {object}  response.ErrorResponse
func (du *desenvolvedorUseCase) Delete(w http.ResponseWriter, r *http.Request) {
	paramID := r.PathValue("id")
	_, err := uuid.Parse(paramID)
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, invalidParameter, err)
		return
	}

	hasDesenvolvedor, err := du.desenvolvedorRepository.HasDesenvolvedorByID(paramID)
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, errorDeletingData, err)
		return
	}
	if !hasDesenvolvedor {
		response.WriteError(w, http.StatusBadRequest, errorDeletingData, errors.New(developerNotFound))
		return
	}

	err = du.desenvolvedorRepository.Delete(paramID)
	if err != nil {
		response.WriteError(w, http.StatusInternalServerError, errorDeletingData, err)
		return
	}

	response.Write(w, http.StatusNoContent, nil)
}

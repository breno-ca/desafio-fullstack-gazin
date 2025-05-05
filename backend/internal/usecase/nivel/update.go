package nivel

import (
	"net/http"

	"backend/internal/entity"
	"backend/pkg/request"
	"backend/pkg/response"

	"github.com/google/uuid"
)

// @Summary      Atualizar nível
// @Description  Atualiza um nível existente
// @Tags         Níveis
// @Router       /niveis/{id} [put]
// @Accept       json
// @Produce      json
// @Param        id       path      int               true   "ID do nível"
// @Param        request  body      presenter.Nivel   true   "Dados atualizados do nível"
// @Success      200      {object}  presenter.Nivel
// @Failure      400      {object}  response.ErrorResponse
// @Failure      422      {object}  response.ErrorResponse
func (nu *nivelUseCase) Update(w http.ResponseWriter, r *http.Request) {
	paramID := r.PathValue("id")
	_, err := uuid.Parse(paramID)
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, errorSavingData, err)
		return
	}

	var nivel entity.Nivel
	err = request.RetrieveRequestData(r, &nivel)
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, invalidDataShouldAdequate, err)
		return
	}
	nivel.ID = paramID

	err = nivel.Validate()
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, invalidDataShouldAdequate, err)
		return
	}

	err = nu.nivelRepository.Update(nivel)
	if err != nil {
		response.WriteError(w, http.StatusUnprocessableEntity, err.Error(), err)
		return
	}

	response.Write(w, http.StatusOK, nivel)
}

package nivel

import (
	"net/http"

	"backend/internal/entity"
	"backend/pkg/request"
	"backend/pkg/response"

	"github.com/google/uuid"
)

// @Summary      Criar nível
// @Description  Cria um novo nível
// @Tags         Níveis
// @Router       /niveis [post]
// @Accept       json
// @Produce      json
// @Param        request  body      entity.Nivel  true  "Dados do nível"
// @Success      201      {object}  presenter.Nivel
// @Failure      400      {object}  response.ErrorResponse
// @Failure      500      {object}  response.ErrorResponse
func (nu *nivelUseCase) Create(w http.ResponseWriter, r *http.Request) {
	var nivel entity.Nivel
	err := request.RetrieveRequestData(r, &nivel)
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, invalidDataShouldAdequate, err)
		return
	}

	err = nivel.Validate()
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, invalidDataShouldAdequate, err)
		return
	}

	uuid, err := uuid.NewV7()
	if err != nil {
		response.WriteError(w, http.StatusInternalServerError, errorSavingData, err)
		return
	}

	nivel.ID = uuid.String()

	err = nu.nivelRepository.Create(nivel)
	if err != nil {
		response.WriteError(w, http.StatusInternalServerError, errorSavingData, err)
		return
	}

	response.Write(w, http.StatusCreated, nivel)
}

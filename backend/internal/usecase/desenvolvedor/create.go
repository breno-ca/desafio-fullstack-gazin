package desenvolvedor

import (
	"net/http"

	"backend/internal/entity"
	"backend/pkg/request"
	"backend/pkg/response"

	"github.com/google/uuid"
)

// @Summary      Criar desenvolvedor
// @Description  Cria um novo desenvolvedor
// @Tags         Desenvolvedores
// @Router       /desenvolvedores [post]
// @Accept       json
// @Produce      json
// @Param        request  body      entity.Desenvolvedor     true  "Dados do desenvolvedor"
// @Success      201      {object}  presenter.Desenvolvedor
// @Failure      400      {object}  response.ErrorResponse
// @Failure      500      {object}  response.ErrorResponse
func (du *desenvolvedorUseCase) Create(w http.ResponseWriter, r *http.Request) {
	var desenvolvedor entity.Desenvolvedor
	err := request.RetrieveRequestData(r, &desenvolvedor)
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, invalidRequestBody, err)
		return
	}

	err = desenvolvedor.Validate()
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, invalidDataShouldAdequate, err)
		return
	}

	uuid, err := uuid.NewV7()
	if err != nil {
		response.WriteError(w, http.StatusInternalServerError, errorSavingData, err)
		return
	}

	desenvolvedor.ID = uuid.String()

	err = du.desenvolvedorRepository.Create(desenvolvedor)
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, errorSavingData, err)
		return
	}

	response.Write(w, http.StatusCreated, desenvolvedor)
}

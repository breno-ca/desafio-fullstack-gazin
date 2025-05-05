package desenvolvedor

import (
	"net/http"

	"backend/internal/entity"
	"backend/pkg/request"
	"backend/pkg/response"

	"github.com/google/uuid"
)

// @Summary      Atualizar desenvolvedor
// @Description  Atualiza um desenvolvedor existente
// @Tags         Desenvolvedores
// @Router       /desenvolvedores/{id} [put]
// @Accept       json
// @Produce      json
// @Param        id       path      int                          true  "ID do desenvolvedor"
// @Param        request  body      entity.Desenvolvedor  true  "Dados atualizados do desenvolvedor"
// @Success      200      {object}  presenter.Desenvolvedor
// @Failure      400      {object}  response.ErrorResponse
// @Failure      422      {object}  response.ErrorResponse
func (du *desenvolvedorUseCase) Update(w http.ResponseWriter, r *http.Request) {
	paramID := r.PathValue("id")
	_, err := uuid.Parse(paramID)
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, invalidParameter, err)
		return
	}

	var desenvolvedor entity.Desenvolvedor
	err = request.RetrieveRequestData(r, &desenvolvedor)
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, invalidRequestBody, err)
		return
	}
	desenvolvedor.ID = paramID

	err = desenvolvedor.Validate()
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, invalidDataShouldAdequate, err)
		return
	}

	err = du.desenvolvedorRepository.Update(desenvolvedor)
	if err != nil {
		response.WriteError(w, http.StatusUnprocessableEntity, err.Error(), err)
		return
	}

	response.Write(w, http.StatusOK, desenvolvedor)
}

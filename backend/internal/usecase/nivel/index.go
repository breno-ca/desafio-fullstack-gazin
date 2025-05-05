package nivel

import (
	"errors"
	"net/http"
	"strconv"

	"backend/internal/presenter"
	"backend/pkg/pagination"
	"backend/pkg/response"
)

// @Summary      Listar níveis
// @Description  Retorna todos os níveis cadastrados
// @Tags         Níveis
// @Router      /niveis [get]
// @Produce     json
// @Param       page    query     int     false  "Número da página"
// @Param       order   query     string  false  "Campo para ordenação (ex: nome)"
// @Param       mode    query     string  false  "Direção da ordenação (ASC ou DESC)"
// @Param       search  query     string  false  "Texto de busca"
// @Success     200     {object}  presenter.NivelPaginatedResponse 	"Lista de níveis"
// @Failure 	404     {object}  response.ErrorResponse
// @Failure 	500     {object}  response.ErrorResponse
func (nu *nivelUseCase) Index(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	search := "%" + params.Get("search") + "%" // Default empty
	order := params.Get("order")               // Default empty
	mode := params.Get("mode")                 // Default empty

	page := 1 // Default Page
	if params.Get("page") != "" {
		number, err := strconv.Atoi(params.Get("page"))
		if err == nil && number > 1 {
			page = number
		}
	}
	perPage := 10 // Default Items Per Page
	if params.Get("per_page") != "" {
		number, err := strconv.Atoi(params.Get("per_page"))
		if err == nil && number > 0 {
			perPage = number
		}
	}
	pg := pagination.NewPagination(page, perPage)

	niveis, total, err := nu.nivelRepository.Index(search, order, mode, pg)
	if err != nil {
		response.WriteError(w, http.StatusInternalServerError, errorListingData, err)
		return
	}
	pg.Total = total

	if pg.Total == 0 {
		response.WriteError(w, http.StatusNotFound, errorListingData, errors.New(nothingToList))
		return
	}

	response.Write(w, http.StatusOK,
		presenter.NivelPaginatedResponse{
			Data: niveis,
			Meta: presenter.Meta{
				Total:       pg.Total,
				PerPage:     pg.Limit,
				CurrentPage: pg.Page,
				LastPage:    pg.LastPage(),
			},
		},
	)
}

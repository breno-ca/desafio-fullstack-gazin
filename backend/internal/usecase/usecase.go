package usecase

import "net/http"

type NivelUseCase interface {
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Index(w http.ResponseWriter, r *http.Request)
	IndexForSelectOptions(w http.ResponseWriter, r *http.Request)
}

type DesenvolvedorUseCase interface {
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Index(w http.ResponseWriter, r *http.Request)
}

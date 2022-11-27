package handler

import (
	"encoding/json"
	"fullcycle-api-challenge/server/usecase"
	"net/http"
)

type DollarCotationHandler struct {
	usecase.DollarCotationUseCase
}

func (handler *DollarCotationHandler) Get(w http.ResponseWriter, r *http.Request) {
	cotation, err := handler.DollarCotationUseCase.DollarCotation(r.Context())
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(JSONErr{Message: err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cotation)
}

type JSONErr struct {
	Message string `json:"error"`
}

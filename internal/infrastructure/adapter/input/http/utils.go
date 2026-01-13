package http

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/squillteam/powerflix-backend/internal/domain/entity"
)

type ErrorResponse struct {
	Message string `json:"message"`
	Code    string `json:"code"`
	Status  int    `json:"status"`
}

func respondWithJson(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(payload)
}

func handleError(w http.ResponseWriter, err error) {
	switch {
	case
		errors.Is(err, entity.ErrNameTooShort),
		errors.Is(err, entity.ErrEmptyDescription),
		errors.Is(err, entity.ErrEmptyCover):
		errorResponse := ErrorResponse{
			Message: err.Error(),
			Code:    "VALIDATION_ERROR",
			Status:  http.StatusBadRequest,
		}
		respondWithJson(w, http.StatusBadRequest, errorResponse)
	default:
		errorResponse := ErrorResponse{
			Message: "Internal Server Error",
			Code:    "INTERNAL_SERVER_ERROR",
			Status:  http.StatusInternalServerError,
		}
		respondWithJson(w, http.StatusInternalServerError, errorResponse)
	}
}

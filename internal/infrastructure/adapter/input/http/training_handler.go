package http

import (
	"encoding/json"
	"net/http"

	"github.com/squillteam/powerflix-backend/internal/application/port/input"
)

type TrainingHandler struct {
	getAllUseCase input.GetAllTrainingUseCase
	createUseCase input.CreateTrainingUseCase
}

func NewTrainingHandler(
	getAllUseCase input.GetAllTrainingUseCase,
	createUseCase input.CreateTrainingUseCase,
) *TrainingHandler {
	return &TrainingHandler{
		getAllUseCase: getAllUseCase,
		createUseCase: createUseCase,
	}
}

func (h *TrainingHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	trainings, err := h.getAllUseCase.Execute()

	if err != nil {
		handleError(w, err)
		return
	}

	trainingsResponse := make([]TrainingResponse, len(trainings))

	for i, training := range trainings {
		trainingsResponse[i] = TrainingResponse{
			ID:          training.ID,
			Name:        training.Name,
			Description: training.Description,
			Cover:       training.Cover,
		}
	}

	respondWithJson(w, http.StatusOK, trainingsResponse)
}

func (h *TrainingHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateTrainingRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		errorResponse := ErrorResponse{
			Message: "Invalid JSON format",
			Code:    "INVALID_JSON",
			Status:  http.StatusBadRequest,
		}
		respondWithJson(w, http.StatusBadRequest, errorResponse)
		return
	}

	trainingInput := &input.TrainingInput{
		Name:        req.Name,
		Description: req.Description,
		Cover:       req.Cover,
	}

	training, err := h.createUseCase.Execute(trainingInput)

	if err != nil {
		handleError(w, err)
		return
	}

	trainingResponse := TrainingResponse{
		ID:          training.ID,
		Name:        training.Name,
		Description: training.Description,
		Cover:       training.Cover,
	}

	respondWithJson(w, http.StatusCreated, trainingResponse)
}

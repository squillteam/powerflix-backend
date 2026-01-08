package http

import "github.com/google/uuid"

type CreateTrainingRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Cover       string `json:"cover"`
}

type TrainingResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Cover       string    `json:"cover"`
}

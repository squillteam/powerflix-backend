package input

import (
	"github.com/squillteam/powerflix-backend/internal/domain/entity"
)

type TrainingInput struct {
	Name        string
	Description string
	Cover       string
}

type GetAllTrainingUseCase interface {
	Execute() ([]*entity.Training, error)
}

type CreateTrainingUseCase interface {
	Execute(trainingInput *TrainingInput) (*entity.Training, error)
}

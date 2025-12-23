package usecase

import (
	"github.com/squillteam/powerflix-backend/internal/application/port/input"
	"github.com/squillteam/powerflix-backend/internal/application/port/output"
	"github.com/squillteam/powerflix-backend/internal/domain/entity"
)

type trainingUseCaseImpl struct {
	trainingRepo output.TrainingRepository
}

func NewTrainingUseCase(trainingRepo output.TrainingRepository) input.GetAllTrainingUseCase {
	return &trainingUseCaseImpl{
		trainingRepo: trainingRepo,
	}
}

func (u trainingUseCaseImpl) Execute() ([]entity.Training, error) {
	trainings, err := u.trainingRepo.GetAll()

	if err != nil {
		return nil, err
	}

	return trainings, nil
}

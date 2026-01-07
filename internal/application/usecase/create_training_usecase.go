package usecase

import (
	"fmt"

	"github.com/squillteam/powerflix-backend/internal/application/port/input"
	"github.com/squillteam/powerflix-backend/internal/application/port/output"
	"github.com/squillteam/powerflix-backend/internal/domain/entity"
)

type createTrainingUseCase struct {
	trainingRepository output.TrainingRepository
}

func NewCreateTrainingUseCase(trainingRepository output.TrainingRepository) input.CreateTrainingUseCase {
	return &createTrainingUseCase{
		trainingRepository: trainingRepository,
	}
}

func (c createTrainingUseCase) Execute(trainingInput *input.TrainingInput) (*entity.Training, error) {
	training, err := entity.NewTraining(trainingInput.Name, trainingInput.Description, trainingInput.Cover)

	if err != nil {
		return nil, fmt.Errorf("Failed to create training %w", err)
	}

	savedTraining, err := c.trainingRepository.Save(training)

	if err != nil {
		return nil, fmt.Errorf("Failed to save training: %w", err)
	}

	return savedTraining, nil
}

package usecase

import (
	"github.com/squillteam/powerflix-backend/internal/application/port/input"
	"github.com/squillteam/powerflix-backend/internal/application/port/output"
	"github.com/squillteam/powerflix-backend/internal/domain/entity"
)

// createTrainingUseCase provides the default implementation of the training creation use case.
type createTrainingUseCase struct {
	trainingRepository output.TrainingRepository
}

// NewCreateTrainingUseCase creates a new CreateTrainingUseCase that persists trainings using the given repository.
func NewCreateTrainingUseCase(trainingRepository output.TrainingRepository) input.CreateTrainingUseCase {
	return &createTrainingUseCase{
		trainingRepository: trainingRepository,
	}
}

// Execute creates a new Training entity from the provided input, saves it, and returns the created training.
func (c createTrainingUseCase) Execute(trainingInput *input.TrainingInput) (*entity.Training, error) {
	training := entity.NewTraining(trainingInput.Name, trainingInput.Description, trainingInput.Cover)

	savedTraining, err := c.trainingRepository.Save(training)

	if err != nil {
		return nil, err
	}

	return savedTraining, nil
}

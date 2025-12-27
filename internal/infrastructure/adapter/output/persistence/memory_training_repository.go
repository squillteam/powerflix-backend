package persistence

import (
	"sync"

	"github.com/squillteam/powerflix-backend/internal/application/port/output"
	"github.com/squillteam/powerflix-backend/internal/domain/entity"
)

type MemoryTrainingRepository struct {
	trainings []*entity.Training
	mu        sync.RWMutex
}

func NewMemoryTrainingRepository() output.TrainingRepository {
	return &MemoryTrainingRepository{
		trainings: []*entity.Training{},
	}
}

func (r *MemoryTrainingRepository) Save(training *entity.Training) (*entity.Training, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.trainings = append(r.trainings, training)

	return training, nil
}

func (r *MemoryTrainingRepository) GetAll() ([]*entity.Training, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make([]*entity.Training, 0, len(r.trainings))
	for _, trainingPtr := range r.trainings {
		result = append(result, trainingPtr)
	}

	return result, nil
}

package persistence

import (
	"sync"

	"github.com/squillteam/powerflix-backend/internal/application/port/output"
	"github.com/squillteam/powerflix-backend/internal/domain/entity"
)

type MemoryTrainingRepository struct {
	trainings map[string]*entity.Training
	mu        sync.RWMutex
}

func NewMemoryTrainingRepository() output.TrainingRepository {
	return &MemoryTrainingRepository{
		trainings: map[string]*entity.Training{
			"1": &entity.Training{
				ID:          "1",
				Name:        "Arnold Biceps",
				Description: "Arnold Biceps from hell",
				Cover:       "https://example.com/cover.jpg",
			},
			"2": &entity.Training{
				ID:          "2",
				Name:        "Arnold Triceps",
				Description: "Arnold Triceps from hell",
				Cover:       "https://example.com/cover.jpg",
			},
			"3": &entity.Training{
				ID:          "3",
				Name:        "Arnold Chest",
				Description: "Arnold Chest from hell",
				Cover:       "https://example.com/cover.jpg",
			},
			"4": &entity.Training{
				ID:          "4",
				Name:        "Arnold Back",
				Description: "Arnold Back from hell",
				Cover:       "https://example.com/cover.jpg",
			},
			"5": &entity.Training{
				ID:          "5",
				Name:        "Arnold Legs",
				Description: "Arnold Legs from hell",
				Cover:       "https://example.com/cover.jpg",
			},
		},
	}
}

func (r *MemoryTrainingRepository) GetAllTrainings() ([]entity.Training, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := []entity.Training{}

	for _, trainingPtr := range r.trainings {
		result = append(result, *trainingPtr)
	}

	return result, nil
}

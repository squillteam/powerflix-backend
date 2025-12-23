package output

import (
	"github.com/squillteam/powerflix-backend/internal/domain/entity"
)

type TrainingRepository interface {
	GetAll() ([]entity.Training, error)
	Save(entity.Training) (entity.Training, error)
}

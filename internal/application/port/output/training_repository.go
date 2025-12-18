package output

import (
	"github.com/squillteam/powerflix-backend/internal/domain/entity"
)

type TrainingRepository interface {
	GetAllTraining() ([]entity.Training, error)
}

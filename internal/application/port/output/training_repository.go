package output

import (
	"github.com/squillteam/powerflix-backend/internal/domain/entity"
)

type TrainingRepository interface {
	GetAllTrainings() ([]entity.Training, error)
}

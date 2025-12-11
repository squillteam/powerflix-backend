package input

type TrainingInput struct {
	ID          string
	Name        string
	Description string
	Cover       string
}

type TrainingService interface {
	CreateTraining()
}

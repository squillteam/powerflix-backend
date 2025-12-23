package entity

type Training struct {
	ID          string
	Name        string
	Description string
	Cover       string
}

func NewTraining(name, description, cover string) *Training {
	return &Training{
		Name:        name,
		Description: description,
		Cover:       cover,
	}
}

package entity

import (
	"github.com/google/uuid"
)

type Training struct {
	ID          uuid.UUID
	Name        string
	Description string
	Cover       string
}

func NewTraining(name, description, cover string) *Training {
	return &Training{
		ID:          uuid.New(),
		Name:        name,
		Description: description,
		Cover:       cover,
	}
}

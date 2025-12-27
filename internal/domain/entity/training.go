package entity

import (
	"errors"
	"unicode/utf8"

	"github.com/google/uuid"
)

var (
	ErrNameTooShort     = errors.New("Name must be at least 3 characters long")
	ErrEmptyDescription = errors.New("Description cannot be empty")
	ErrEmptyCover       = errors.New("Cover cannot be empty")
)

type Training struct {
	ID          uuid.UUID
	Name        string
	Description string
	Cover       string
}

func NewTraining(name, description, cover string) (*Training, error) {
	if name == "" || utf8.RuneCountInString(name) < 3 {
		return nil, ErrNameTooShort
	}

	if description == "" {
		return nil, ErrEmptyDescription
	}

	if cover == "" {
		return nil, ErrEmptyCover
	}

	return &Training{
		ID:          uuid.New(),
		Name:        name,
		Description: description,
		Cover:       cover,
	}, nil
}

package main

import (
	"fmt"

	"github.com/squillteam/powerflix-backend/internal/domain/entity"
)

func main() {
	fmt.Println("Hello World!")

	training := entity.Training{
		ID:          "1",
		Name:        "Arnold Biceps",
		Description: "Anorld Biceps from hell",
		Cover:       "https://example.com/cover.jpg",
	}

	fmt.Printf("%+v\n", training)
}

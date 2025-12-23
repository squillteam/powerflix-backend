package main

import (
	"encoding/json"
	"fmt"

	"github.com/squillteam/powerflix-backend/internal/application/usecase"
	"github.com/squillteam/powerflix-backend/internal/infrastructure/adapter/output/persistence"
)

func main() {
	repo := persistence.NewMemoryTrainingRepository()
	trainingUseCase := usecase.NewTrainingUseCase(repo)

	trainings, err := trainingUseCase.Execute()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	jsonData, err := json.MarshalIndent(trainings, "", " ")
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
		return
	}

	fmt.Println(string(jsonData))
}

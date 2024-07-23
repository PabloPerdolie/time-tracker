package main

import (
	"EffectiveMobileTestTask/internal/app"
	_ "EffectiveMobileTestTask/internal/docs"
	"context"
	"log"
)

// @title Time Tracker API
// @version 1.0
// @description This is a sample server for a time tracking application.
// @host localhost:8080
// @BasePath /
func main() {
	a, err := app.NewApp(context.Background())
	if err != nil {
		log.Fatalf("failed to init app: %s", err.Error())
	}

	err = a.Run()
	if err != nil {
		log.Fatalf("failed to run app: %s", err.Error())
	}
}

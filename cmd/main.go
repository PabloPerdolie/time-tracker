package main

import (
	"EffectiveMobileTestTask/internal/app"
	"context"
	"log"
)

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

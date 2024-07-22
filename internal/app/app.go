package app

import (
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

const defaultPort = "8080"

type App struct {
	serviceProvider *serviceProvider
	router          *mux.Router
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}
	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initServiceProvider,
		a.initServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) Run() error {
	return a.runServer()
}

func (a *App) initServer(ctx context.Context) error {

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return a.serviceProvider.initServices()
}

func (a *App) runServer() error {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	log.Printf("connect to http://localhost:%s/", port)
	log.Fatal(http.ListenAndServe(":"+port, a.router))

	return nil
}

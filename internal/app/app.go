package app

import (
	"log"
	"os"
	"wallet-test/config"
	"wallet-test/internal/databases/repos"
	"wallet-test/internal/server"
	"wallet-test/internal/usecase"
)

type App struct {
	config *config.Config
}

func (app *App) Run() error {
	db, err := app.connPostgres()
	if err != nil {
		return err
	}

	repositories := repos.New(db)
	myUseCase := usecase.NewUseCase(repositories)

	newServer := server.NewServer(*app.config, myUseCase)

	err = newServer.Run()
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	os.Exit(0)
	return nil
}

func NewApp(config *config.Config) *App {
	return &App{config: config}
}

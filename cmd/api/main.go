package main

import (
	"log"
	"os"
	"wallet-test/config"
	"wallet-test/internal/app"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	application := app.NewApp(cfg)
	err = application.Run()
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
}

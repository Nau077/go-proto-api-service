package main

import (
	"context"
	"flag"
	"log"

	"github.com/Nau077/golang-pet-first/internal/app"
	_ "github.com/jackc/pgx/stdlib"
)

var pathConfig string

func init() {
	flag.StringVar(&pathConfig, "config", "./config.json", "path to configuration file")
}

func main() {
	flag.Parse()

	ctx := context.Background()
	a, err := app.NewApp(ctx, "")

	if err != nil {
		log.Fatalf("failed to create app: %s", err.Error())
	}

	err = a.Run()
	if err != nil {
		log.Fatalf("failed to run app %s", err.Error())
	}
}

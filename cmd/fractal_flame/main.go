package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/werener/fractal-flame/internal/application/handlers"
	"github.com/werener/fractal-flame/internal/application/usecase"
	"github.com/werener/fractal-flame/internal/infrastructure/cli"
	"github.com/werener/fractal-flame/pkg/random"
)

func main() {
	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGINT,
	)
	defer stop()

	saver := handlers.FractalSaver{}
	generator := handlers.FractalGenerator{}
	randomizer := random.NewGenerator()

	fractalService := usecase.NewFractalService(saver, generator, *randomizer)

	app := cli.NewApp(fractalService)

	err := app.Run(ctx, os.Args)
	if err != nil {
		log.Println(err)
	}
}

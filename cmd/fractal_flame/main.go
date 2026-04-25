package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/werener/fractal-flame/internal/infrastructure/cli"
)

func main() {
	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	err := cli.Run(ctx, os.Args)

	log.Println(err)
}

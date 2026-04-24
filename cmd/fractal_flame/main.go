package main

import (
	"context"
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

	cli.Run(ctx, os.Args)
}

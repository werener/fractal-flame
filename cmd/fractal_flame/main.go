package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	Run(ctx, os.Args)
}

func Run(ctx context.Context, args []string) {
	for _, arg := range args {
		fmt.Println(arg)
	}
}

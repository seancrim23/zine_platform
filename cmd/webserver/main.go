package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"zine_platform/server"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	server, err := server.NewZinePlatformServer(ctx, server.LoadConfig())
	if err != nil {
		fmt.Println("failed to create new mailto generator server: ", err)
	}

	err = server.StartZinePlatformServer(ctx)
	if err != nil {
		fmt.Println("failed to start mailto generator server: ", err)
	}
}

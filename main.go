package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/juntakoman123/go_todo_app/config"
)

func main() {
	err := run(context.Background())
	if err != nil {
		log.Printf("Failed to terminate server: %v", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	cfg, err := config.New()
	if err != nil {
		return err
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Fatalf("failed to listen port %d: %v", cfg.Port, err)
	}
	url := fmt.Sprintf("http://%s", listener.Addr().String())
	log.Printf("start with: %v", url)

	mux := NewMux()
	server := NewServer(listener, mux)
	return server.Run(ctx)

}

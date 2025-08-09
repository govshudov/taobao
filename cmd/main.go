package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	configs "taobao/internal/config"
	"taobao/internal/handlers"
	"time"
)

const (
	ctxWaitTime       = 5 * time.Second
	ReadHeaderTimeout = 5 * time.Second
	ReadWriteTimeout  = 10 * time.Second
	IdleTimeout       = 60 * time.Second
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	cfg := configs.GetConfig()

	r := handlers.Manager(ctx, "https://www.xportcn.com/api")

	port := cfg.Listen.Port
	server := &http.Server{
		Addr:              port,
		Handler:           r,
		ReadHeaderTimeout: ReadHeaderTimeout,
		ReadTimeout:       ReadWriteTimeout,
		WriteTimeout:      ReadWriteTimeout,
		IdleTimeout:       IdleTimeout,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("failed to start server: %v", err)
		}
	}()
	log.Printf("Server started on port %s", port)

	<-ctx.Done()

	log.Println("got interruption signal")

	ctx, cancel := context.WithTimeout(context.Background(), ctxWaitTime)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("Server shutdown failed: %v", err)
	}
}

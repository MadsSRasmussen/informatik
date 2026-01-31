package main

import (
	"context"
	"fmt"
	"informatik/api/internal/ai"
	"informatik/api/internal/config"
	"informatik/api/internal/db"
	"informatik/api/internal/server"
	"informatik/api/internal/store"
	"io"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func run(ctx context.Context, w io.Writer, getenv func(string) string) error {
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt)
	defer stop()

	_, err := config.New(getenv)
	if err != nil {
		return err
	}

	db, err := db.NewSQLClient(getenv)
	if err != nil {
		return err
	}

	store := store.New(db)

	mistral, err := ai.NewMistralClient(getenv)
	if err != nil {
		return err
	}

	srv := server.New(store, mistral)
	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: srv,
	}

	go func() {
		<-ctx.Done()

		shutdownCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			fmt.Fprintf(w, "Shutdown error: %v", err)
		}
	}()

	fmt.Fprintln(w, "Server listening on port :8080")
	return httpServer.ListenAndServe()
}

func main() {
	ctx := context.Background()

	if err := run(ctx, os.Stdout, os.Getenv); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

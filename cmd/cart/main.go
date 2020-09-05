package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/threkk/cart-service/internal/cart"
)

func main() {
	svc := cart.NewService()
	srv := &http.Server{
		Handler:        svc,
		Addr:           ":3000",
		WriteTimeout:   15 * time.Second,
		ReadTimeout:    15 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		log.Fatal(srv.ListenAndServe())
	}()

	sig := make(chan os.Signal, 1)

	signal.Notify(sig, os.Interrupt)
	signal.Notify(sig, os.Kill)

	<-sig

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	go func() {
		srv.Shutdown(ctx)
	}()

	<-ctx.Done()

	os.Exit(1)
}

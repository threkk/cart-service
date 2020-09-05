package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/threkk/cart-service/internal/cart"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	svc := cart.NewService()
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	db.AutoMigrate(&cart.Cart{})
	db.AutoMigrate(&cart.Item{})

	svc.DB = db

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

	os.Exit(0)
}

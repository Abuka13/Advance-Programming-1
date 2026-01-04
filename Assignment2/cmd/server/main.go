package main

import (
	"assignment2/internal/server"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	srv := server.New()

	mux := http.NewServeMux()
	mux.HandleFunc("POST /data", srv.PostData)
	mux.HandleFunc("GET /data", srv.GetData)
	mux.HandleFunc("DELETE /data/{key}", srv.DeleteData)
	mux.HandleFunc("GET /stats", srv.Stats)

	// Запускаем worker
	go srv.StartWorker()

	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// Запускаем сервер
	go func() {
		fmt.Println("Server running on :8080")
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Error: %v\n", err)
		}
	}()

	// Ждем Ctrl+C
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	// Останавливаем
	fmt.Println("Shutting down...")
	srv.Shutdown()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	httpServer.Shutdown(ctx)

	fmt.Println("Server stopped")
}

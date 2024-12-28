package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dev-rathore/go/rest-apis/internal/config"
	"github.com/dev-rathore/go/rest-apis/internal/http/handlers/student"
	"github.com/dev-rathore/go/rest-apis/internal/storage/sqlite"
)

func main() {
  // Load config
  cfg := config.MustLoad()

  // Database setup
  storage, err := sqlite.New(cfg)

  if err != nil {
    log.Fatal(err)
  }

  slog.Info("Storage initialized", slog.String("env", cfg.Env), slog.String("version", "1.0.0"))

  // Setup router
  router := http.NewServeMux()

  router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Welcome to REST APIs in Go"))
  })

  router.HandleFunc("GET /api/students", student.GetList(storage))
  router.HandleFunc("POST /api/students", student.Create(storage))
  router.HandleFunc("GET /api/students/{id}", student.GetById(storage))

  server := http.Server {
    // Addr: cfg.HTTPServer.Addr,
    Addr: cfg.Addr,
    Handler: router,
  }

  slog.Info("Server started at", slog.String("address", cfg.HTTPServer.Addr))
  // fmt.Printf("Server started at %s", cfg.HTTPServer.Addr)

  done := make(chan os.Signal, 1)

  signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

  go func() {
    err := server.ListenAndServe()
  
    if err != nil {
      log.Fatal(("Failed to start server"))
    }
  }()

  <-done

  slog.Info("Shutting down the server")

  ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)

  defer cancel()

  // err := server.Shutdown(ctx)

  // if err != nil {
  //   slog.Error("Failed to shutdown server", slog.String("error", err.Error()))
  // }

  // Short form of above statements
  if err := server.Shutdown(ctx); err != nil {
    slog.Error("Failed to shutdown server", slog.String("error", err.Error()))
  }

  slog.Info("Server shutdown successfully")
}
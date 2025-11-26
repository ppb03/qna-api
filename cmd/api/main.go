// main.go
package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ppb03/qna-api/internal/config"
	"github.com/ppb03/qna-api/internal/handler"
	"github.com/ppb03/qna-api/internal/model"
	"github.com/ppb03/qna-api/internal/service"
	"github.com/ppb03/qna-api/internal/repository"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	if err := config.Load(); err != nil {
		slog.Error("failed to load config: " +  err.Error())
		os.Exit(1)
	}

	db, err := gorm.Open(postgres.Open(config.DBDSN), &gorm.Config{})
	if err != nil {
		slog.Error("failed to connect to database: " + err.Error())
		os.Exit(1)
	}

	if err := db.AutoMigrate(&model.Question{}, &model.Answer{}); err != nil {
		slog.Error("failed to run migrations", "error", err)
		os.Exit(1)
	}

	questionRepository := repository.NewPostgresQuestionRepository(db)
	answerRepository := repository.NewPostgresAnswerRepository(db)

	questionService := service.NewQuestionService(questionRepository)
	answerService := service.NewAnswerService(answerRepository, questionRepository)

	router := handler.NewRouter(questionService, answerService)

	port := config.ServerPort
	server := &http.Server{
		Addr:         ":" + port,
		Handler:      handler.LoggingMiddleware(router),
		ReadTimeout:  8 * time.Second,
		WriteTimeout: 16 * time.Second,
		IdleTimeout:  16 * time.Second,
	}
	
	slog.Info("server starting", "port", port)
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("server failed to start", "error", err)
			os.Exit(1)
		}
	}()
	
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-done
	slog.Info("server shutting down gracefully")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("server shutdown failed", "error", err)
		os.Exit(1)
	}

	slog.Info("server stopped")
}
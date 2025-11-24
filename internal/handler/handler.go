package handler

import (
	"net/http"
	"log/slog"
	"time"

	"github.com/ppb03/question-answer-api/internal/service"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		slog.Info("request started", "method", r.Method, "path", r.URL.Path, "remote_addr", r.RemoteAddr)

		next.ServeHTTP(w, r)

		duration := time.Since(start)
		slog.Info("request completed", "method", r.Method, "path", r.URL.Path, "duration", duration)
	})
}

func NewRouter(questionSvc service.QuestionService, answerSvc service.AnswerService) *http.ServeMux {
	mux := http.NewServeMux()

	// Questions routes
	mux.HandleFunc("GET /questions/", GetAllQuestions(questionSvc))
	mux.HandleFunc("POST /questions/", CreateQuestion(questionSvc))
	mux.HandleFunc("GET /questions/{id}", GetQuestionWithAnswers(questionSvc))
	mux.HandleFunc("DELETE /questions/{id}", DeleteQuestion(questionSvc))

	// Answers routes
	mux.HandleFunc("POST /questions/{id}/answers/", CreateAnswer(answerSvc))
	mux.HandleFunc("GET /answers/{id}", GetAnswer(answerSvc))
	mux.HandleFunc("DELETE /answers/{id}", DeleteAnswer(answerSvc))

	return mux
}
package handler

import (
	"net/http"
	"log/slog"
	"time"

	"github.com/ppb03/qna-api/internal/service"
)

// LoggingMiddleware adds logging to each request with measurement of the time spent on servicing it.
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		slog.Info("request started", "method", r.Method, "path", r.URL.Path, "remote_addr", r.RemoteAddr)

		next.ServeHTTP(w, r)

		duration := time.Since(start)
		slog.Info("request completed", "method", r.Method, "path", r.URL.Path, "duration", duration)
	})
}

// NewRouter creates a new HTTP serve mux with registered handlers.
func NewRouter(questionService service.QuestionService, answerService service.AnswerService) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /questions/", createQuestion(questionService))
	mux.HandleFunc("DELETE /questions/{id}", deleteQuestion(questionService))
	mux.HandleFunc("GET /questions/", getAllQuestions(questionService))
	mux.HandleFunc("GET /questions/{id}", getQuestionWithAnswers(questionService))

	mux.HandleFunc("POST /questions/{id}/answers/", createAnswer(answerService))
	mux.HandleFunc("DELETE /answers/{id}", deleteAnswer(answerService))
	mux.HandleFunc("GET /answers/{id}", getAnswer(answerService))

	return mux
}
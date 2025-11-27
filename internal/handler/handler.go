// Package handler implements HTTP handlers and middleware for question and answer management.
package handler

import (
	"log/slog"
	"net/http"
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
	mux.HandleFunc("GET /questions/{id}", getQuestionByID(questionService))
	mux.HandleFunc("GET /questions/", getAllQuestions(questionService))

	mux.HandleFunc("POST /questions/{id}/answers/", createAnswer(answerService))
	mux.HandleFunc("DELETE /answers/{id}", deleteAnswer(answerService))
	mux.HandleFunc("GET /answers/{id}", getAnswer(answerService))

	return mux
}

// ! Костыль.
func errorStatusCode(err error) int {
	serviceErrMapping := map[error]int{
		service.ErrEmptyText:     http.StatusBadRequest,
		service.ErrEmptyUserID:   http.StatusBadRequest,
		service.ErrInvalidUserID: http.StatusBadRequest,

		service.ErrQuestionNotExists: http.StatusNotFound,
		service.ErrAnswerNotExists:   http.StatusNotFound,

		service.ErrRepositoryFailure: http.StatusInternalServerError,
	}

	statusCode, ok := serviceErrMapping[err]
	if !ok {
		statusCode = http.StatusInternalServerError
	}

	return statusCode
}

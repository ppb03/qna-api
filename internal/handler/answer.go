package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ppb03/qna-api/internal/service"
)

func createAnswer(svc service.AnswerService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		questionID, err := strconv.ParseUint(r.PathValue("id"), 10, 32)
		if err != nil {
			http.Error(w, "invalid question id", http.StatusBadRequest)
			return
		}

		rbody := struct {
			UserID string `json:"user_id"`
			Text string `json:"text"`
		}{}

		if err := json.NewDecoder(r.Body).Decode(&rbody); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		answer, err := svc.Create(r.Context(), uint(questionID), rbody.UserID, rbody.Text)
		if err != nil {
			http.Error(w, err.Error(), errorStatusCode(err))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(answer)
	}
}

func getAnswer(svc service.AnswerService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseUint(r.PathValue("id"), 10, 32)
		if err != nil {
			http.Error(w, "invalid answer id", http.StatusBadRequest)
			return
		}

		answer, err := svc.GetByID(r.Context(), uint(id))
		if err != nil {
			http.Error(w, err.Error(), errorStatusCode(err))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(answer)
	}
}

func deleteAnswer(svc service.AnswerService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseUint(r.PathValue("id"), 10, 32)
		if err != nil {
			http.Error(w, "invalid answer id", http.StatusBadRequest)
			return
		}

		if err := svc.Delete(r.Context(), uint(id)); err != nil {
			http.Error(w, err.Error(), errorStatusCode(err))
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

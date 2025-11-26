package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ppb03/qna-api/internal/service"
)

func createQuestion(svc service.QuestionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rbody := struct {
			Text string `json:"text"`
		}{}
		
		if err := json.NewDecoder(r.Body).Decode(&rbody); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		
		question, err := svc.Create(r.Context(), rbody.Text)
		if err != nil {
			http.Error(w, err.Error(), errorStatusCode(err))
			return
		}
		
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(question)
	}
}

func deleteQuestion(svc service.QuestionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseUint(r.PathValue("id"), 10, 32)
		if err != nil {
			http.Error(w, "invalid question ID", http.StatusBadRequest)
			return
		}

		if err := svc.Delete(r.Context(), uint(id)); err != nil {
			http.Error(w, err.Error(), errorStatusCode(err))
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func getQuestionByID(svc service.QuestionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseUint(r.PathValue("id"), 10, 32)
		if err != nil {
			http.Error(w, "invalid question ID", http.StatusBadRequest)
			return
		}

		
		question, err := svc.GetByID(r.Context(), uint(id))
		if err != nil {
			http.Error(w, err.Error(), errorStatusCode(err))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(question)
	}
}

func getAllQuestions(svc service.QuestionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		questions, err := svc.GetAll(r.Context())
		if err != nil {
			http.Error(w, err.Error(), errorStatusCode(err))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(questions)
	}
}

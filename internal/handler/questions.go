package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ppb03/question-answer-api/internal/models"
	"github.com/ppb03/question-answer-api/internal/service"
)

func GetAllQuestions(questionSvc service.QuestionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		questions, err := questionSvc.GetAll(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(questions)
	}
}

func CreateQuestion(questionSvc service.QuestionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var question models.Question
		if err := json.NewDecoder(r.Body).Decode(&question); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := questionSvc.Create(r.Context(), &question); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(question)
	}
}

func GetQuestionWithAnswers(questionSvc service.QuestionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseUint(r.PathValue("id"), 10, 32)
		if err != nil {
			http.Error(w, "invalid question id", http.StatusBadRequest)
			return
		}

		question, err := questionSvc.GetWithAnswers(r.Context(), uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if question == nil {
			http.Error(w, "question not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(question)
	}
}

func DeleteQuestion(questionSvc service.QuestionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseUint(r.PathValue("id"), 10, 32)
		if err != nil {
			http.Error(w, "invalid question id", http.StatusBadRequest)
			return
		}

		if err := questionSvc.Delete(r.Context(), uint(id)); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
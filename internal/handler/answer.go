package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ppb03/question-answer-api/internal/models"
	"github.com/ppb03/question-answer-api/internal/service"
)

func CreateAnswer(answerSvc service.AnswerService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var answer models.Answer
		if err := json.NewDecoder(r.Body).Decode(&answer); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		questionID, err := strconv.ParseUint(r.PathValue("id"), 10, 32)
		if err != nil {
			http.Error(w, "invalid question id", http.StatusBadRequest)
			return
		}
		answer.QuestionID = uint(questionID)

		if err := answerSvc.Create(r.Context(), &answer); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(answer)
	}
}

func GetAnswer(answerSvc service.AnswerService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseUint(r.PathValue("id"), 10, 32)
		if err != nil {
			http.Error(w, "invalid answer id", http.StatusBadRequest)
			return
		}

		answer, err := answerSvc.GetByID(r.Context(), uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if answer == nil {
			http.Error(w, "answer not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(answer)
	}
}

func DeleteAnswer(answerSvc service.AnswerService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseUint(r.PathValue("id"), 10, 32)
		if err != nil {
			http.Error(w, "invalid answer id", http.StatusBadRequest)
			return
		}

		if err := answerSvc.Delete(r.Context(), uint(id)); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
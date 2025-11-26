package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ppb03/qna-api/internal/model"
	"github.com/ppb03/qna-api/internal/service"
)

func getAllQuestions(questionService service.QuestionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		questions, err := questionService.GetAll(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(questions)
	}
}

func createQuestion(questionService service.QuestionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		questionText := struct {
			Text string `json:"text"`
		}{}
		
		if err := json.NewDecoder(r.Body).Decode(&questionText); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		
		if questionText.Text == "" {
			http.Error(w, "text cannot be empty", http.StatusBadRequest)
			return
		}
		
		question := model.Question{Text: questionText.Text}
		if err := questionService.Create(r.Context(), &question); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(question)
	}
}

func getQuestionWithAnswers(questionService service.QuestionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseUint(r.PathValue("id"), 10, 32)
		if err != nil {
			http.Error(w, "invalid question id", http.StatusBadRequest)
			return
		}

		question, err := questionService.GetWithAnswers(r.Context(), uint(id))
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

func deleteQuestion(questionService service.QuestionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseUint(r.PathValue("id"), 10, 32)
		if err != nil {
			http.Error(w, "invalid question id", http.StatusBadRequest)
			return
		}

		if err := questionService.Delete(r.Context(), uint(id)); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ppb03/qna-api/internal/mocks"
	"github.com/ppb03/qna-api/internal/model"
	"github.com/ppb03/qna-api/internal/repository"
	"github.com/ppb03/qna-api/internal/service"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateQuestion_Success(t *testing.T) {
	mockQuestionRepo := new(mocks.MockQuestionRepository)
	questionService := service.NewQuestionService(mockQuestionRepo)
	handler := NewRouter(questionService, nil)

	expectedQuestion := &model.Question{ID: 1, Text: "Test question"}
	
	mockQuestionRepo.On("Create", mock.Anything, mock.MatchedBy(func(q *model.Question) bool {
		return q.Text == expectedQuestion.Text
	})).Return(expectedQuestion, nil)

	requestBody := map[string]string{"text": "Test question"}
	body, _ := json.Marshal(requestBody)

	req := httptest.NewRequest("POST", "/questions/", bytes.NewReader(body))
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)

	var response model.Question
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, expectedQuestion.ID, response.ID)
	assert.Equal(t, expectedQuestion.Text, response.Text)

	mockQuestionRepo.AssertExpectations(t)
}

func TestCreateQuestion_ErrEmptyText(t *testing.T) {
	mockQuestionRepo := new(mocks.MockQuestionRepository)
	questionService := service.NewQuestionService(mockQuestionRepo)
	handler := NewRouter(questionService, nil)

	requestBody := map[string]string{"text": ""}
	body, _ := json.Marshal(requestBody)

	req := httptest.NewRequest("POST", "/questions/", bytes.NewReader(body))
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Contains(t, rr.Body.String(), service.ErrEmptyText.Error())
}

func TestGetQuestion_Success(t *testing.T) {
	mockQuestionRepo := new(mocks.MockQuestionRepository)
	questionService := service.NewQuestionService(mockQuestionRepo)
	handler := NewRouter(questionService, nil)

	expectedQuestion := &model.Question{
		ID:      1,
		Text:    "Test question",
		Answers: []model.Answer{},
	}
	
	mockQuestionRepo.On("GetByID", mock.Anything, uint(1)).
		Return(expectedQuestion, nil)

	req := httptest.NewRequest("GET", "/questions/1", nil)
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response model.Question
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, expectedQuestion.ID, response.ID)
	assert.Equal(t, expectedQuestion.Text, response.Text)

	mockQuestionRepo.AssertExpectations(t)
}

func TestGetQuestion_ErrQuestionNotExists(t *testing.T) {
	mockQuestionRepo := new(mocks.MockQuestionRepository)
	questionService := service.NewQuestionService(mockQuestionRepo)
	handler := NewRouter(questionService, nil)

	mockQuestionRepo.On("GetByID", mock.Anything, uint(999)).
		Return((*model.Question)(nil), repository.ErrQuestionNotFound)

	req := httptest.NewRequest("GET", "/questions/999", nil)
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
	assert.Contains(t, rr.Body.String(), service.ErrQuestionNotExists.Error())

	mockQuestionRepo.AssertExpectations(t)
}

func TestGetQuestion_ErrInvalidID(t *testing.T) {
	mockQuestionRepo := new(mocks.MockQuestionRepository)
	questionService := service.NewQuestionService(mockQuestionRepo)
	handler := NewRouter(questionService, nil)

	req := httptest.NewRequest("GET", "/questions/-123", nil)
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Contains(t, rr.Body.String(), "invalid question ID")
}

func TestGetAllQuestions_Success(t *testing.T) {
	mockQuestionRepo := new(mocks.MockQuestionRepository)
	questionService := service.NewQuestionService(mockQuestionRepo)
	handler := NewRouter(questionService, nil)

	expectedQuestions := []model.Question{
		{ID: 1, Text: "Question 1"},
		{ID: 2, Text: "Question 2"},
	}
	mockQuestionRepo.On("GetAll", mock.Anything).Return(expectedQuestions, nil)

	req := httptest.NewRequest("GET", "/questions/", nil)
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response []model.Question
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Len(t, response, 2)
	assert.Equal(t, expectedQuestions[0].ID, response[0].ID)
	assert.Equal(t, expectedQuestions[0].Text, response[0].Text)

	mockQuestionRepo.AssertExpectations(t)
}

func TestDeleteQuestion_Success(t *testing.T) {
	mockQuestionRepo := new(mocks.MockQuestionRepository)
	questionService := service.NewQuestionService(mockQuestionRepo)
	handler := NewRouter(questionService, nil)

	mockQuestionRepo.On("Delete", mock.Anything, uint(1)).Return(nil)

	req := httptest.NewRequest("DELETE", "/questions/1", nil)
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNoContent, rr.Code)
	mockQuestionRepo.AssertExpectations(t)
}

func TestDeleteQuestion_ErrQuestionNotExists(t *testing.T) {
	mockQuestionRepo := new(mocks.MockQuestionRepository)
	questionService := service.NewQuestionService(mockQuestionRepo)
	handler := NewRouter(questionService, nil)

	mockQuestionRepo.On("Delete", mock.Anything, uint(999)).
		Return(repository.ErrQuestionNotFound)

	req := httptest.NewRequest("DELETE", "/questions/999", nil)
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
	assert.Contains(t, rr.Body.String(), service.ErrQuestionNotExists.Error())

	mockQuestionRepo.AssertExpectations(t)
}

func TestCreateAnswer_Success(t *testing.T) {
	mockQuestionRepo := new(mocks.MockQuestionRepository)
	mockAnswerRepo := new(mocks.MockAnswerRepository)

	questionService := service.NewQuestionService(mockQuestionRepo)
	answerService := service.NewAnswerService(mockAnswerRepo, mockQuestionRepo)
	handler := NewRouter(questionService, answerService)

	question := &model.Question{ID: 1, Text: "Test question"}
	expectedAnswer := &model.Answer{
		ID:         1,
		QuestionID: 1,
		UserID:     "123e4567-e89b-12d3-a456-426614174000",
		Text:       "Test answer",
	}

	mockQuestionRepo.On("GetByID", mock.Anything, uint(1)).Return(question, nil)
	mockAnswerRepo.On("Create", mock.Anything, mock.MatchedBy(func(a *model.Answer) bool {
		return a.QuestionID == 1 && 
		       a.UserID == "123e4567-e89b-12d3-a456-426614174000" && 
		       a.Text == "Test answer"
	})).Return(expectedAnswer, nil)

	requestBody := map[string]string{
		"user_id": "123e4567-e89b-12d3-a456-426614174000",
		"text":    "Test answer",
	}
	body, _ := json.Marshal(requestBody)

	req := httptest.NewRequest("POST", "/questions/1/answers/", bytes.NewReader(body))
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)

	var response model.Answer
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, expectedAnswer.ID, response.ID)
	assert.Equal(t, expectedAnswer.Text, response.Text)

	mockQuestionRepo.AssertExpectations(t)
	mockAnswerRepo.AssertExpectations(t)
}

func TestCreateAnswer_ErrEmptyText(t *testing.T) {
	mockQuestionRepo := new(mocks.MockQuestionRepository)
	mockAnswerRepo := new(mocks.MockAnswerRepository)

	questionService := service.NewQuestionService(mockQuestionRepo)
	answerService := service.NewAnswerService(mockAnswerRepo, mockQuestionRepo)
	handler := NewRouter(questionService, answerService)

	requestBody := map[string]string{
		"user_id": "123e4567-e89b-12d3-a456-426614174000",
		"text":    "",
	}
	body, _ := json.Marshal(requestBody)

	req := httptest.NewRequest("POST", "/questions/1/answers/", bytes.NewReader(body))
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Contains(t, rr.Body.String(), service.ErrEmptyText.Error())
}

func TestCreateAnswer_ErrEmptyUserID(t *testing.T) {
	mockQuestionRepo := new(mocks.MockQuestionRepository)
	mockAnswerRepo := new(mocks.MockAnswerRepository)

	questionService := service.NewQuestionService(mockQuestionRepo)
	answerService := service.NewAnswerService(mockAnswerRepo, mockQuestionRepo)
	handler := NewRouter(questionService, answerService)

	requestBody := map[string]string{
		"user_id": "",
		"text":    "Test answer",
	}
	body, _ := json.Marshal(requestBody)

	req := httptest.NewRequest("POST", "/questions/1/answers/", bytes.NewReader(body))
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Contains(t, rr.Body.String(), service.ErrEmptyUserID.Error())
}

func TestCreateAnswer_ErrInvalidUserID(t *testing.T) {
	mockQuestionRepo := new(mocks.MockQuestionRepository)
	mockAnswerRepo := new(mocks.MockAnswerRepository)

	questionService := service.NewQuestionService(mockQuestionRepo)
	answerService := service.NewAnswerService(mockAnswerRepo, mockQuestionRepo)
	handler := NewRouter(questionService, answerService)

	requestBody := map[string]string{
		"user_id": "invalid-uuid",
		"text":    "Test answer",
	}
	body, _ := json.Marshal(requestBody)

	req := httptest.NewRequest("POST", "/questions/1/answers/", bytes.NewReader(body))
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Contains(t, rr.Body.String(), service.ErrInvalidUserID.Error())
}

func TestCreateAnswer_ErrQuestionNotExists(t *testing.T) {
	mockQuestionRepo := new(mocks.MockQuestionRepository)
	mockAnswerRepo := new(mocks.MockAnswerRepository)

	questionService := service.NewQuestionService(mockQuestionRepo)
	answerService := service.NewAnswerService(mockAnswerRepo, mockQuestionRepo)
	handler := NewRouter(questionService, answerService)

	mockQuestionRepo.On("GetByID", mock.Anything, uint(999)).
		Return((*model.Question)(nil), repository.ErrQuestionNotFound)

	requestBody := map[string]string{
		"user_id": "123e4567-e89b-12d3-a456-426614174000",
		"text":    "Test answer",
	}
	body, _ := json.Marshal(requestBody)

	req := httptest.NewRequest("POST", "/questions/999/answers/", bytes.NewReader(body))
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
	assert.Contains(t, rr.Body.String(), service.ErrQuestionNotExists.Error())

	mockQuestionRepo.AssertExpectations(t)
}

func TestGetAnswer_Success(t *testing.T) {
	mockQuestionRepo := new(mocks.MockQuestionRepository)
	mockAnswerRepo := new(mocks.MockAnswerRepository)

	questionService := service.NewQuestionService(mockQuestionRepo)
	answerService := service.NewAnswerService(mockAnswerRepo, mockQuestionRepo)
	handler := NewRouter(questionService, answerService)

	expectedAnswer := &model.Answer{
		ID:         1,
		QuestionID: 1,
		UserID:     "123e4567-e89b-12d3-a456-426614174000",
		Text:       "Test answer",
	}
	
	mockAnswerRepo.On("GetByID", mock.Anything, uint(1)).Return(expectedAnswer, nil)

	req := httptest.NewRequest("GET", "/answers/1", nil)
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response model.Answer
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, expectedAnswer.ID, response.ID)
	assert.Equal(t, expectedAnswer.Text, response.Text)

	mockAnswerRepo.AssertExpectations(t)
}

func TestGetAnswer_ErrAnswerNotExists(t *testing.T) {
	mockQuestionRepo := new(mocks.MockQuestionRepository)
	mockAnswerRepo := new(mocks.MockAnswerRepository)

	questionService := service.NewQuestionService(mockQuestionRepo)
	answerService := service.NewAnswerService(mockAnswerRepo, mockQuestionRepo)
	handler := NewRouter(questionService, answerService)

	mockAnswerRepo.On("GetByID", mock.Anything, uint(999)).
		Return((*model.Answer)(nil), repository.ErrAnswerNotFound)

	req := httptest.NewRequest("GET", "/answers/999", nil)
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
	assert.Contains(t, rr.Body.String(), service.ErrAnswerNotExists.Error())

	mockAnswerRepo.AssertExpectations(t)
}

func TestDeleteAnswer_Success(t *testing.T) {
	mockQuestionRepo := new(mocks.MockQuestionRepository)
	mockAnswerRepo := new(mocks.MockAnswerRepository)

	questionService := service.NewQuestionService(mockQuestionRepo)
	answerService := service.NewAnswerService(mockAnswerRepo, mockQuestionRepo)
	handler := NewRouter(questionService, answerService)

	mockAnswerRepo.On("Delete", mock.Anything, uint(1)).Return(nil)

	req := httptest.NewRequest("DELETE", "/answers/1", nil)
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNoContent, rr.Code)
	mockAnswerRepo.AssertExpectations(t)
}

func TestDeleteAnswer_ErrAnswerNotExists(t *testing.T) {
	mockQuestionRepo := new(mocks.MockQuestionRepository)
	mockAnswerRepo := new(mocks.MockAnswerRepository)
	questionService := service.NewQuestionService(mockQuestionRepo)
	answerService := service.NewAnswerService(mockAnswerRepo, mockQuestionRepo)
	handler := NewRouter(questionService, answerService)

	mockAnswerRepo.On("Delete", mock.Anything, uint(999)).
		Return(repository.ErrAnswerNotFound)

	req := httptest.NewRequest("DELETE", "/answers/999", nil)
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
	assert.Contains(t, rr.Body.String(), service.ErrAnswerNotExists.Error())

	mockAnswerRepo.AssertExpectations(t)
}
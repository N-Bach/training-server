package controller

import (
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/assert"
	"testing"
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"net/http"
	"bytes"
	"errors"
	"entity"
	"requestModel"
)

type mockFeedbackRepo struct {
	mock.Mock
}

func (m *mockFeedbackRepo) Save(feedback *entity.Feedback) error {
	args := m.Called(feedback)
	return args.Error(0)
}

func TestAddFeedback_Test1(t *testing.T) {
	myCtrl := &Controller{}
	mockRepo := &mockFeedbackRepo{}
	mockRepo.On("Save", mock.Anything).Return(nil)
	myCtrl.FeedbackRepo = mockRepo

	// Prepare request
	rBody := requestModel.RequestFeedback{
		AuthorId: "dmqdobnwghe123492j3nnk23",
		Title: "Test title",
		Description: "Test desciption",
	}
	result, _ := json.Marshal(rBody)
	req := httptest.NewRequest("POST", "http://localhost:8080/feedback", bytes.NewBuffer(result))
	req.Header.Set("Content-Type","application/json")
	w := httptest.NewRecorder()

	// Excute controller
	myCtrl.AddFeedback(w, req)

	// Test controller
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, []byte(`"Feedback added"`), body)
}

// Testcase: type mismatch in request
func TestAddFeedback_Test2(t *testing.T) {
	myCtrl := &Controller{}
	mockRepo := &mockFeedbackRepo{}
	mockRepo.On("Save", mock.Anything).Return(errors.New("Cannot parse"))
	myCtrl.FeedbackRepo = mockRepo

	// Prepare request
	type mockObject struct {
		Id string `gorethink:"id,omitempty" json:"id"`
		AuthorId string `gorethink:"authorId" json:"authorId"`
		Title string `gorethink:"title" json:"title"`
		Description int `gorethink:"description" json:"description"`
	}

	rBody := mockObject{
		AuthorId: "",
		Title: "Test title",
		Description: 3,
	}
	result, _ := json.Marshal(rBody)
	req := httptest.NewRequest("POST", "http://localhost:8080/feedback", bytes.NewBuffer(result))
	req.Header.Set("Content-Type","application/json")
	w := httptest.NewRecorder()

	// Excute controller
	myCtrl.AddFeedback(w, req)

	// Test controller
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	assert.Equal(t, []byte(`"Cannot parse feedback from body"`), body)
}

func TestAddFeedback_Test3(t *testing.T) {
	myCtrl := &Controller{}
	mockRepo := &mockFeedbackRepo{}
	mockRepo.On("Save", mock.Anything).Return(errors.New("Cannot save"))
	myCtrl.FeedbackRepo = mockRepo

	// Prepare request
	rBody := requestModel.RequestFeedback{
		AuthorId: "dmqdobnwghe123492j3nnk23",
		Title: "",
		Description: "Test desciption",
	}
	result, _ := json.Marshal(rBody)
	req := httptest.NewRequest("POST", "http://localhost:8080/feedback", bytes.NewBuffer(result))
	req.Header.Set("Content-Type","application/json")
	w := httptest.NewRecorder()

	// Excute controller
	myCtrl.AddFeedback(w, req)

	// Test controller
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	assert.Equal(t, []byte(`"Cannot create new feedback"`), body)
}

// Testcase: cannot save to db
func TestAddFeedback_Test4(t *testing.T) {
	myCtrl := &Controller{}
	mockRepo := &mockFeedbackRepo{}
	mockRepo.On("Save", mock.Anything).Return(errors.New("Cannot save"))
	myCtrl.FeedbackRepo = mockRepo

	// Prepare request
	rBody := requestModel.RequestFeedback{
		AuthorId: "dmqdobnwghe123492j3nnk23",
		Title: "test title",
		Description: "Test desciption",
	}
	result, _ := json.Marshal(rBody)
	req := httptest.NewRequest("POST", "http://localhost:8080/feedback", bytes.NewBuffer(result))
	req.Header.Set("Content-Type","application/json")
	w := httptest.NewRecorder()

	// Excute controller
	myCtrl.AddFeedback(w, req)

	// Test controller
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	assert.Equal(t, []byte(`"Cannot save feedback"`), body)
}

package controller

import (
	"testing"
	"io/ioutil"
	"net/http/httptest"
	"net/http"
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"entity"
	"time"
	"errors"
)

var (
	TOKEN = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjA0NGUxYTQwLWU2YTktNDhjNC1iNDlkLTg3OGE0NmZlODRjZiJ9.5C57xtuVpLRqh17nDnaa-8ESQs7Elewsw_OAZ8Ry-0E"
	INVALID_TOKEN = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjA0NGUxYTQwLWU2YTktNDhjNC1iNDlkLTg3OGE0NmZlODRjZiJ9.5C57xtuVpLRqh17nDnaa-8ESQs7Elewsw_OAZ8Ry-0A"
	URL = "http://localhost:8080/lessons"
)

type mockLessonRepo struct{
	mock.Mock
}

func (m *mockLessonRepo) Save(lesson *entity.Lesson) error {
	args := m.Called(lesson)
	return args.Error(0)
}
	
func (m *mockLessonRepo) GetOne(id string) (* entity.Lesson, error) {
	return nil, nil
}

func (m *mockLessonRepo) AddEnroll(lesson *entity.Lesson, userId string) error {
	args := m.Called(lesson, userId)
	return args.Error(0)
}

func TestAddLesson_Test1(t *testing.T) {
	myCtrl := &Controller{}
	mockRepo := &mockLessonRepo{}
	mockRepo.On("Save", mock.Anything).Return(nil)
	myCtrl.LessonRepo = mockRepo

	// Prepare request
	rBody := entity.RequestLesson{
		Date: time.Now(),
		Location: "Test location",
		Period: entity.Period{
			From: time.Now(), 
			To: time.Now().Add(30),
		},
		Description: "Test description",
	}
	result, _ := json.Marshal(rBody)
	req := httptest.NewRequest("POST", URL, bytes.NewBuffer(result))
	req.Header.Set("Content-Type","application/json")
	req.Header.Set("Authorization","Bearer " + TOKEN)
	w := httptest.NewRecorder()

	// Excute controller
	myCtrl.AddLesson(w, req)

	// Test controller
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, []byte(`"Lesson created"`), body)
}

func TestAddLesson_Test2(t *testing.T) {
	myCtrl := &Controller{}
	mockRepo := &mockLessonRepo{}
	mockRepo.On("Save", mock.Anything).Return(errors.New("Empty Header"))
	myCtrl.LessonRepo = mockRepo

	// Prepare request
	rBody := entity.RequestLesson{
		Date: time.Now(),
		Location: "Test location",
		Period: entity.Period{
			From: time.Now(), 
			To: time.Now().Add(30),
		},
		Description: "Test description",
	}
	result, _ := json.Marshal(rBody)
	req := httptest.NewRequest("POST", URL, bytes.NewBuffer(result))
	req.Header.Set("Content-Type","application/json")
	//req.Header.Set("Authorization","Bearer " + TOKEN)
	w := httptest.NewRecorder()

	// Excute controller
	myCtrl.AddLesson(w, req)

	// Test controller
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	assert.Equal(t, []byte(`"Request do not have token"`), body)
}

func TestAddLesson_Test3(t *testing.T) {
	myCtrl := &Controller{}
	mockRepo := &mockLessonRepo{}
	mockRepo.On("Save", mock.Anything).Return(errors.New("Empty location"))
	myCtrl.LessonRepo = mockRepo

	// Prepare request
	rBody := entity.RequestLesson{
		Date: time.Now(),
		Location: "",
		Period: entity.Period{
			From: time.Now(), 
			To: time.Now().Add(30),
		},
		Description: "Test description",
	}
	result, _ := json.Marshal(rBody)
	req := httptest.NewRequest("POST", URL, bytes.NewBuffer(result))
	req.Header.Set("Content-Type","application/json")
	req.Header.Set("Authorization","Bearer " + TOKEN)
	w := httptest.NewRecorder()

	// Excute controller
	myCtrl.AddLesson(w, req)

	// Test controller
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	assert.Equal(t, []byte(`"Cannot create new lesson"`), body)
}

func TestAddLesson_Test4(t *testing.T) {
	myCtrl := &Controller{}
	mockRepo := &mockLessonRepo{}
	mockRepo.On("Save", mock.Anything).Return(errors.New("Empty location"))
	myCtrl.LessonRepo = mockRepo

	// Prepare request
	rBody := entity.RequestLesson{
		Date: time.Now(),
		Location: "Test location",
		Period: entity.Period{
			From: time.Now(), 
			To: time.Now().Add(30),
		},
		Description: "Test description",
	}
	result, _ := json.Marshal(rBody)
	req := httptest.NewRequest("POST", URL, bytes.NewBuffer(result))
	req.Header.Set("Content-Type","application/json")
	req.Header.Set("Authorization","Bearer " + TOKEN)
	w := httptest.NewRecorder()

	// Excute controller
	myCtrl.AddLesson(w, req)

	// Test controller
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	assert.Equal(t, []byte(`"Cannot save new lesson"`), body)
}



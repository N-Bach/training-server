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
	"errors"
)

type mockUserRepo struct {
	mock.Mock
}

func (*mockUserRepo) GetByEmail(email string) (*entity.User, error) {
	return nil, nil
}
func (m *mockUserRepo) Save(user *entity.RequestUser) error {
	// return errors.New("asdasdasdasd")

	
	args := m.Called(user)
	return args.Error(0)
}

func TestRegisterUser_Test1(t *testing.T) {
	myCtrl := &Controller{}
	mockRepo := &mockUserRepo{}
	mockRepo.On("Save", mock.Anything).Return(nil)
	myCtrl.UserRepo = mockRepo

	// Prepare request
	rBody := map[string]string{
		"email": "testmail@gmail.com",
		"password": "abc",
	}
	result, _ := json.Marshal(rBody)
	req := httptest.NewRequest("POST", "http://localhost:8080/register", bytes.NewBuffer(result))
	req.Header.Set("Content-Type","application/json")
	w := httptest.NewRecorder()

	// Excute controller
	myCtrl.RegisterUser(w, req)

	// Test controller
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, []byte(`"Saved success"`), body)
}

func TestRegisterUser_Test2(t *testing.T) {
	// test for wrong request format

	myCtrl := &Controller{}
	mockRepo := &mockUserRepo{}
	mockRepo.On("Save", mock.Anything).Return(errors.New("Ko save dc"))
	myCtrl.UserRepo = mockRepo

	// Prepare request
	rBody := map[string]string{
		"email": "testmail@gmail.com",
	}
	result, _ := json.Marshal(rBody)
	req := httptest.NewRequest("POST", "http://localhost:8080/register", bytes.NewBuffer(result))
	req.Header.Set("Content-Type","application/json")
	w := httptest.NewRecorder()

	// Excute controller
	myCtrl.RegisterUser(w, req)

	// Test controller
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	assert.Equal(t, []byte(`"Empty field(s)"`), body)
}

func TestRegisterUser_Test3(t *testing.T) {
	// test for wrong request format

	myCtrl := &Controller{}
	mockRepo := &mockUserRepo{}
	mockRepo.On("Save", mock.Anything).Return(errors.New("Ko save dc"))
	myCtrl.UserRepo = mockRepo

	// Prepare request
	type mockTest struct {
		Email string `json:"email"`
		Password int `json:"password"`
	}

	rBody := mockTest{
		Email: "testmail@gmail.com",
		Password: 1,
	}
	result, _ := json.Marshal(rBody)
	req := httptest.NewRequest("POST", "http://localhost:8080/register", bytes.NewBuffer(result))
	req.Header.Set("Content-Type","application/json")
	w := httptest.NewRecorder()

	// Excute controller
	myCtrl.RegisterUser(w, req)

	// Test controller
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	assert.Equal(t, []byte(`"Incorrect user request format"`), body)
}

func TestRegisterUser_Test4(t *testing.T) {
	// test for wrong request format

	myCtrl := &Controller{}
	mockRepo := &mockUserRepo{}
	mockRepo.On("Save", mock.Anything).Return(errors.New("Ko save dc"))
	myCtrl.UserRepo = mockRepo

	// Prepare request
	rBody := entity.RequestUser{
		Email: "testmail@gmail.com",
		Password: "qwerty",
	}
	result, _ := json.Marshal(rBody)
	req := httptest.NewRequest("POST", "http://localhost:8080/register", bytes.NewBuffer(result))
	req.Header.Set("Content-Type","application/json")
	w := httptest.NewRecorder()

	// Excute controller
	myCtrl.RegisterUser(w, req)

	// Test controller
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	assert.Equal(t, []byte(`"Cannot save user"`), body)
}


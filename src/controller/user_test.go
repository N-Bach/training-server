package controller

import (
	"testing"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"bytes"
	"encoding/json"
)

func TestRegisterUser(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "<html><body>Hello World!</body></html>")
	}

	rBody := map[string]string{
		"email": "testmail@gmail.com",
		"password": "abc",
	}
	result, _ := json.Marshal(rBody)

	req := httptest.NewRequest("POST", "http://localhost:8080/register", bytes.NewBuffer(result))
	req.Header.Set("Content-Type","application/json")
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(string(body))
}
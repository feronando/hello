package internal_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"workshop.go/hello/internal"
)

func Test_handler_ServerHTTP(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "/test", nil)
	response := httptest.NewRecorder()

	handler := internal.NewHandler()
	handler.ServeHTTP(response, request)

	expectedCode := http.StatusOK
	expectedBody := "hello, test!"

	if response.Code != expectedCode {
		t.Errorf("result: '%d', expected: '%d'", response.Code, expectedCode)
	}

	if response.Body.String() != expectedBody {
		t.Errorf("result: '%s', expected: '%s'", response.Body.String(), expectedBody)
	}
}

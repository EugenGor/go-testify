package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func getResponseRecorder(url string) *httptest.ResponseRecorder {
	req := httptest.NewRequest("GET", url, nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	return responseRecorder
}
func TestMainHandlerWhenOk(t *testing.T) {
	responseRecorder := getResponseRecorder("/cafe?count=2&city=moscow")
	expectedCode := http.StatusOK
	cnt := 2
	assert.Equal(t, expectedCode, responseRecorder.Code)
	require.NotEmpty(t, responseRecorder.Body)
	body := strings.Split(responseRecorder.Body.String(), ",")
	assert.Len(t, body, cnt)
}
func TestMainHandlerWhenWrongCity(t *testing.T) {
	responseRecorder := getResponseRecorder("/cafe?count=2&city=omsk")
	expectedCode := http.StatusBadRequest
	assert.Equal(t, expectedCode, responseRecorder.Code)
	require.NotEmpty(t, responseRecorder.Body)
	body := responseRecorder.Body.String()
	expectedBody := "wrong city value"
	assert.Equal(t, expectedBody, body)
}
func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	responseRecorder := getResponseRecorder("/cafe?count=10&city=moscow")
	expectedCode := http.StatusOK
	cnt := 4
	assert.Equal(t, expectedCode, responseRecorder.Code)
	require.NotEmpty(t, responseRecorder.Body)
	body := strings.Split(responseRecorder.Body.String(), ",")
	assert.Len(t, body, cnt)
}

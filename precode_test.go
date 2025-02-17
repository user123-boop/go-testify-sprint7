package main

import (
	"net/http"
	"net/http/httptest"

	//"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenBodyOkAndNotEmpty(t *testing.T) {
	//totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// здесь нужно добавить необходимые проверки

	body := responseRecorder.Body.String()                 //тело ответа
	require.NotEmpty(t, body)                              //проверяем пустое ли тело
	require.Equal(t, http.StatusOK, responseRecorder.Code) //проверяем статус ответа
}

func TestMainHandlerWhenWrongCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscownew", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	body := responseRecorder.Body.String()

	bodyB := "wrong city value"
	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	assert.Equal(t, bodyB, body)
}
func TestMainHandlerWhenCountMoreThanTotalCount(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	body := responseRecorder.Body.String()

	sliceCafe := strings.Split(body, ",")
	assert.Len(t, sliceCafe, totalCount)

}

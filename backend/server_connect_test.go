package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServerIsUp(t *testing.T) {
	// Создаем новый запрос
	req, err := http.NewRequest("GET", "http://localhost:4444/api/orders", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Создаем записывающее устройство для записи ответа
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetOrdersHandler) // передача обработчика

	// Вызываем обработчик
	handler.ServeHTTP(rr, req)

	// Проверяем статус-код
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("expected status OK; got %v", status)
	} else {
		t.Logf("Test passed: received status %v", status) // Логируем успешный статус
	}
}

package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type Response struct {
	Message string `json:"message"`
}

// TestHelloEndpoint проверяет эндпоинт /hello
func TestHelloEndpoint(t *testing.T) {
	// Создаём тестовый запрос
	req, err := http.NewRequest("GET", "/hello", nil)
	if err != nil {
		t.Fatalf("Ошибка при создании запроса: %v", err)
	}

	// Создаём тестовый ResponseRecorder
	rr := httptest.NewRecorder()

	// Создаём новый сервир и вызываем handler
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)

	// Выполняем запрос
	mux.ServeHTTP(rr, req)

	// 1. Проверка кода ответа (200 OK)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Ожидался статус %d, но получено %d", http.StatusOK, status)
	}

	// 2. Проверка заголовка Content-Type
	contentType := rr.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Ожидался Content-Type 'application/json', но получено '%s'", contentType)
	}

	// 3. Проверка тела ответа
	body := rr.Body.String()

	// Проверка, что тело содержит поле message
	if !strings.Contains(body, "message") {
		t.Errorf("Тело ответа не содержит поле 'message': %s", body)
	}

	// Проверка, что тело содержит "hello, world!"
	if !strings.Contains(body, "hello, world!") {
		t.Errorf("Тело ответа не содержит 'hello, world!': %s", body)
	}

	// Дополнительная проверка: парсинг JSON и проверка значения
	var response Response
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Errorf("Ошибка при парсинге JSON: %v", err)
	}

	if response.Message != "hello, world!" {
		t.Errorf("Ожидалось message 'hello, world!', но получено '%s'", response.Message)
	}
}

// TestHelloEndpointMethod проверяет, что только GET запросы работают
func TestHelloEndpointMethod(t *testing.T) {
	// Создаём POST запрос (должен вернуть 404 или другой статус)
	req, err := http.NewRequest("POST", "/hello", nil)
	if err != nil {
		t.Fatalf("Ошибка при создании запроса: %v", err)
	}

	rr := httptest.NewRecorder()

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)
	mux.ServeHTTP(rr, req)

	// POST запрос не должен обрабатываться (вернёт 404)
	if status := rr.Code; status == http.StatusOK {
		t.Errorf("POST запрос к /hello не должен возвращать 200 OK, но получил статус %d", status)
	}
}

// TestHelloEndpointPath проверяет, что только /hello работает
func TestHelloEndpointPath(t *testing.T) {
	// Создаём запрос к несуществующему эндпоинту
	req, err := http.NewRequest("GET", "/notfound", nil)
	if err != nil {
		t.Fatalf("Ошибка при создании запроса: %v", err)
	}

	rr := httptest.NewRecorder()

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)
	mux.ServeHTTP(rr, req)

	// Несуществующий путь должен вернуть 404
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("Ожидался статус 404 для несуществующего пути, но получено %d", status)
	}
}

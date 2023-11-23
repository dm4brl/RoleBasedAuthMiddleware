package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAdmin(t *testing.T) {
	request, err := http.NewRequest("GET", "/admin", nil)
	if err != nil {
		t.Fatal(err)
	}
	request.Header.Set("X-User-Role", "admin")

	// Создание виртуального сервера.
	recorder := httptest.NewRecorder()
	handler := http.Handler(RoleBasedAuthMiddleware([]string{"admin", "superadmin"}, http.HandlerFunc(AdminHandler)))
	handler.ServeHTTP(recorder, request)

	// Проверка кода состояния и ответа.
	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, recorder.Code)
	}

	// Проверка тела ответа.
	expectedResponse := "Admin Resource"
	body, err := ioutil.ReadAll(recorder.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(body) != expectedResponse {
		t.Errorf("Incorrect body. Expected: %s, got: %s", expectedResponse, body)
	}
}

func TestUser(t *testing.T) {
	request, err := http.NewRequest("GET", "/user", nil)
	if err != nil {
		t.Fatal(err)
	}
	request.Header.Set("X-User-Role", "user")

	// Создание виртуального сервера.
	recorder := httptest.NewRecorder()
	handler := http.Handler(RoleBasedAuthMiddleware([]string{"user"}, http.HandlerFunc(UserHandler)))
	handler.ServeHTTP(recorder, request)

	// Проверка кода состояния и ответа.
	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, recorder.Code)
	}

	// Проверка тела ответа.
	expectedResponse := "User Resource"
	body, err := ioutil.ReadAll(recorder.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(body) != expectedResponse {
		t.Errorf("Incorrect body. Expected: %s, got: %s", expectedResponse, body)
	}
}

func TestForbidden(t *testing.T) {
	request, err := http.NewRequest("GET", "/admin", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Создание виртуального сервера.
	recorder := httptest.NewRecorder()
	handler := http.Handler(RoleBasedAuthMiddleware([]string{"user"}, http.HandlerFunc(AdminHandler)))
	handler.ServeHTTP(recorder, request)

	// Проверка кода состояния (должен быть http.StatusForbidden).
	if recorder.Code != http.StatusForbidden {
		t.Errorf("Expected status code %d, got %d", http.StatusForbidden, recorder.Code)
	}
}

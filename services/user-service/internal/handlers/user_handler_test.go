package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"user-service/internal/models"
	"user-service/internal/repository"
)

func setupUserHandler() *UserHandler {
	return NewUserHandler(repository.NewInMemoryUserRepository())
}

func TestCreateUser_Success(t *testing.T) {
	h := setupUserHandler()
	body := bytes.NewBufferString(`{"name":"Test","email":"t@example.com","password":"p"}`)
	req := httptest.NewRequest(http.MethodPost, "/users", body)
	rec := httptest.NewRecorder()

	h.CreateUser(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("expected 201 got %d", rec.Code)
	}
	var resp models.Response
	_ = json.Unmarshal(rec.Body.Bytes(), &resp)
	if !resp.Success {
		t.Error("expected success true")
	}
}

func TestCreateUser_ValidationError(t *testing.T) {
	h := setupUserHandler()
	body := bytes.NewBufferString(`{"name":"","email":"","password":""}`)
	req := httptest.NewRequest(http.MethodPost, "/users", body)
	rec := httptest.NewRecorder()

	h.CreateUser(rec, req)
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected 400 got %d", rec.Code)
	}
}

func TestLogin_InvalidCredentials(t *testing.T) {
	h := setupUserHandler()
	// create a user
	createBody := bytes.NewBufferString(`{"name":"Test","email":"t@example.com","password":"secret"}`)
	creq := httptest.NewRequest(http.MethodPost, "/users", createBody)
	cres := httptest.NewRecorder()
	h.CreateUser(cres, creq)

	// attempt login with wrong password
	loginBody := bytes.NewBufferString(`{"email":"t@example.com","password":"wrong"}`)
	lreq := httptest.NewRequest(http.MethodPost, "/auth/login", loginBody)
	lres := httptest.NewRecorder()
	h.Login(lres, lreq)
	if lres.Code != http.StatusUnauthorized {
		t.Fatalf("expected 401 got %d", lres.Code)
	}
}

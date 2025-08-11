package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"product-service/internal/repository"
)

func setupProductHandler() *ProductHandler {
	return NewProductHandler(repository.NewInMemoryProductRepository())
}

func TestCreateProduct_Success(t *testing.T) {
	h := setupProductHandler()
	body := bytes.NewBufferString(`{"name":"Test","description":"d","category":"Cat","price":10.5,"stock":5}`)
	req := httptest.NewRequest(http.MethodPost, "/products", body)
	rec := httptest.NewRecorder()

	h.CreateProduct(rec, req)
	if rec.Code != http.StatusCreated {
		t.Fatalf("expected 201 got %d", rec.Code)
	}
}

func TestCreateProduct_ValidationError(t *testing.T) {
	h := setupProductHandler()
	body := bytes.NewBufferString(`{"name":"","description":"d","category":"","price":0}`)
	req := httptest.NewRequest(http.MethodPost, "/products", body)
	rec := httptest.NewRecorder()

	h.CreateProduct(rec, req)
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected 400 got %d", rec.Code)
	}
}

func TestListProducts_Filter(t *testing.T) {
	h := setupProductHandler()
	req := httptest.NewRequest(http.MethodGet, "/products?category=Electronics&in_stock=true", nil)
	rec := httptest.NewRecorder()

	h.ListProducts(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200 got %d", rec.Code)
	}
	// Basic validation of JSON structure
	var js map[string]interface{}
	if err := json.Unmarshal(rec.Body.Bytes(), &js); err != nil {
		t.Fatalf("invalid json: %v", err)
	}
}

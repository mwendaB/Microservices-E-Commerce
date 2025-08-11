package handlers

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"order-service/internal/models"
	"order-service/internal/repository"
)

type mockClient struct {
	userErr   error
	itemsErr  error
	items     []models.OrderItem
}

func (m *mockClient) CheckUserExists(userID string) error { return m.userErr }
func (m *mockClient) ValidateOrderItems(items []models.CreateOrderItem) ([]models.OrderItem, error) {
	if m.itemsErr != nil { return nil, m.itemsErr }
	return m.items, nil
}

func TestCreateOrder_Success(t *testing.T) {
	repo := repository.NewInMemoryOrderRepository()
	mock := &mockClient{items: []models.OrderItem{models.NewOrderItem("p1","Prod",10,1)}}
	h := NewOrderHandler(repo, mock)
	body := bytes.NewBufferString(`{"user_id":"u1","items":[{"product_id":"p1","quantity":1}]}`)
	req := httptest.NewRequest(http.MethodPost, "/orders", body)
	rec := httptest.NewRecorder()

	h.CreateOrder(rec, req)
	if rec.Code != http.StatusCreated {
		t.Fatalf("expected 201 got %d", rec.Code)
	}
}

func TestCreateOrder_InvalidUser(t *testing.T) {
	repo := repository.NewInMemoryOrderRepository()
	mock := &mockClient{userErr: errors.New("user not found")}
	h := NewOrderHandler(repo, mock)
	body := bytes.NewBufferString(`{"user_id":"bad","items":[{"product_id":"p1","quantity":1}]}`)
	req := httptest.NewRequest(http.MethodPost, "/orders", body)
	rec := httptest.NewRecorder()

	h.CreateOrder(rec, req)
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected 400 got %d", rec.Code)
	}
}

func TestUpdateOrderStatus_InvalidStatus(t *testing.T) {
	repo := repository.NewInMemoryOrderRepository()
	mock := &mockClient{}
	h := NewOrderHandler(repo, mock)
	// create base order directly
	o := models.NewOrder("u1", []models.OrderItem{models.NewOrderItem("p1","Prod",10,1)})
	_ = repo.Create(o)

	body := bytes.NewBufferString(`{"status":"wrong"}`)
	req := httptest.NewRequest(http.MethodPatch, "/orders/"+o.ID+"/status", body)
	rec := httptest.NewRecorder()

	h.UpdateOrderStatus(rec, req)
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected 400 got %d", rec.Code)
	}
}

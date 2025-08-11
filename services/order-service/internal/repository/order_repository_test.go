package repository

import (
	"testing"
	"order-service/internal/models"
)

func TestInMemoryOrderRepository_CreateAndGet(t *testing.T) {
	repo := NewInMemoryOrderRepository()
	order := models.NewOrder("user1", []models.OrderItem{{ProductID: "p1", Quantity: 2}})
	if err := repo.Create(order); err != nil {
		t.Fatalf("create failed: %v", err)
	}
	got, err := repo.GetByID(order.ID)
	if err != nil {
		t.Fatalf("get failed: %v", err)
	}
	if got.UserID != order.UserID {
		t.Errorf("expected userID %s got %s", order.UserID, got.UserID)
	}
}

func TestInMemoryOrderRepository_GetByUserID_List_Delete(t *testing.T) {
	repo := NewInMemoryOrderRepository()
	o1 := models.NewOrder("u1", []models.OrderItem{{ProductID: "p1", Quantity: 1}})
	o2 := models.NewOrder("u1", []models.OrderItem{{ProductID: "p2", Quantity: 3}})
	o3 := models.NewOrder("u2", []models.OrderItem{{ProductID: "p3", Quantity: 2}})
	_ = repo.Create(o1)
	_ = repo.Create(o2)
	_ = repo.Create(o3)

	u1Orders, err := repo.GetByUserID("u1")
	if err != nil {
		t.Fatalf("GetByUserID failed: %v", err)
	}
	if len(u1Orders) != 2 {
		t.Errorf("expected 2 orders for u1 got %d", len(u1Orders))
	}

	all, _ := repo.List()
	if len(all) != 3 {
		t.Errorf("expected 3 total orders got %d", len(all))
	}

	if err := repo.Delete(o2.ID); err != nil {
		t.Fatalf("delete failed: %v", err)
	}
	if _, err := repo.GetByID(o2.ID); err == nil {
		t.Error("expected error for deleted order")
	}
}

func TestInMemoryOrderRepository_Update(t *testing.T) {
	repo := NewInMemoryOrderRepository()
	o := models.NewOrder("u3", []models.OrderItem{{ProductID: "p9", Quantity: 4}})
	_ = repo.Create(o)
	o.Status = models.OrderStatusConfirmed
	if err := repo.Update(o); err != nil {
		t.Fatalf("update failed: %v", err)
	}
	got, _ := repo.GetByID(o.ID)
	if got.Status != models.OrderStatusConfirmed {
		t.Errorf("expected status confirmed got %s", got.Status)
	}
}

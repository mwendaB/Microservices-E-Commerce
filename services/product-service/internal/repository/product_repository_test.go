package repository

import (
	"testing"
	"product-service/internal/models"
)

func TestInMemoryProductRepository_CreateAndGet(t *testing.T) {
	repo := NewInMemoryProductRepository()
	p := models.NewProduct("Test Product", "Desc", "Category", 10.0, 5, "img")
	if err := repo.Create(p); err != nil {
		t.Fatalf("create failed: %v", err)
	}
	if err := repo.Create(models.NewProduct("Test Product", "Desc2", "Category", 11.0, 2, "img2")); err == nil {
		t.Error("expected duplicate name error")
	}
	got, err := repo.GetByID(p.ID)
	if err != nil {
		t.Fatalf("get failed: %v", err)
	}
	if got.Name != p.Name {
		t.Errorf("expected name %s got %s", p.Name, got.Name)
	}
}

func TestInMemoryProductRepository_Filtering(t *testing.T) {
	repo := NewInMemoryProductRepository()
	_ = repo.Create(models.NewProduct("Cheap", "", "Electronics", 5, 1, ""))
	_ = repo.Create(models.NewProduct("Mid", "", "Electronics", 50, 0, ""))
	_ = repo.Create(models.NewProduct("Expensive", "", "Electronics", 500, 3, ""))

	filter := &models.ProductFilter{MinPrice: 10, MaxPrice: 400, InStock: true, Category: "Electronics"}
	list, err := repo.List(filter)
	if err != nil {
		t.Fatalf("list failed: %v", err)
	}
	for _, p := range list {
		if p.Price < 10 || p.Price > 400 || p.Stock <= 0 || p.Category != "Electronics" {
			t.Error("filter returned invalid product")
		}
	}
}

func TestInMemoryProductRepository_UpdateStock(t *testing.T) {
	repo := NewInMemoryProductRepository()
	p := models.NewProduct("Stock Item", "", "Cat", 9.9, 10, "")
	_ = repo.Create(p)
	if err := repo.UpdateStock(p.ID, 25); err != nil {
		t.Fatalf("update stock failed: %v", err)
	}
	got, _ := repo.GetByID(p.ID)
	if got.Stock != 25 {
		t.Errorf("expected stock 25 got %d", got.Stock)
	}
	if err := repo.UpdateStock(p.ID, -5); err == nil {
		t.Error("expected negative stock error")
	}
}

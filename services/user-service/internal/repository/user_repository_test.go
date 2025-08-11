package repository

import (
	"testing"
	"user-service/internal/models"
)

func TestInMemoryUserRepository_CreateAndGet(t *testing.T) {
	repo := NewInMemoryUserRepository()
	user := models.NewUser("Alice", "alice@example.com", "password123")

	if err := repo.Create(user); err != nil {
		t.Errorf("expected create success, got error: %v", err)
	}

	// Duplicate email
	dup := models.NewUser("Alice2", "alice@example.com", "pass")
	if err := repo.Create(dup); err == nil {
		t.Error("expected duplicate email error, got nil")
	}

	fetched, err := repo.GetByID(user.ID)
	if err != nil {
		t.Fatalf("expected fetch success, got error: %v", err)
	}
	if fetched.Email != user.Email {
		t.Errorf("expected email %s, got %s", user.Email, fetched.Email)
	}
	if fetched.Password != "" { // password should be blanked in GetByID
		t.Error("expected password to be stripped in fetched user")
	}
}

func TestInMemoryUserRepository_UpdateAndDelete(t *testing.T) {
	repo := NewInMemoryUserRepository()
	user := models.NewUser("Bob", "bob@example.com", "password")
	_ = repo.Create(user)

	user.Name = "Bob Updated"
	if err := repo.Update(user); err != nil {
		t.Fatalf("update failed: %v", err)
	}

	fetched, _ := repo.GetByID(user.ID)
	if fetched.Name != "Bob Updated" {
		t.Errorf("expected updated name, got %s", fetched.Name)
	}

	if err := repo.Delete(user.ID); err != nil {
		t.Fatalf("delete failed: %v", err)
	}
	if _, err := repo.GetByID(user.ID); err == nil {
		t.Error("expected error fetching deleted user")
	}
}

func TestInMemoryUserRepository_List(t *testing.T) {
	repo := NewInMemoryUserRepository()
	_ = repo.Create(models.NewUser("A", "a@example.com", "p"))
	_ = repo.Create(models.NewUser("B", "b@example.com", "p"))

	users, err := repo.List()
	if err != nil {
		t.Fatalf("list failed: %v", err)
	}
	if len(users) != 2 {
		t.Errorf("expected 2 users, got %d", len(users))
	}
	for _, u := range users {
		if u.Password != "" {
			t.Error("expected stripped password in list")
		}
	}
}

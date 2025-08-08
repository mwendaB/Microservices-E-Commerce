package repository

import (
	"errors"
	"sync"
	"user-service/internal/models"
)

// UserRepository defines the interface for user data operations
type UserRepository interface {
	Create(user *models.User) error
	GetByID(id string) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	Update(user *models.User) error
	Delete(id string) error
	List() ([]*models.User, error)
}

// InMemoryUserRepository implements UserRepository using in-memory storage
// In production, this would be replaced with a database implementation
type InMemoryUserRepository struct {
	users map[string]*models.User
	mutex sync.RWMutex
}

// NewInMemoryUserRepository creates a new in-memory user repository
func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[string]*models.User),
	}
}

// Create adds a new user to the repository
func (r *InMemoryUserRepository) Create(user *models.User) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	// Check if user with email already exists
	for _, existingUser := range r.users {
		if existingUser.Email == user.Email {
			return errors.New("user with this email already exists")
		}
	}

	r.users[user.ID] = user
	return nil
}

// GetByID retrieves a user by their ID
func (r *InMemoryUserRepository) GetByID(id string) (*models.User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	user, exists := r.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}

	// Return a copy to prevent external modification
	userCopy := *user
	userCopy.Password = "" // Don't return password
	return &userCopy, nil
}

// GetByEmail retrieves a user by their email address
func (r *InMemoryUserRepository) GetByEmail(email string) (*models.User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	for _, user := range r.users {
		if user.Email == email {
			return user, nil // Return with password for authentication
		}
	}

	return nil, errors.New("user not found")
}

// Update modifies an existing user
func (r *InMemoryUserRepository) Update(user *models.User) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.users[user.ID]; !exists {
		return errors.New("user not found")
	}

	r.users[user.ID] = user
	return nil
}

// Delete removes a user from the repository
func (r *InMemoryUserRepository) Delete(id string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.users[id]; !exists {
		return errors.New("user not found")
	}

	delete(r.users, id)
	return nil
}

// List returns all users (without passwords)
func (r *InMemoryUserRepository) List() ([]*models.User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	users := make([]*models.User, 0, len(r.users))
	for _, user := range r.users {
		userCopy := *user
		userCopy.Password = "" // Don't return passwords
		users = append(users, &userCopy)
	}

	return users, nil
}

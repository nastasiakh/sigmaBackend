package repositories

import (
	"sigma/internal/entities"
	"sync"
)

type UserRepository struct {
	mu    sync.RWMutex
	users []entities.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (repo *UserRepository) GetAll() ([]entities.User, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	return repo.users, nil
}

func (repo *UserRepository) Create(user entities.User) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	user.ID = uint(len(repo.users) + 1)
	repo.users = append(repo.users, user)
	return nil
}

func (repo *UserRepository) GetByID(id uint) *entities.User {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	for _, user := range repo.users {
		if user.ID == id {
			return &user
		}
	}

	return nil
}

func (repo *UserRepository) Update(user entities.User) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	for i, existingUser := range repo.users {
		if existingUser.ID == user.ID {
			repo.users[i] = user
			return
		}
	}
}

func (repo *UserRepository) Delete(id uint) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	var updatedUsers []entities.User
	for _, existingUser := range repo.users {
		if existingUser.ID != id {
			updatedUsers = append(updatedUsers, existingUser)
		}
	}

	repo.users = updatedUsers
	return nil
}

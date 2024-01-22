package services

import (
	"sigma/internal/entities"
	"sigma/internal/handlers/repositories"
)

type UserService struct {
	userRepo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{userRepo: repo}
}

func (s *UserService) CreateUserService(user entities.User) error {
	if err := s.userRepo.Create(user); err != nil {
		return err
	}
	return nil
}

func (s *UserService) GetAllUsersService() ([]entities.User, error) {
	return s.userRepo.GetAll()
}

func (s *UserService) GetUserByIDService(id uint) *entities.User {
	return s.userRepo.GetByID(id)
}

func (s *UserService) UpdateUserService(user entities.User) {
	s.userRepo.Update(user)
}

func (s *UserService) DeleteUserService(id uint) error {
	err := s.userRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

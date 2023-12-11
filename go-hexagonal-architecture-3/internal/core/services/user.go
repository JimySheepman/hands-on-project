package services

import (
	"Hexagonal-Architecture/internal/adapters/repository"
	"Hexagonal-Architecture/internal/core/domain"
	"Hexagonal-Architecture/internal/core/ports"
)

type UserService struct {
	repo ports.UserRepository
}

func NewUserService(repo ports.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (u *UserService) CreateUser(email, password string) (*domain.User, error) {
	return u.repo.CreateUser(email, password)
}

func (u *UserService) ReadUser(id string) (*domain.User, error) {
	return u.repo.ReadUser(id)
}

func (u *UserService) ReadUsers() ([]*domain.User, error) {
	return u.repo.ReadUsers()
}

func (u *UserService) UpdateUser(id, email, password string) error {
	return u.repo.UpdateUser(id, email, password)
}

func (u *UserService) DeleteUser(id string) error {
	return u.repo.DeleteUser(id)
}

func (u *UserService) LoginUser(email, password string) (*repository.LoginResponse, error) {
	return u.repo.LoginUser(email, password)
}

func (u *UserService) UpdateMembershipStatus(id string, status bool) error {
	return u.repo.UpdateMembershipStatus(id, status)
}

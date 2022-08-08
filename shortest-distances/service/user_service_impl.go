package service

import (
	"context"
	"shortest-distances/repository"
)

type UserServiceImpl struct {
	repo repository.UserRepository
}

func NewUserServiceImpl(repo repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		repo: repo,
	}
}

func (i *UserServiceImpl) Login(ctx context.Context) error {
	return nil
}

func (i *UserServiceImpl) Register(ctx context.Context) error {
	return nil
}

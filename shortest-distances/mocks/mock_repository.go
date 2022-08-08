package mocks

import "context"

type postgreMockRepository struct {
	login    map[int]error
	register map[int]error
}

func NewMockRepository() *postgreMockRepository {
	return &postgreMockRepository{
		login:    map[int]error{},
		register: map[int]error{},
	}
}

func (r *postgreMockRepository) Login(ctx context.Context) error {
	return nil
}

func (r *postgreMockRepository) Register(ctx context.Context) error {
	return nil
}

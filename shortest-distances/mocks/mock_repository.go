package mocks

type postgreMockRepository struct {
	Login    map[int]error
	Register map[int]error
}

func NewMockRepository() *postgreMockRepository {
	return &postgreMockRepository{
		Login:    map[int]error{},
		Register: map[int]error{},
	}
}

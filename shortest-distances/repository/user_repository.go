package repository

import (
	"context"
)

type UserRepository interface {
	Login(context.Context) error
	Register(context.Context) error
}

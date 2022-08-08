package service

import (
	"context"
)

type UserService interface {
	Login(context.Context) error
	Register(context.Context) error
}

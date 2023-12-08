//go:generate mockgen -package=mockups -destination=../../mocks/mockups/uidgem.go -source=uidgen.go

package uidgen

import "github.com/google/uuid"

type UIDGen interface {
	New() string
}

type uidgen struct{}

func New() UIDGen {
	return &uidgen{}
}

func (u uidgen) New() string {
	return uuid.New().String()
}

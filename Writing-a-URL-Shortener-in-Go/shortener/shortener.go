package shortener

import (
	"fmt"

	"github.com/teris-io/shortid"

	"go-short/storage"
)

type Shortener interface {
	Shorten(url string) (string, error)
}

type Store struct {
	storage storage.Storage
}

var generateID = func() (string, error) {
	return shortid.Generate()
}

func (u *Store) Shorten(url string) (string, error) {
	id, err := generateID()
	if err != nil {
		return "", fmt.Errorf("generating id for %s: %v", url, err)
	}

	exists, err := u.storage.Exists(id)
	if err != nil {
		return "", fmt.Errorf("checking existance of key %s for url %s: %v", id, url, err)
	} else if exists {
		return "", fmt.Errorf("key %s for url %s already exists in storage", id, url)
	}

	if err := u.storage.Set(id, url); err != nil {
		return "", fmt.Errorf("storing key %s for url %s: %v", id, url, err)
	}

	return id, nil
}

func (u Store) Get(key string) (string, error) {
	return u.storage.Get(key)
}

func New(storage storage.Storage) *Store {
	return &Store{storage: storage}
}

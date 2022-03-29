package shortener

import (
	"errors"
	"fmt"
	"testing"
)

type fakeStorage struct {
	storage map[string]string
}

func (f fakeStorage) Get(key string) (string, error) {
	if key == "ERROR" {
		return "", fmt.Errorf("error")
	}

	return f.storage[key], nil
}

func (f fakeStorage) Exists(key string) (bool, error) {
	if key == "ERROR" {
		return false, fmt.Errorf("error")
	}

	_, ok := f.storage[key]
	return ok, nil
}

func (f fakeStorage) Set(key string, value interface{}) error {
	if key == "ERROR" {
		return fmt.Errorf("error")
	}

	switch val := value.(type) {
	case string:
		f.storage[key] = val
		return nil
	default:
		return fmt.Errorf("Unable to store: %v", val)
	}
}

func createURLStorage(prefill map[string]string) *Store {
	if prefill == nil {
		prefill = make(map[string]string)
	}

	fakeStorage := fakeStorage{prefill}
	return &Store{fakeStorage}
}

func TestShortenGenerateID(t *testing.T) {
	tests := []struct {
		key         string
		generateID  func() (string, error)
		shouldError bool
	}{
		{"should succeed", func() (string, error) { return "KEY", nil }, false},
		{"should error", func() (string, error) { return "", errors.New("Error") }, true},
	}

	idGenerator := generateID

	for _, test := range tests {
		store := createURLStorage(nil)
		generateID = test.generateID

		_, err := store.Shorten(test.key)
		if test.shouldError && err == nil {
			t.Errorf("store.Shorten(%s) expected error, found none", test.key)
		}

		if !test.shouldError && err != nil {
			t.Errorf("store.Shorten(%s) found error, expected none: %v", test.key, err)
		}
	}

	generateID = idGenerator
}
func TestShortenShouldErrorWhenKeyExistsOrExistsFails(t *testing.T) {

}

func TestShortenShouldErrorWhenSetFails(t *testing.T) {

}

func TestGet(t *testing.T) {
	tests := []struct {
		key         string
		url         string
		shouldError bool
	}{
		{"exists", "myurl", false},
		{"ERROR", "", true},
	}

	for _, test := range tests {
		store := createURLStorage(map[string]string{test.key: test.url})

		url, err := store.Get(test.key)
		if test.shouldError && err == nil {
			t.Errorf("store.Get(%s) expected error, found none", test.key)
			continue
		} else if !test.shouldError && err != nil {
			t.Errorf("store.Get(%s) found error, expected none", test.key)
			continue
		}

		if url != test.url {
			t.Errorf("store.Get(%s)=%s, expected %s", test.key, url, test.url)
		}
	}
}

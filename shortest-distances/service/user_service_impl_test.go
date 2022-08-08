package service

import (
	"shortest-distances/mocks"
	"testing"
)

func TestLogin(t *testing.T) {
	r := mocks.NewMockRepository()
	s := NewUserServiceImpl(r)
	_ = s

	tests := []struct {
		testName string
	}{
		{testName: "test"},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			t.Log(test.testName)
		})
	}
}

func TestRegister(t *testing.T) {
	r := mocks.NewMockRepository()
	s := NewUserServiceImpl(r)
	_ = s

	tests := []struct {
		testName string
	}{
		{testName: "test"},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			t.Log(test.testName)
		})
	}
}

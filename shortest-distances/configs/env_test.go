package configs

import (
	"errors"
	"reflect"
	"testing"
)

func TestConfigLoad(t *testing.T) {
	var tests = []struct {
		testName string
		path     string
		expected error
	}{
		{
			testName: "loading error",
			path:     "",
			expected: errors.New("Error loading .env file"),
		},
		{
			testName: "successful loading",
			path:     "../.env",
			expected: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			actual := ConfigLoad(test.path)
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("actual:%v \nexpected:%v", actual, test.expected)
			}
		})
	}
}

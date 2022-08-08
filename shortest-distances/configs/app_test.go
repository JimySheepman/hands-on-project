package configs

import (
	"reflect"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
)

func TestAppConfig(t *testing.T) {

	type Views interface{}
	type Ctx struct{}
	type ErrorHandler = func(*Ctx, error) error

	var tests = []struct {
		Config fiber.Config
	}{
		{
			Config: fiber.Config{
				AppName:       "dispatcher-api",
				ReadTimeout:   time.Second * time.Duration(30),
				WriteTimeout:  time.Second * time.Duration(30),
				CaseSensitive: true,
				BodyLimit:     64 * 1024 * 1024,
				Concurrency:   256 * 1024,
				IdleTimeout:   10 * time.Second,
			},
		},
	}

	for _, test := range tests {
		t.Run("successful config", func(t *testing.T) {
			actual := AppConfig()
			expected := test.Config

			if !reflect.DeepEqual(actual, expected) {
				t.Errorf("actual:%v \nexpected:%v", actual, expected)
			}
		})
	}
}

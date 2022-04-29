package routes

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/twinj/uuid"
)

func TestPublicRoutes(t *testing.T) {
	if err := godotenv.Load("../../.env.test"); err != nil {
		panic(err)
	}

	tests := []struct {
		description   string
		route         string
		expectedError bool
		expectedCode  int
	}{
		{
			description:   "get book by ID",
			route:         "/api/v1/book/" + uuid.New().String(),
			expectedError: false,
			expectedCode:  404,
		},
		{
			description:   "get book by invalid ID (non UUID)",
			route:         "/api/v1/book/123456",
			expectedError: false,
			expectedCode:  500,
		},
	}

	app := fiber.New()

	PublicRoutes(app)

	for _, test := range tests {
		req := httptest.NewRequest("GET", test.route, nil)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req, -1) // the -1 disables request latency

		assert.Equalf(t, test.expectedError, err != nil, test.description)

		if test.expectedError {
			continue
		}

		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}

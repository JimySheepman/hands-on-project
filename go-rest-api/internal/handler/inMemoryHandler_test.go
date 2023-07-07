package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/JimySheepman/go-rest-api/internal/model"
)

type TestMemoryRequestPayload struct {
	Key   string `json:"key"`
	Value int    `json:"value"`
}

func TestPostInMemeoryDataHandler(t *testing.T) {
	t.Run("In-memory POST Endpoint", func(t *testing.T) {
		testBody := &model.MemoryRequestPayload{
			Key:   "active-tabs",
			Value: "getir",
		}

		body, _ := json.Marshal(testBody)
		req, err := http.NewRequest(http.MethodPost, "/api/v1/in-memory", strings.NewReader(string(body)))
		if err != nil {
			t.Errorf("Request creation failed: ERROR: %v", err)
		}

		res := httptest.NewRecorder()
		handler := http.HandlerFunc(PostInMemeoryDataHandler())
		handler.ServeHTTP(res, req)

		expectedResponse := model.MemoryRequestPayload{
			Key:   "active-tabs",
			Value: "getir",
		}

		marshalExpectedResponse, _ := json.Marshal(expectedResponse)
		// Because of the "\n"  taken when reading res.Body.String() length one more room than expected length. We add "\n" .
		expected := string(marshalExpectedResponse) + "\n"
		if res.Body.String() != expected {
			t.Errorf("Handler returned unexpected body:\n got: %v \nwant: %v", res.Body.String(), string(marshalExpectedResponse))
		}
	})

	t.Run("unmarshal request body", func(t *testing.T) {
		testBody := &TestMemoryRequestPayload{
			Key:   "active-tabs",
			Value: 1,
		}

		body, _ := json.Marshal(testBody)
		req, err := http.NewRequest(http.MethodPost, "/api/v1/in-memory", strings.NewReader(string(body)))
		if err != nil {
			t.Errorf("Request creation failed: ERROR: %v", err)
		}

		res := httptest.NewRecorder()
		handler := http.HandlerFunc(PostInMemeoryDataHandler())
		handler.ServeHTTP(res, req)
		expectedResponse := model.MemoryErrorResponsePayload{
			Code:    2,
			Message: "Error: could not complete unmarshal body",
		}

		marshalExpectedResponse, _ := json.Marshal(expectedResponse)
		// Because of the "\n"  taken when reading res.Body.String() length one more room than expected length. We add "\n" .
		expected := string(marshalExpectedResponse) + "\n"
		if res.Body.String() != expected {
			t.Errorf("Handler returned unexpected body:\n got: %v \nwant: %v", res.Body.String(), string(marshalExpectedResponse))
		}
	})
}

func TestGetInMemeoryDataHandler(t *testing.T) {
	t.Run("In-memory GET Endpoint", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/v1/in-memory?key=active-tabs", nil)
		res := httptest.NewRecorder()
		handler := http.HandlerFunc(GetInMemeoryDataHandler())
		handler.ServeHTTP(res, req)

		expectedResponse := model.MemoryRequestPayload{
			Key:   req.FormValue("key"),
			Value: "getir",
		}

		marshalExpectedResponse, _ := json.Marshal(expectedResponse)
		// Because of the "\n"  taken when reading res.Body.String() length one more room than expected length. We add "\n" .
		expected := string(marshalExpectedResponse) + "\n"
		if res.Body.String() != expected {
			t.Errorf("Handler returned unexpected body:\n got: %v \nwant: %v", res.Body.String(), string(marshalExpectedResponse))
		}
	})

	t.Run("In-memory GET Endpoint", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/v1/in-memory?value=active-tabs", nil)
		res := httptest.NewRecorder()
		handler := http.HandlerFunc(GetInMemeoryDataHandler())
		handler.ServeHTTP(res, req)

		expectedResponse := model.MemoryErrorResponsePayload{
			Code:    3,
			Message: "Error: Url Param 'key' is missing",
		}

		marshalExpectedResponse, _ := json.Marshal(expectedResponse)
		// Because of the "\n"  taken when reading res.Body.String() length one more room than expected length. We add "\n" .
		expected := string(marshalExpectedResponse) + "\n"
		if res.Body.String() != expected {
			t.Errorf("Handler returned unexpected body:\n got: %v \nwant: %v", res.Body.String(), string(marshalExpectedResponse))
		}
	})
}

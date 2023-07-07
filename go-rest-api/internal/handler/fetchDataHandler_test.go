package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/JimySheepman/go-rest-api/config/env"
	"github.com/JimySheepman/go-rest-api/internal/model"
)

func init() {
	_, err := env.LoadEnvironmentConfigure("../../.env")
	if err != nil {
		log.Fatal("Loading .env file failed")
	}
}

type TestRecordsRequestPayload struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	MinCount  int    `json:"minCount"`
	MaxCount  string `json:"maxCount"`
}

func TestPostFetchDataHandler(t *testing.T) {

	t.Run("status method allowed POST", func(t *testing.T) {
		testBody := &model.RecordsRequestPayload{
			StartDate: "2016-01-26",
			EndDate:   "2018-02-02",
			MinCount:  2700,
			MaxCount:  3000,
		}

		body, _ := json.Marshal(testBody)
		req, err := http.NewRequest(http.MethodPost, "/api/v1/fetch-data", strings.NewReader(string(body)))
		if err != nil {
			t.Errorf("Request creation failed: ERROR: %v", err)
		}

		res := httptest.NewRecorder()
		handler := assertHandler()
		handler.ServeHTTP(res, req)

		if req.Method != "POST" {
			t.Errorf("Request method is not 'POST': got\n %v want\n %v", req.Method, http.MethodPost)
		}
	})

	t.Run("status method not allowed GET", func(t *testing.T) {
		testBody := &model.RecordsRequestPayload{
			StartDate: "2016-01-26",
			EndDate:   "2018-02-02",
			MinCount:  2700,
			MaxCount:  3000,
		}

		body, _ := json.Marshal(testBody)
		req, err := http.NewRequest(http.MethodGet, "/api/v1/fetch-data", strings.NewReader(string(body)))
		if err != nil {
			t.Errorf("Request creation failed: ERROR: %v", err)
		}

		res := httptest.NewRecorder()
		handler := assertHandler()
		handler.ServeHTTP(res, req)

		if req.Method != "GET" {
			t.Errorf("Request method is not 'POST': got\n %v want\n %v", req.Method, http.MethodPost)
		}
	})

	t.Run("succsess result", func(t *testing.T) {
		testBody := &model.RecordsRequestPayload{
			StartDate: "2015-01-26",
			EndDate:   "2016-02-01",
			MinCount:  3000,
			MaxCount:  3010,
		}

		body, _ := json.Marshal(testBody)
		req, err := http.NewRequest(http.MethodPost, "/api/v1/fetch-data", strings.NewReader(string(body)))
		if err != nil {
			t.Errorf("Request creation failed: ERROR: %v", err)
		}

		res := httptest.NewRecorder()
		handler := assertHandler()
		handler.ServeHTTP(res, req)

		stringToTime, _ := time.Parse(time.RFC3339, "2015-01-19T14:27:54.01Z")
		expectedResponse := &model.RecordsResponsePayload{
			Code:    0,
			Message: "Succsess",
			Records: []model.Record{
				{
					Key:        "aCnXSuEJ",
					CreatedAt:  stringToTime,
					TotalCount: 3007,
				},
			},
		}

		marshalExpectedResponse, _ := json.Marshal(expectedResponse)
		// Because of the "\n"  taken when reading res.Body.String() length one more room than expected length. We add "\n" .
		expected := string(marshalExpectedResponse) + "\n"
		if res.Body.String() != expected {
			t.Errorf("Handler returned unexpected body:\n got: %v \nwant: %v", res.Body.String(), string(marshalExpectedResponse))
		}
	})

	t.Run("could not complete unmarshal body", func(t *testing.T) {
		testBody := &TestRecordsRequestPayload{
			StartDate: "2016-01-26",
			EndDate:   "2018-02-02",
			MinCount:  2700,
			MaxCount:  "3000",
		}

		body, _ := json.Marshal(testBody)
		req, err := http.NewRequest(http.MethodPost, "/api/v1/fetch-data", strings.NewReader(string(body)))
		if err != nil {
			t.Errorf("Request creation failed: ERROR: %v", err)
		}

		res := httptest.NewRecorder()
		handler := assertHandler()
		handler.ServeHTTP(res, req)

		expectedResponse := &model.RecordsResponsePayload{
			Code:    2,
			Message: "Error: could not complete unmarshal body",
			Records: []model.Record{},
		}

		marshalExpectedResponse, _ := json.Marshal(expectedResponse)
		// Because of the "\n"  taken when reading res.Body.String() length one more room than expected length. We add "\n" .
		expected := string(marshalExpectedResponse) + "\n"

		if res.Body.String() != expected {
			t.Errorf("Handler returned unexpected body:\n got: %v \nwant:%v", res.Body.String(), string(marshalExpectedResponse))
		}
	})

	t.Run("wrong time format", func(t *testing.T) {
		testBody := &model.RecordsRequestPayload{
			StartDate: "2016-01-26",
			EndDate:   "2018-2-02",
			MinCount:  2700,
			MaxCount:  3000,
		}

		body, _ := json.Marshal(testBody)
		req, err := http.NewRequest(http.MethodPost, "/api/v1/fetch-data", strings.NewReader(string(body)))
		if err != nil {
			t.Errorf("Request creation failed: ERROR: %v", err)
		}

		res := httptest.NewRecorder()
		handler := assertHandler()
		handler.ServeHTTP(res, req)

		expectedResponse := &model.RecordsResponsePayload{
			Code:    3,
			Message: "Error: wrong time format ",
			Records: []model.Record{},
		}
		marshalExpectedResponse, _ := json.Marshal(expectedResponse)
		expected := string(marshalExpectedResponse) + "\n"

		if res.Body.String() != expected {
			t.Errorf("Handler returned unexpected body: got\n %v want\n %v", res.Body.String(), expected)
		}
	})

}

func assertHandler() http.HandlerFunc {
	handler := http.HandlerFunc(PostFetchDataHandler())
	return handler
}

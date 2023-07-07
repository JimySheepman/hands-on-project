package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/JimySheepman/go-rest-api/internal/model"
)

func PostInMemeoryDataHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var memoryRequestPayload model.MemoryRequestPayload

		req, err := ioutil.ReadAll(r.Body)
		if err != nil {
			json.NewEncoder(w).Encode(model.MemoryErrorResponsePayload{
				Code:    1,
				Message: "Error: could not complete read from request body",
			})
			log.Println("Could not complete read from request body")
			return
		}

		err = json.Unmarshal(req, &memoryRequestPayload)
		if err != nil {
			json.NewEncoder(w).Encode(model.MemoryErrorResponsePayload{
				Code:    2,
				Message: "Error: could not complete unmarshal body",
			})
			log.Println("Could not complete unmarshal body")
			return
		}

		json.NewEncoder(w).Encode(memoryRequestPayload)
	}
}

func GetInMemeoryDataHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.FormValue("key")
		if key == "" {
			json.NewEncoder(w).Encode(model.MemoryErrorResponsePayload{
				Code:    3,
				Message: "Error: Url Param 'key' is missing",
			})
			log.Println("Url Param 'key' is missing")
			return
		}

		json.NewEncoder(w).Encode(model.MemoryRequestPayload{
			Key:   key,
			Value: "getir",
		})
	}
}

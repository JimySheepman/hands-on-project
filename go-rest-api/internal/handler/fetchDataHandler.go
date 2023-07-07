package handler

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/JimySheepman/go-rest-api/config/db"
	"github.com/JimySheepman/go-rest-api/internal/model"
	"github.com/JimySheepman/go-rest-api/internal/times"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func PostFetchDataHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var recordsRequestPayload model.RecordsRequestPayload

		req, err := ioutil.ReadAll(r.Body)
		if err != nil {
			json.NewEncoder(w).Encode(model.RecordsResponsePayload{
				Code:    1,
				Message: "Error: could not complete read from request body",
				Records: []model.Record{},
			})
			log.Println("Could not complete read from request body")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = json.Unmarshal(req, &recordsRequestPayload)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(model.RecordsResponsePayload{
				Code:    2,
				Message: "Error: could not complete unmarshal body",
				Records: []model.Record{},
			})
			log.Println("Could not complete unmarshal body")
			return
		}

		err = CheckTimeValidation(w, recordsRequestPayload)
		if err != nil {
			log.Println(err)
			return
		}

		results := DataAggregationFromDatabase(recordsRequestPayload)

		json.NewEncoder(w).Encode(model.RecordsResponsePayload{
			Code:    0,
			Message: "Succsess",
			Records: results,
		})
	}
}

func CheckTimeValidation(w http.ResponseWriter, recordsRequestPayload model.RecordsRequestPayload) error {
	if times.TimeFormatValidator(recordsRequestPayload.StartDate) || times.TimeFormatValidator(recordsRequestPayload.EndDate) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.RecordsResponsePayload{
			Code:    3,
			Message: "Error: wrong time format ",
			Records: []model.Record{},
		})
		return errors.New("Wrong time format ")
	}
	return nil
}

func DataAggregationFromDatabase(recordsRequestPayload model.RecordsRequestPayload) []model.Record {

	startDate := times.TimeConverter(recordsRequestPayload.StartDate)
	endDate := times.TimeConverter(recordsRequestPayload.EndDate)

	matchStartDate := bson.D{{"$match", bson.D{{"createdAt", bson.D{{"$gte", startDate}}}}}}
	matchEndDate := bson.D{{"$match", bson.D{{"createdAt", bson.D{{"$lt", endDate}}}}}}
	unwindCounts := bson.D{{"$unwind", "$counts"}}
	groupCount := bson.D{{"$group", bson.D{{"_id", bson.D{{"key", "$key"}, {"createdAt", "$createdAt"}}}, {"totalCount", bson.D{{"$sum", "$counts"}}}}}}
	matchMinCount := bson.D{{"$match", bson.D{{"totalCount", bson.D{{"$gte", recordsRequestPayload.MinCount}}}}}}
	matchMaxCount := bson.D{{"$match", bson.D{{"totalCount", bson.D{{"$lt", recordsRequestPayload.MaxCount}}}}}}
	projectQuery := bson.D{{"$project", bson.D{{"_id", 0}, {"key", "$_id.key"}, {"createdAt", "$_id.createdAt"}, {"totalCount", 1}}}}

	db, err := db.ConnectDB()
	if err != nil {
		log.Fatal("Cannot connect to database")
	}

	coll := db.Collection("records")

	cursor, err := coll.Aggregate(context.TODO(), mongo.Pipeline{matchStartDate, matchEndDate, unwindCounts, groupCount, matchMinCount, matchMaxCount, projectQuery})
	if err != nil {
		log.Println(err)
	}

	var results []model.Record
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Println(err)
	}

	return results
}

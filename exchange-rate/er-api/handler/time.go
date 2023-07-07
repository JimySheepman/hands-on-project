package handler

import (
	"encoding/json"
	"er-api/db"
	"er-api/model"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"time"
)

const (
	layoutISO = "2006-01-02"
	layoutUS  = "January 2, 2006"
)

func FetchTimeCurrencyHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var timeRequestPayload model.TimeRequestPayload

	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorTemplate{
			Result:    "Error",
			ErrorType: "Could not complete read from request body",
		})
		log.Println("Could not complete read from request body")
		return
	}

	err = json.Unmarshal(req, &timeRequestPayload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorTemplate{
			Result:    "Error",
			ErrorType: "Could not complete unmarshal body",
		})
		log.Println("Could not complete unmarshal body")
		return
	}

	if TimeFormatValidator(timeRequestPayload.StartDate) || TimeFormatValidator(timeRequestPayload.EndDate) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorTemplate{
			Result:    "Error",
			ErrorType: "Wrong time format",
		})
		return
	}

	db, err := db.ConnectDB()
	if err != nil {
		log.Fatal("Cannot connect to database")
	}

	var currencies model.Currencies
	var currencyArray []model.Currencies

	query := `select * from currency where created_at>=$1 and created_at<=$2;`
	rows, err := db.Query(query, TimeConverter(timeRequestPayload.StartDate), TimeConverter(timeRequestPayload.EndDate))
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	for rows.Next() {

		if err := rows.Scan(&currencies.Id, &currencies.BaseCode, &currencies.TargetCode, &currencies.ConversionRate, &currencies.CreatedAt); err != nil {
			log.Fatal(err)
		}
		currencyArray = append(currencyArray, currencies)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	if len(currencyArray) == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorTemplate{
			Result:    "Error",
			ErrorType: "You entered a wrong query parameter",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(currencyArray)
	return
}

func TimeConverter(date string) int64 {
	t, _ := time.Parse(layoutISO, date)
	return t.Unix()
}

func TimeFormatValidator(str string) bool {
	re := regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)

	submatchall := re.FindAllString(str, -1)
	if len(submatchall) == 0 {
		return true
	}
	return false
}

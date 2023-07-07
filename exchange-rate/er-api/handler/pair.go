package handler

import (
	"encoding/json"
	"er-api/db"
	"er-api/model"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func FetchPairConversionHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	baseCode := vars["base"]
	targetCode := vars["target"]
	if baseCode == "" || targetCode == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorTemplate{
			Result:    "Error",
			ErrorType: "You entered a missing query parameter",
		})
		return
	}

	db, err := db.ConnectDB()
	if err != nil {
		log.Fatal("Cannot connect to database")
	}

	var currencies model.Currencies
	var currencyArray []model.Currencies

	query := `select * from currency where base_code=$1 and target_code=$2;`
	rows, err := db.Query(query, "USD", "TRY")
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

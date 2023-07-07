package task

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"er-api-consumer/model"
)

const (
	MAIN_PATH   = "https://v6.exchangerate-api.com/v6/"
	LATEST_PATH = "/latest/"
	PAIR_PATH   = "/pair/"
)

var queryPaths = [3]string{"TRY", "USD", "EUR"}

func Task() {
	for _, path := range queryPaths {
		fecthValue := GetAllValueFromAPI(LATEST_PATH, path)
		unmarshalData := UnmarshalValueForRates(fecthValue)
		SentToQueueByPairvalue(unmarshalData)
	}
}

func GetAllValueFromAPI(subPath string, path string) []byte {
	res, err := http.Get(MAIN_PATH + os.Getenv("API_KEY") + subPath + path)
	if err != nil {
		log.Fatal(err)
	}

	req, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Could not complete read from request body")
	}

	return req
}

func UnmarshalValueForRates(req []byte) model.Rates {
	var rates model.Rates
	err := json.Unmarshal(req, &rates)
	if err != nil {
		log.Fatal("Could not complete unmarshal body")
	}

	return rates
}

func UnmarshalValueForCurrencies(req []byte) model.Currencies {
	var currencies model.Currencies
	err := json.Unmarshal(req, &currencies)
	if err != nil {
		log.Fatal("Could not complete unmarshal body")
	}

	return currencies
}

func SentToQueueByPairvalue(rates model.Rates) {
	for target, conversion_rate := range rates.ConversionRates {
		data := model.Currencies{
			BaseCode:       rates.BaseCode,
			TargetCode:     target,
			ConversionRate: conversion_rate,
			CreatedAt:      time.Now().UTC().Unix(),
		}

		Send(data, rates.BaseCode)
	}
}

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

type CovidResult struct {
	City                        string  `json:"city"`
	CityIbgeCode                string  `json:"city_ibge_code"`
	Confirmed                   int     `json:"confirmed"`
	ConfirmedPer100KInhabitants float64 `json:"confirmed_per_100k_inhabitants"`
	Date                        string  `json:"date"`
	DeathRate                   float64 `json:"death_rate"`
	Deaths                      int     `json:"deaths"`
	EstimatedPopulation         int     `json:"estimated_population"`
	EstimatedPopulation2019     int     `json:"estimated_population_2019"`
	IsLast                      bool    `json:"is_last"`
	OrderForPlace               int     `json:"order_for_place"`
	PlaceType                   string  `json:"place_type"`
	State                       string  `json:"state"`
}

type CovidCase struct {
	Count   int64         `json:"count"`
	Next    string        `json:"next"`
	Prev    string        `json:"prev"`
	Results []CovidResult `json:"results"`
}

func main() {

	os.Setenv("APP_PORT", ":8080")
	APP_PORT := os.Getenv("APP_PORT")
	http.HandleFunc("/covid", GetCovidHandler)
	error := http.ListenAndServe(APP_PORT, nil)

	if error != nil {
		fmt.Fprintf(os.Stderr, "Falha ao executar servidor")
	}

	os.Unsetenv("APP_PORT")
}

func GetCovidHandler(response http.ResponseWriter, request *http.Request) {

	page := request.URL.Query().Get("page")
	covidData, error := GetCovidData(page)

	if error != nil {
		fmt.Fprintf(os.Stderr, "Falha ao consultar dados covid 19")
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.Header().Set("X-Amzn-Idempotency", "a1b2c3-d4e5f6")
	json.NewEncoder(response).Encode(covidData)
}

func GetCovidData(page string) (*CovidCase, error) {

	url := fmt.Sprintf("https://brasil.io/api/v1/dataset/covid19/caso/data/?page=%v", page)

	header := http.Header{}
	header.Set("Authorization", fmt.Sprintf("Token %v", os.Getenv("BRASIL_API_KEY")))
	client := &http.Client{}
	req, httpError := http.NewRequest("GET", url, nil)
	req.Header = header

	if httpError != nil {
		return nil, errors.New("falha ao configurar cliente http")
	}

	res, clientError := client.Do(req)

	if clientError != nil {
		return nil, errors.New("falha ao consultar api covid 19")
	}

	defer res.Body.Close()

	bytes, ioError := io.ReadAll(res.Body)
	if ioError != nil {
		return nil, errors.New("falha ao consumir resultado da requisição")
	}

	covidCase := CovidCase{}
	json.Unmarshal(bytes, &covidCase)
	return &covidCase, nil
}

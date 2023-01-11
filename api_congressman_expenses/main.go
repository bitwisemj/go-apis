package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

type CongressmanExpense struct {
	AuthorName  string `json:"author_name"`
	AuthorURL   string `json:"author_url"`
	CodeURL     string `json:"code_url"`
	Description string `json:"description"`
	ID          string `json:"id"`
	LicenseName string `json:"license_name"`
	LicenseURL  string `json:"license_url"`
	Links       []struct {
		Title string `json:"title"`
		URL   string `json:"url"`
	} `json:"links"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	SourceName  string `json:"source_name"`
	SourceURL   string `json:"source_url"`
	CollectedAt string `json:"collected_at"`
	Tables      []struct {
		Fields []struct {
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"fields"`
		Name       string `json:"name"`
		DataURL    string `json:"data_url"`
		ImportDate string `json:"import_date"`
	} `json:"tables"`
}

type APICall struct{}

func main() {

	os.Setenv("APP_PORT", ":8080")
	APP_PORT := os.Getenv("APP_PORT")

	serveMux := http.NewServeMux()
	serveMux.Handle("/congressman", APICall{})
	http.ListenAndServe(APP_PORT, serveMux)
}

func (apiCall APICall) ServeHTTP(response http.ResponseWriter, request *http.Request) {

	congressmanExpense, error := GetCongressmanExpenses()

	if error != nil {
		panic(error)
	}

	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(congressmanExpense)
}

func GetCongressmanExpenses() (*CongressmanExpense, error) {

	url := "https://brasil.io/api/v1/dataset/gastos-deputados/"

	request, httpError := http.NewRequest("GET", url, nil)

	if httpError != nil {
		return nil, errors.New("falha ao configurar requisição")
	}

	request.Header.Set("Authorization", fmt.Sprintf("Token %v", os.Getenv("BRASIL_API_KEY")))
	client := http.Client{}
	response, requestError := client.Do(request)

	if requestError != nil {
		return nil, errors.New("falha ao executar requisição http")
	}

	defer response.Body.Close()

	bytes, error := io.ReadAll(response.Body)

	if error != nil {
		return nil, errors.New("falha ao consumir body de resposta")
	}

	congressmanExpense := CongressmanExpense{}
	json.Unmarshal(bytes, &congressmanExpense)

	return &congressmanExpense, nil
}

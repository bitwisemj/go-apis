package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

type DogResult struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

type DogRequest struct {
}

func main() {

	os.Setenv("APP_PORT", ":8080")
	APP_PORT := os.Getenv("APP_PORT")
	serveMux := http.NewServeMux()
	serveMux.Handle("/dogs", DogRequest{})
	http.ListenAndServe(APP_PORT, serveMux)

	os.Unsetenv("APP_PORT")
}

func GetDogImage() (*DogResult, error) {

	url := "https://dog.ceo/api/breeds/image/random"

	request, httpError := http.Get(url)

	if httpError != nil {
		return nil, errors.New("falha ao consultar api de dogs")
	}

	defer request.Body.Close()
	bytes, error := io.ReadAll(request.Body)

	if error != nil {
		return nil, errors.New("falha ao ler resposta")
	}

	dogResult := DogResult{}
	json.Unmarshal(bytes, &dogResult)
	return &dogResult, nil
}

func (dogRequest DogRequest) ServeHTTP(response http.ResponseWriter, request *http.Request) {

	dogResult, apiError := GetDogImage()

	if apiError != nil {
		fmt.Fprintf(os.Stderr, "Falha ao consultar dados")
	}

	response.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(response)
	encoder.Encode(dogResult)
}

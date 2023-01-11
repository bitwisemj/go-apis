package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Address struct {
	ZipCode      string `json:"cep"`
	Street       string `json:"logradouro"`
	Complement   string `json:"complemento"`
	Neighborhood string `json:"bairro"`
	Place        string `json:"localidade"`
	State        string `json:"uf"`
	Ibge         string `json:"ibge"`
	Gia          string `json:"gia"`
	Ddd          string `json:"ddd"`
	Siafi        string `json:"siafi"`
}

func main() {
	os.Setenv("APP_PORT", ":8082")
	port := os.Getenv("APP_PORT")

	http.HandleFunc("/address", GetAddressRequest)
	http.ListenAndServe(port, nil)

	os.Unsetenv("APP_PORT")
}

func GetAddressRequest(response http.ResponseWriter, request *http.Request) {

	if request.URL.Path != "/address" {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	zipCode := request.URL.Query().Get("zipCode")
	if zipCode == "" {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	address, error := GetAddress(zipCode)
	if error != nil {
		panic(error)
	}

	response.Header().Set("Content-Type", "application/json")
	response.Header().Set("X-Amzn-Key", "a1b2c3-d4e5c3")

	json.NewEncoder(response).Encode(address)
}

func GetAddress(zipCode string) (*Address, error) {

	url := fmt.Sprintf("https://viacep.com.br/ws/%v/json/", zipCode)

	req, err := http.Get(url)

	if err != nil {
		return nil, errors.New("falha ao consultar cep")
	}

	defer req.Body.Close()
	bytes, _ := io.ReadAll(req.Body)
	address := Address{}
	json.Unmarshal(bytes, &address)

	return &address, nil
}

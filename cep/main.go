package main

import (
	"encoding/json"
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

	fileName := "address.json"
	for _, zipCode := range os.Args[1:] {

		address := ToAddress(GetAddress(zipCode))
		jsonString := ToJSON(address)

		WriteFile(fileName, jsonString)
		fmt.Printf("%v\n", jsonString)
	}
}

func WriteFile(fileName, content string) {

	file, error := os.Create(fileName)

	if error != nil {
		fmt.Fprintf(os.Stderr, "Falha ao criar arquivo")
	}
	defer file.Close()
	file.WriteString(content)
}

func GetAddress(zipCode string) string {

	url := fmt.Sprintf("https://viacep.com.br/ws/%v/json/", zipCode)
	request, error := http.Get(url)

	if error != nil {
		fmt.Fprintf(os.Stderr, "Falha ao consultar endereço")
	}

	defer request.Body.Close()
	bytes, err := io.ReadAll(request.Body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Falha ao obter body da response")
	}

	content := string(bytes)
	return content
}

func ToAddress(jsonString string) Address {

	address := Address{}
	var ptrAddress *Address = &address
	json.Unmarshal([]byte(jsonString), ptrAddress)

	return address
}

func ToJSON(address Address) string {

	bytes, error := json.Marshal(address)

	if error != nil {
		fmt.Fprintf(os.Stderr, "Falha ao converter objeto para json")
	}

	content := string(bytes)
	return content
}

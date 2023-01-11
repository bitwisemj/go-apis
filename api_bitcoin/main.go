package main

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
)

type Asset struct {
	Id                string `json:"id"`
	Rank              string `json:"rank"`
	Symbol            string `json:"symbol"`
	Name              string `json:"name"`
	Supply            string `json:"supply"`
	MaxSupply         string `json:"maxSupply"`
	MarketCapUsd      string `json:"marketCapUsd"`
	VolumeUsd24Hr     string `json:"volumeUsd24Hr"`
	PriceUsd          string `json:"priceUsd"`
	ChangePercent24Hr string `json:"changePercent24Hr"`
	Vwap24Hr          string `json:"vWap24Hr"`
	Explorer          string `json:"explorer"`
}

type AssetData struct {
	Data []Asset `json:"data"`
}

type AssetHandler struct{}

func main() {

	os.Setenv("APP_PORT", ":8080")
	APP_PORT := os.Getenv("APP_PORT")
	serveMux := http.NewServeMux()
	serveMux.Handle("/assets", AssetHandler{})
	http.ListenAndServe(APP_PORT, serveMux)
}

func (assetHandler AssetHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {

	assetData, error := GetAssets()

	if error != nil {
		log.Fatal("Could not retrieve assets")
	}

	encoder := json.NewEncoder(response)
	encoder.Encode(assetData)
}

func GetAssets() (*AssetData, error) {

	url := "https://api.coincap.io/v2/assets"
	request, httpError := http.Get(url)

	if httpError != nil {
		return nil, errors.New("fail to retrieve bitcoin assets")
	}

	defer request.Body.Close()
	bytes, error := io.ReadAll(request.Body)

	if error != nil {
		return nil, errors.New("could not read request body")
	}

	assetData := AssetData{}
	json.Unmarshal(bytes, &assetData)

	return &assetData, nil
}

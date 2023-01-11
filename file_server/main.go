package main

import (
	"log"
	"net/http"
	"os"
)

func main() {

	log.Print("Starting application")

	os.Setenv("APP_PORT", ":8080")
	APP_PORT := os.Getenv("APP_PORT")

	log.Printf("Application listening on %v", APP_PORT)
	fileServer := http.FileServer(http.Dir("./public"))
	serveMux := http.NewServeMux()
	serveMux.Handle("/", fileServer)
	http.ListenAndServe(APP_PORT, serveMux)
}

package main

import (
	"JMeter_Go/internal/config"
	"JMeter_Go/internal/handlers"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/hello", handlers.HelloHandler).Methods("GET")
	r.HandleFunc("/upload-jmx", handlers.UploadFileHandler).Methods("POST")
	r.HandleFunc("/run-tests", handlers.TestRunnerHandler).Methods("GET")

	cfg := config.LoadConfig()
	port := cfg.Port

	fmt.Println("App running on Port", port)

	http.ListenAndServe(port, r)

}

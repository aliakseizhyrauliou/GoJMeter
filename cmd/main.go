package main

import (
	"JMeter_Go/internal/config"
	"JMeter_Go/internal/handlers"
	"fmt"
	"net/http"
)

func main() {

	cfg := config.LoadConfig()
	port := cfg.Port

	fmt.Println("App running on Port", port)

	http.HandleFunc("/hello", handlers.HelloHandler)
	http.HandleFunc("/upload-jmx", handlers.UploadFileHandler)
	http.HandleFunc("/run-tests", handlers.TestRunnerHandler)

	http.ListenAndServe(port, nil)
}

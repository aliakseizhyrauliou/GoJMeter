package main


import (
	"fmt"
	"net/http"
	"JMeter_Go/internal/handlers"
	"JMeter_Go/internal/config"
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
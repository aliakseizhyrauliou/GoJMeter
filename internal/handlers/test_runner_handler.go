package handlers

import (
	"fmt"
	"net/http"
	"JMeter_Go/internal/jmeter"
	"log"
)

func TestRunnerHandler(w http.ResponseWriter, r *http.Request) {
	userCount := r.URL.Query().Get("userCount")
	if userCount != "10" && userCount != "100" && userCount != "1000" {
		http.Error(w, "Invalid userCount value. Supported values are: 10, 100, 1000", http.StatusBadRequest)
		log.Println("UserCount:", userCount)
		return
	}


	err := jmeter.RunJMeterTest(userCount)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to run JMeter test for %s users", userCount), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "JMeter test for %s users completed successfully", userCount)
}
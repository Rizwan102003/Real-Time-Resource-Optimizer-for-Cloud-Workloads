package main

import (
	"cloud-optimizer/services"
	"log"
	"net/http"
)

func main() {
	instanceID := ""
	projectID := "real-time-resource-optimizer"
	services.FetchGCPMetrics(instanceID, projectID)
	http.HandleFunc("/metrics", controllers.FetchMetricsHandler)
	log.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

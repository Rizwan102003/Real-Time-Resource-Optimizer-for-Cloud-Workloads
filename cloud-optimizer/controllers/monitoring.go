package controllers

import (
	"cloud-optimizer/services"
	"encoding/json"
	"fmt"
	"net/http"
)

func FetchMetricsHandler(w http.ResponseWriter, r *http.Request) {
	projectID := r.URL.Query().Get("project_id")
	instanceID := r.URL.Query().Get("instance_id")

	if projectID == "" || instanceID == "" {
		http.Error(w, "project_id and instance_id are required", http.StatusBadRequest)
		return
	}

	// Fetch metrics
	metrics, err := services.FetchGCPMetrics(instanceID, projectID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching metrics: %v", err), http.StatusInternalServerError)
		return
	}

	// Return metrics as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(metrics)
}

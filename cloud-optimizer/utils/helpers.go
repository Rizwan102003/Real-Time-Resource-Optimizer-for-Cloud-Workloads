package utils

import (
	"fmt"
	"log"
)

// LogError logs error messages with a consistent format
func LogError(message string, err error) {
	log.Printf("ERROR: %s: %v", message, err)
}

// FormatMetric formats the metric data for readability
func FormatMetric(metricData string) string {
	// Example formatting logic
	return fmt.Sprintf("Formatted: %s", metricData)
}

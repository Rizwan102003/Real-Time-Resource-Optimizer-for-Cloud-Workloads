package main

import (
	"cloud-optimizer/services"
)

func main() {
	instanceID := "4112819316644587855"
	projectID := "real-time-resource-optimizer"
	services.FetchGCPMetrics(instanceID, projectID)
}

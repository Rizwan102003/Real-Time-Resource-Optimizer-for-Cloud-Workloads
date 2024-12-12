package main

import (
	"cloud-optimizer/services"
)

func main() {
	instanceID := ""
	projectID := "real-time-resource-optimizer"
	services.FetchGCPMetrics(instanceID, projectID)
}

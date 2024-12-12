package services

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/genproto/googleapis/monitoring/v3"
)

// FetchGCPMetrics fetches CPU utilization metrics from GCP
func FetchGCPMetrics(instanceID string, projectID string) ([]*monitoring.TimeSeries, error) {
	// Create a Cloud Monitoring client
	ctx := context.Background()
	client, err := monitoring.NewMetricClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("Error creating client: %v", err)
	}
	defer client.Close()

	// Define the query
	startTime := time.Now().Add(-1 * time.Hour).Format(time.RFC3339)
	endTime := time.Now().Format(time.RFC3339)
	filter := fmt.Sprintf(
		`metric.type="compute.googleapis.com/instance/cpu/utilization" AND resource.labels.instance_id="%s"`,
		instanceID,
	)

	req := &monitoring.ListTimeSeriesRequest{
		Name:   "projects/" + projectID,
		Filter: filter,
		Interval: &monitoring.TimeInterval{
			StartTime: &startTime,
			EndTime:   &endTime,
		},
		Aggregation: &monitoring.Aggregation{
			AlignmentPeriod:  "300s", // 5 minutes
			PerSeriesAligner: monitoring.Aggregation_ALIGN_MEAN,
		},
	}

	// Fetch data
	it := client.ListTimeSeries(ctx, req)
	var metrics []*monitoring.TimeSeries
	for {
		series, err := it.Next()
		if err != nil {
			if err.Error() == "iterator done" {
				break
			}
			return nil, fmt.Errorf("Error fetching metrics: %v", err)
		}
		metrics = append(metrics, series)
	}

	return metrics, nil
}

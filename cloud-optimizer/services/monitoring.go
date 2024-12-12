package services

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/genproto/googleapis/monitoring/v3"
)

func FetchGCPMetrics(instanceID string, projectID string) {
	// Create a Cloud Monitoring client
	ctx := context.Background()
	client, err := monitoring.NewMetricClient(ctx)
	if err != nil {
		fmt.Println("Error creating client:", err)
		return
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
	for {
		series, err := it.Next()
		if err != nil {
			if err.Error() == "iterator done" {
				break
			}
			fmt.Println("Error fetching metrics:", err)
			return
		}

		// Print time series data
		fmt.Println("Metric:", series.GetMetric().GetType())
		for _, point := range series.GetPoints() {
			fmt.Printf("Value at %s: %v\n", point.GetInterval().GetEndTime(), point.GetValue().GetDoubleValue())
		}
	}
}

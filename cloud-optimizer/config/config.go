package config

import (
	"log"
	"os"
)

var (
	ProjectID  = os.Getenv("PROJECT_ID")
	InstanceID = os.Getenv("INSTANCE_ID")
)

func LoadConfig() {
	if ProjectID == "" || InstanceID == "" {
		log.Fatal("GCP_PROJECT_ID and GCP_INSTANCE_ID must be set in environment variables")
	}
}

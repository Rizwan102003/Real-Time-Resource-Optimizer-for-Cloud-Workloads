# Real-Time-Resource-Optimizer-for-Cloud-Workloads
A tool that dynamically analyzes and optimizes cloud resource utilization for cost efficiency and performance.

## Features

- Real-time monitoring of cloud workloads on GCP.
- Fetch CPU utilization metrics using **Google Cloud Monitoring**.
- Database integration to store and analyze metrics.
- Optimization suggestions based on resource usage.
- REST API for querying metrics and recommendations.

## Tech Stack:
# Backend: ![Go](https://img.shields.io/badge/Go-00ADD8?style=flat&logo=go&logoColor=white)
# Cloud SDKs: ![GCP](https://img.shields.io/badge/GCP-4285F4?style=flat&logo=google-cloud&logoColor=white)  ![Google Cloud Monitoring](https://img.shields.io/badge/Google_Cloud_Monitoring-FF9900?style=flat&logo=google-cloud&logoColor=white)
# Database: ![PostgreSQL](https://img.shields.io/badge/PostgreSQL-336791?style=flat&logo=postgresql&logoColor=white)
![Gin](https://img.shields.io/badge/Gin-67C7B5?style=flat&logo=gin&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-2496ED?style=flat&logo=docker&logoColor=white)


## Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/yourusername/cloud-optimizer.git
    cd cloud-optimizer
    ```

2. Install dependencies:
    ```bash
    go mod tidy
    ```

3. Set up your **GCP credentials** and enable required APIs.

4. Run the application:
    ```bash
    go run main.go
    ```

## License

MIT License. See `LICENSE` for more information.

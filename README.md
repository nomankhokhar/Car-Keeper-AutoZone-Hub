Here is the professional version of your README. I have organized it into logical sections, formatted the code blocks for readability, and corrected minor typographical errors while retaining 100% of your original information.

-----

# Car-Keeper-AutoZone-Hub

This repository contains the backend service for the Car-Keeper AutoZone Hub. Below are the instructions for building, running, and deploying the application using Docker Compose, local development environments, manual Docker containers, and Kubernetes.

## 1\. Quick Start (Docker Compose)

The easiest way to build and run the application along with all dependencies (PostgreSQL, Jaeger, Prometheus, Grafana) is via Docker Compose.

### Build and Run

Go to Car-Keeper-AutoZone-Hub/docker

```bash
docker-compose up --build -d
```

### Accessing Internal Components

| Component | Action / Command | URL / Details |
| :--- | :--- | :--- |
| **PostgreSQL** | Access Database | `docker exec -it my-postgres psql -U caruser -d car` |
| **Jaeger UI** | View Traces | [http://localhost:16686](https://www.google.com/search?q=http://localhost:16686) |
| **Prometheus UI** | View Metrics | [http://localhost:9090](https://www.google.com/search?q=http://localhost:9090) |
| **Grafana UI** | View Dashboards | [http://localhost:3000](https://www.google.com/search?q=http://localhost:3000) (User/Pass: `admin`) |

-----

## 2\. Local Development (Running from Source)

If you wish to run the Go application locally on your machine (without Dockerizing the app itself), follow these steps.

### Step 1: Start the Database

Run a standalone PostgreSQL container:

```bash
docker run --name my-postgres \
  -e POSTGRES_USER=caruser \
  -e POSTGRES_PASSWORD=carpassword \
  -e POSTGRES_DB=car \
  -p 5432:5432 \
  -d postgres:latest
```

### Step 2: Configure Environment Variables

Set the necessary environment variables for your local session:

```bash
export DB_HOST=localhost
export DB_PORT=5432
export DB_NAME=car
export DB_USER=caruser
export DB_PASSWORD=carpassword
```

### Step 3: Run the Application

Execute the Go entry point:

```bash
go run main.go
```

-----

Check the application logs to ensure it connects successfully to the database.
http://localhost:8080/health should return a healthy status.

## 3\. Manual Docker Deployment

To pull the latest images from Docker Hub and run the application manually using a custom Docker network.

### Step 1: Pull Images

```bash
docker pull nomanali1114/car_keeper_backend:latest
docker pull postgres:latest
```

### Step 2: Create Network

Create a dedicated bridge network for the containers to communicate:

```bash
docker network create car-network
```

### Step 3: Run PostgreSQL

Start the database attached to the network:

```bash
docker run --name my-postgres --network car-network \
  -e POSTGRES_USER=caruser \
  -e POSTGRES_PASSWORD=carpassword \
  -e POSTGRES_DB=car \
  -p 5432:5432 \
  -d postgres:latest
```

*Note: Wait for a few seconds to ensure PostgreSQL is fully up and running before starting the application.*

### Step 4: Run Car-Keeper-AutoZone-Hub

Start the application container attached to the network:

```bash
docker run --name golang-app -p 8080:8080 --network car-network \
  -e DB_HOST=my-postgres \
  -e DB_PORT=5432 \
  -e DB_NAME=car \
  -e DB_USER=caruser \
  -e DB_PASSWORD=carpassword \
  nomanali1114/car_keeper_backend:latest
```

-----

## 4\. Kubernetes Deployment

Below are the access points for the application and observability tools when deployed to a Kubernetes cluster using Traefik as the Load Balancer.

**Note:** Replace `{EXTERNAL-IP}` with the actual LoadBalancer IP assigned by your cloud provider.

| Component | Access URL |
| :--- | :--- |
| **Car-Keeper API** | `http://{EXTERNAL-IP}:8080/health` |
| **Prometheus UI** | `http://{EXTERNAL-IP}:9090/metrics` |
| **Grafana UI** | `http://{EXTERNAL-IP}:3000` |
| **Jaeger UI** | `http://{EXTERNAL-IP}:16686` |
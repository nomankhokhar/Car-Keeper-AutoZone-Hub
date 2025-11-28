![Project Screenshot](./image/demo_work.png)

# Full Monitoring System for Car-Keeper-AutoZone-Hub

**Monitoring Architecture**
The application is monitored using a modern observability pipeline designed to provide real-time visibility into system behavior:

* **Jaeger (Tracing):** Tracks requests as they flow through the system, allowing for quick identification of bottlenecks and failures.
* **Prometheus (Metrics):** Scrapes and stores time-series data, such as API request counts, error rates, and resource usage.
* **Grafana (Visualization):** Connects to Prometheus to display actionable data on customizable dashboards, making it easy to monitor system stability at a glance.

This repository contains the backend service for the Car-Keeper AutoZone Hub. Below are the instructions for building, running, and deploying the application using Docker Compose, local development environments, manual Docker containers, and Kubernetes.

## Postman Collection
postman_collection.json contains the Postman collection for testing the API endpoints.

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

---

## 4. Kubernetes Deployment

Navigate to the `kubernetes-deployment` directory to locate the manifest files required to deploy the application stack (App, Database, Monitoring) using Traefik as the Ingress Controller.

### Directory Structure

* **`traefik-deployment/`**
    Contains the core Traefik Load Balancer Deployment and Service definitions.
    * *Refer to the internal README.md within this folder for specific Traefik installation instructions.*
* **`Apps-deployment/`**
    Contains the standard Kubernetes manifests (Deployments and Services) for the application components: PostgreSQL, Car-Keeper-AutoZone-Hub, Jaeger, Prometheus, and Grafana.
* **`App-routes/`**
    Contains the Traefik `IngressRoute` custom resources that define how external traffic is routed to specific internal services, including port configurations.

### Deployment & Verification

Once you have applied the configurations using `kubectl apply -f`, the services will be exposed via the External IP assigned to the Traefik Load Balancer.

**Access Points**

Replace `{EXTERNAL-IP}` with the actual public IP address of your Load Balancer (e.g., `129.212.x.x`).

| Component | Access URL |
| :--- | :--- |
| **Car-Keeper API** | `http://{EXTERNAL-IP}:8080/health` |
| **Prometheus UI** | `http://{EXTERNAL-IP}:9090/metrics` |
| **Grafana UI** | `http://{EXTERNAL-IP}:3000` |
| **Jaeger UI** | `http://{EXTERNAL-IP}:16686` |
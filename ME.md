# Car-Keeper-AutoZone-Hub

# To Build and Run the Application
docker-compose up --build -d

# To Access the PostgreSQL Database
docker exec -it my-postgres psql -U caruser -d car

# To Access the Jaeger UI
Open your web browser and navigate to http://localhost:16686

# To Access the Prometheus UI
Open your web browser and navigate to http://localhost:9090

# To Access the Grafana UI
Open your web browser and navigate to http://localhost:3000 (default username and password are both "admin")


# To run the application without docker-compose

# To run the postgres container directly
docker run --name my-postgres   -e POSTGRES_USER=caruser   -e POSTGRES_PASSWORD=carpassword   -e POSTGRES_DB=car   -p 5432:5432   -d postgres:latest


# Set the Environment Variables for Local Development
export DB_HOST=localhost
export DB_PORT=5432
export DB_NAME=car
export DB_USER=caruser
export DB_PASSWORD=carpassword

# Run the Application 
go run main.go



# To Run the Container of Car-Keeper-AutoZone-Hub Directly from Docker Hub Pulls

To pull the latest image from Docker Hub and run it along with a PostgreSQL container, follow these steps:
docker pull nomanali1114/car_keeper_backend:latest
docker pull postgres:latest

## First Create a Docker Network
docker network create car-network

## Then Run the PostgreSQL Container
docker run --name my-postgres --network car-network -e POSTGRES_USER=caruser -e POSTGRES_PASSWORD=carpassword -e POSTGRES_DB=car -p 5432:5432 -d postgres:latest

Wait form a few seconds to ensure PostgreSQL is up and running.

## Finally, Run the Car-Keeper-AutoZone-Hub Container
docker run --name golang-app -p 8080:8080 --network car-network -e DB_HOST=my-postgres -e DB_PORT=5432 -e DB_NAME=car -e DB_USER=caruser -e DB_PASSWORD=carpassword nomanali1114/car_keeper_backend:latest


# To Deploy the Application to Kubernetes Cluster Using kubectl with Load Balancer Traefik

# To Access the Prometheus UI in the Cluster
http://{EXTERNAL-IP}:9090/metrics

# To Access the Grafana UI in the Cluster
http://{EXTERNAL-IP}:3000

# To Access the Jaeger UI in the Cluster
http://{EXTERNAL-IP}:16686

# To access the Car-Keeper-AutoZone-Hub API in the Cluster
http://{EXTERNAL-IP}:8080/health




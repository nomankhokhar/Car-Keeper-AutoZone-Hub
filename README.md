# Car-Keeper-AutoZone-Hub

# To Build and Run the Application
docker-compose up --build -d

# To Access the PostgreSQL Database
docker exec -it my-postgres psql -U caruser -d car

# To run the postgres container directly
docker run --name my-postgres   -e POSTGRES_USER=caruser   -e POSTGRES_PASSWORD=carpassword   -e POSTGRES_DB=car   -p 5432:5432   -d postgres:latest

# Environment Variables for Local Development
export DB_HOST=localhost
export DB_PORT=5432
export DB_NAME=car
export DB_USER=caruser
export DB_PASSWORD=carpassword

# To Run the Application Locally
go run main.go

# To Access the Jaeger UI
Open your web browser and navigate to http://localhost:16686

# To Access the Prometheus UI
Open your web browser and navigate to http://localhost:9090

# To Access the Grafana UI
Open your web browser and navigate to http://localhost:3000 (default username and password are both "admin")
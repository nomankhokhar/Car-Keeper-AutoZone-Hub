mkdir -p Car_Keeper_backend/cmd/api \
  Car_Keeper_backend/internal/{config,database/migrations,models,repository,service,handler,middleware,dto,validator} \
  Car_Keeper_backend/pkg/{logger,response,utils} \
  Car_Keeper_backend/scripts \
  Car_Keeper_backend/docker

touch Car_Keeper_backend/cmd/api/main.go \
  Car_Keeper_backend/internal/config/config.go \
  Car_Keeper_backend/internal/database/postgres.go \
  Car_Keeper_backend/internal/models/{user.go,product.go} \
  Car_Keeper_backend/internal/repository/{user_repository.go,product_repository.go} \
  Car_Keeper_backend/internal/service/{user_service.go,product_service.go} \
  Car_Keeper_backend/internal/handler/{user_handler.go,product_handler.go} \
  Car_Keeper_backend/internal/middleware/{auth.go,cors.go,logger.go} \
  Car_Keeper_backend/internal/dto/{user_dto.go,product_dto.go} \
  Car_Keeper_backend/internal/validator/validator.go \
  Car_Keeper_backend/pkg/logger/logger.go \
  Car_Keeper_backend/pkg/response/response.go \
  Car_Keeper_backend/pkg/utils/helpers.go \
  Car_Keeper_backend/scripts/{migrate.sh,seed.sh} \
  Car_Keeper_backend/docker/{Dockerfile,docker-compose.yml} \
  Car_Keeper_backend/{.env.example,.gitignore,Makefile,README.md,go.mod}


# Car-Keeper-AutoZone-Hub

# To Build and Run the Application
docker-compose up --build -d

# To Access the PostgreSQL Database
docker exec -it my-postgres psql -U caruser -d car

# To run the postgres container directly
docker run --name my-postgres   -e POSTGRES_USER=caruser   -e POSTGRES_PASSWORD=carpassword   -e POSTGRES_DB=car   -p 5432:5432   -d postgres:alpine

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
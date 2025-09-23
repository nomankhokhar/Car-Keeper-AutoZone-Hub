package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/nomankhokhar/Car-Keeper-AutoZone-Hub/driver"
	carHandler "github.com/nomankhokhar/Car-Keeper-AutoZone-Hub/handler/car"
	engineHandler "github.com/nomankhokhar/Car-Keeper-AutoZone-Hub/handler/engine"
	"github.com/nomankhokhar/Car-Keeper-AutoZone-Hub/middleware"
	carService "github.com/nomankhokhar/Car-Keeper-AutoZone-Hub/service/car"
	engineService "github.com/nomankhokhar/Car-Keeper-AutoZone-Hub/service/engine"
	carStore "github.com/nomankhokhar/Car-Keeper-AutoZone-Hub/store/car"
	engineStore "github.com/nomankhokhar/Car-Keeper-AutoZone-Hub/store/engine"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	driver.InitDB()
	defer driver.CloseDB()

	db := driver.GetDB()
	carStore := carStore.New(db)
	carService := carService.NewCarService(carStore)

	engineStore := engineStore.New(db)
	engineService := engineService.NewEngineService(engineStore)

	carHandler := carHandler.NewCarHandler(carService)
	engineHandler := engineHandler.NewEngineHandler(engineService)

	router := mux.NewRouter()

	// Middleware
	router.Use(middleware.AuthMiddleware)

	schemaFile := "store/schema.sql"
	if err := executeSchema(db, schemaFile); err != nil {
		log.Fatalf("Failed to execute schema: %v", err)
	}

	// Car routes
	router.HandleFunc("/cars/{id}", carHandler.GetCarByID).Methods("GET")
	router.HandleFunc("/cars", carHandler.GetCarsByBrand).Methods("GET")
	router.HandleFunc("/cars", carHandler.CreateCar).Methods("POST")
	router.HandleFunc("/cars/{id}", carHandler.UpdateCar).Methods("PUT")
	router.HandleFunc("/cars/{id}", carHandler.DeleteCar).Methods("DELETE")

	// Engine routes
	router.HandleFunc("/engines/{id}", engineHandler.GetEngineByID).Methods("GET")
	router.HandleFunc("/engines", engineHandler.CreateEngine).Methods("POST")
	router.HandleFunc("/engines/{id}", engineHandler.UpdateEngine).Methods("PUT")
	router.HandleFunc("/engines/{id}", engineHandler.DeleteEngine).Methods("DELETE")

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Starting server on :" + port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("Server failed: %v", err)
	}

}

func executeSchema(db *sql.DB, schemaFile string) error {
	executeSchema, err := os.ReadFile(schemaFile)
	if err != nil {
		return fmt.Errorf("failed to read schema file: %w", err)
	}

	_, err = db.Exec(string(executeSchema))
	if err != nil {
		return fmt.Errorf("failed to execute schema: %w", err)
	}

	return nil
}

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
	LoginHandler "github.com/nomankhokhar/Car-Keeper-AutoZone-Hub/handler/login"
	middleware "github.com/nomankhokhar/Car-Keeper-AutoZone-Hub/middleware"
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

	schemaFile := "store/schema.sql"
	if err := executeSchema(db, schemaFile); err != nil {
		log.Fatalf("Failed to execute schema: %v", err)
	}

	router.HandleFunc("/login", LoginHandler.LoginHandler).Methods("POST")

	// Middleware
	protected := router.PathPrefix("/").Subrouter()
	protected.Use(middleware.AuthMiddleware)
	// Car routes
	protected.HandleFunc("/cars/{id}", carHandler.GetCarByID).Methods("GET")
	protected.HandleFunc("/cars", carHandler.GetCarsByBrand).Methods("GET")
	protected.HandleFunc("/cars", carHandler.CreateCar).Methods("POST")
	protected.HandleFunc("/cars/{id}", carHandler.UpdateCar).Methods("PUT")
	protected.HandleFunc("/cars/{id}", carHandler.DeleteCar).Methods("DELETE")

	// Engine routes
	protected.HandleFunc("/engines/{id}", engineHandler.GetEngineByID).Methods("GET")
	protected.HandleFunc("/engines", engineHandler.CreateEngine).Methods("POST")
	protected.HandleFunc("/engines/{id}", engineHandler.UpdateEngine).Methods("PUT")
	protected.HandleFunc("/engines/{id}", engineHandler.DeleteEngine).Methods("DELETE")

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

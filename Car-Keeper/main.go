package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/nomankhokhar/Car-Keeper-AutoZone-Hub/driver"
	carHandler "github.com/nomankhokhar/Car-Keeper-AutoZone-Hub/handler/car"
	engineHandler "github.com/nomankhokhar/Car-Keeper-AutoZone-Hub/handler/engine"
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

	router.HandleFunc("/cars/{id}", carHandler.GetCarByID).Methods("GET")
	router.HandleFunc("/cars", carHandler.GetCarsByBrand).Methods("GET")
	router.HandleFunc("/cars", carHandler.CreateCar).Methods("POST")
	router.HandleFunc("/cars/{id}", carHandler.UpdateCar).Methods("PUT")
	router.HandleFunc("/cars/{id}", carHandler.DeleteCar).Methods("DELETE")

	router.HandleFunc("/engines/{id}", engineHandler.GetEngineByID).Methods("GET")
	router.HandleFunc("/engines", engineHandler.CreateEngine).Methods("POST")
	router.HandleFunc("/engines/{id}", engineHandler.UpdateEngine).Methods("PUT")
	router.HandleFunc("/engines/{id}", engineHandler.DeleteEngine).Methods("DELETE")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.ListenAndServe(":"+port, router)
	fmt.Println("Starting server on :" + port)
}

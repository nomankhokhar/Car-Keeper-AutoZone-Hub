package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/nomankhokhar/Car-Keeper-AutoZone-Hub/driver"
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

}

package handler

import (
	"Car_Keeper/internal/models"
	"Car_Keeper/internal/service"

	"github.com/gin-gonic/gin"
)

type CarHandler struct {
	service service.CarService
}

func NewCarHandler(service service.CarService) *CarHandler {
	return &CarHandler{service: service}
}

// Get by Car ID
func (h *CarHandler) GetCarByID(c *gin.Context) {
	carID := c.Param("carid")
	car, err := h.service.GetCarByID(carID)
	if err != nil {
		c.JSON(404, gin.H{"message": "Car not found", "error": err.Error()})
		return
	}
	c.JSON(200, car)
}

func (h *CarHandler) GetCarByBrand(c *gin.Context) {
	brand := c.Query("brand")
	cars, err := h.service.GetCarByBrand(brand)
	if err != nil {
		c.JSON(404, gin.H{"message": "Cars not found", "error": err.Error()})
		return
	}
	c.JSON(200, cars)
}

func (h *CarHandler) CreateCar(c *gin.Context) {
	var carReq models.CarRequest
	// Bind JSON request to struct
	if err := c.ShouldBindJSON(&carReq); err != nil {
		c.JSON(400, gin.H{"message": "Invalid request", "error": err.Error()})
		return
	}
	// Call service to create car
	if err := h.service.CreateCar(&carReq); err != nil {
		c.JSON(500, gin.H{"message": "Failed to create car", "error": err.Error()})
		return
	}
	c.JSON(201, gin.H{"message": "Car created successfully"})
}

func (h *CarHandler) UpdateCar(c *gin.Context) {
	carId := c.Param("carid")
	var carReq models.CarRequest
	// Bind JSON request to struct
	if err := c.ShouldBindJSON(&carReq); err != nil {
		c.JSON(400, gin.H{"message": "Invalid request", "error": err.Error()})
		return
	}
	// Call service to update car
	if err := h.service.UpdateCar(carId, &carReq); err != nil {
		c.JSON(500, gin.H{"message": "Failed to update car", "error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Car updated successfully"})
}

func (h *CarHandler) DeleteCar(c *gin.Context) {
	carID := c.Param("carid")
	if err := h.service.DeleteCar(carID); err != nil {
		c.JSON(500, gin.H{"message": "Failed to delete car", "error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Car deleted successfully"})
}

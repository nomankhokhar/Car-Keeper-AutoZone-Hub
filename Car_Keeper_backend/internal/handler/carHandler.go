package handler

import (
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

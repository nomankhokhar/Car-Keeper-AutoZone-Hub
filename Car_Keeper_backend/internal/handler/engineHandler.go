package handler

import (
	"Car_Keeper/internal/models"
	"Car_Keeper/internal/service"

	"github.com/gin-gonic/gin"
)

type EngineHandler struct {
	service service.EngineService
}

func NewEngineHandler(service service.EngineService) *EngineHandler {
	return &EngineHandler{service: service}
}

func (h *EngineHandler) GetEngineByID(c *gin.Context) {

	engineID := c.Param("engineid")
	if engineID == "" {
		c.JSON(400, gin.H{"message": "Engine ID is required"})
		return
	}

	// Call service to get engine by ID
	engine, err := h.service.GetEngineByID(engineID)
	if err != nil {
		c.JSON(404, gin.H{"message": "Engine not found", "error": err.Error()})
		return
	}
	c.JSON(200, engine)
}

func (h *EngineHandler) CreateEngine(c *gin.Context) {
	var engineReq models.EngineRequest
	if err := c.ShouldBindJSON(&engineReq); err != nil {
		c.JSON(400, gin.H{"message": "Invalid request", "error": err.Error()})
		return
	}

	// Call service to create engine
	engine, err := h.service.CreateEngine(&engineReq)
	if err != nil {
		c.JSON(500, gin.H{"message": "Failed to create engine", "error": err.Error()})
		return
	}
	c.JSON(201, engine)
}

func (h *EngineHandler) UpdateEngine(c *gin.Context) {
	engineID := c.Param("engineid")
	if engineID == "" {
		c.JSON(400, gin.H{"message": "Engine ID is required"})
		return
	}

	var engineReq models.EngineRequest
	if err := c.ShouldBindJSON(&engineReq); err != nil {
		c.JSON(400, gin.H{"message": "Invalid request", "error": err.Error()})
		return
	}

	// Call service to update engine
	engine, err := h.service.UpdateEngine(engineID, &engineReq)
	if err != nil {
		c.JSON(500, gin.H{"message": "Failed to update engine", "error": err.Error()})
		return
	}
	c.JSON(200, engine)
}

func (h *EngineHandler) DeleteEngine(c *gin.Context) {
	engineID := c.Param("engineid")
	if engineID == "" {
		c.JSON(400, gin.H{"message": "Engine ID is required"})
		return
	}

	// Call service to delete engine
	if err := h.service.DeleteEngine(engineID); err != nil {
		c.JSON(500, gin.H{"message": "Failed to delete engine", "error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Engine deleted successfully"})
}

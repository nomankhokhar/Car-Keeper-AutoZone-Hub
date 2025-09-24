package engine

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/nomankhokhar/Car-Keeper-AutoZone-Hub/models"
	"github.com/nomankhokhar/Car-Keeper-AutoZone-Hub/service"
	"go.opentelemetry.io/otel"
)

type EngineHandler struct {
	service service.EngineServiceInterface
}

func NewEngineHandler(service service.EngineServiceInterface) *EngineHandler {
	return &EngineHandler{
		service: service,
	}
}

func (e *EngineHandler) GetEngineByID(w http.ResponseWriter, r *http.Request) {
	tracer := otel.Tracer("EngineHandler")
	ctx, span := tracer.Start(r.Context(), "GetEngineByID-Handler")
	defer span.End()

	vars := mux.Vars(r)
	id := vars["id"]

	resp, err := e.service.GetEngineByID(ctx, id)
	if err != nil {
		log.Println("Error: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Println("Fetched Engine: ", resp)
	respBody, err := json.Marshal(resp)
	if err != nil {
		log.Println("Error: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(respBody)
	if err != nil {
		log.Println("Error writing response: ", err)
	}
}

func (e *EngineHandler) CreateEngine(w http.ResponseWriter, r *http.Request) {
	tracer := otel.Tracer("EngineHandler")
	ctx, span := tracer.Start(r.Context(), "CreateEngine-Handler")
	defer span.End()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading request body: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var engineReq models.EngineRequest
	err = json.Unmarshal(body, &engineReq)

	if err != nil {
		log.Println("Error unmarshalling request body: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	createdEngine, err := e.service.CreateEngine(ctx, &engineReq)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error creating engine: ", err)
		return
	}
	fmt.Println("Created Engine: ", createdEngine)
	responseBody, err := json.Marshal(createdEngine)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error marshalling response body: ", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	_, err = w.Write(responseBody)
	if err != nil {
		log.Println("Error writing response: ", err)
	}
}

func (e *EngineHandler) UpdateEngine(w http.ResponseWriter, r *http.Request) {
	tracer := otel.Tracer("EngineHandler")
	ctx, span := tracer.Start(r.Context(), "UpdateEngine-Handler")
	defer span.End()

	vars := mux.Vars(r)
	id := vars["id"]

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading request body:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var engineReq models.EngineRequest
	err = json.Unmarshal(body, &engineReq)
	if err != nil {
		log.Println("Error unmarshalling request body:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	updatedEngine, err := e.service.UpdateEngine(ctx, id, &engineReq)
	if err != nil {
		log.Println("Error updating engine:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseBody, err := json.Marshal(updatedEngine)
	if err != nil {
		log.Println("Error marshalling response body:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if _, err := w.Write(responseBody); err != nil {
		log.Println("Error writing response:", err)
	}
}

func (e *EngineHandler) DeleteEngine(w http.ResponseWriter, r *http.Request) {
	tracer := otel.Tracer("EngineHandler")
	ctx, span := tracer.Start(r.Context(), "DeleteEngine-Handler")
	defer span.End()

	vars := mux.Vars(r)
	id := vars["id"]

	deletedEngine, err := e.service.DeleteEngine(ctx, id)
	if err != nil {
		log.Println("Error deleting engine: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		response := map[string]string{"error": "Failed to delete engine or invalid ID"}
		responseBody, _ := json.Marshal(response)
		_, _ = w.Write(responseBody)
		return
	}

	if deletedEngine.EngineID == uuid.Nil {
		w.WriteHeader(http.StatusNotFound)
		response := map[string]string{"error": "Engine not found"}
		responseBody, _ := json.Marshal(response)
		_, _ = w.Write(responseBody)
		return
	}

	jsonResponse, err := json.Marshal(deletedEngine)
	if err != nil {
		log.Println("Error marshalling response body: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		response := map[string]string{"error": "Failed to process response"}
		jsonResponse, _ := json.Marshal(response)
		_, _ = w.Write(jsonResponse)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(jsonResponse)
	if err != nil {
		log.Println("Error writing response: ", err)
	}
}

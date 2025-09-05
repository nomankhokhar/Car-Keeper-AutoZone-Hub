package engine

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nomankhokhar/Car-Keeper-AutoZone-Hub/models"
	"github.com/nomankhokhar/Car-Keeper-AutoZone-Hub/service"
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
	ctx := r.Context()
	vars := mux.Vars(r)
	id := vars["id"]

	resp, err := e.service.GetEngineByID(ctx, id)
	if err != nil {
		log.Println("Error: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

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
	ctx := r.Context()

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
	ctx := r.Context()
	vars := mux.Vars(r)
	id := vars["id"]

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading request body: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body, err = json.Marshal(body)

	if err != nil {
		log.Println("Error marshalling request body: ", err)
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

	updatedEngine, err := e.service.UpdateEngine(ctx, id, &engineReq)
	if err != nil {
		log.Println("Error updating engine: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseBody, err := json.Marshal(updatedEngine)
	if err != nil {
		log.Println("Error marshalling response body: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(responseBody)
	if err != nil {
		log.Println("Error writing response: ", err)
	}
}

func (e *EngineHandler) DeleteEngine(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	id := vars["id"]

	deletedEngine, err := e.service.DeleteEngine(ctx, id)
	if err != nil {
		log.Println("Error deleting engine: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	responseBody, err := json.Marshal(deletedEngine)
	if err != nil {
		log.Println("Error marshalling response body: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = w.Write(responseBody)
	if err != nil {
		log.Println("Error writing response: ", err)
	}
}

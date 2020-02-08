package handlers

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/hobord/golang-poc-rest/delivery/http/dto"
	"github.com/hobord/golang-poc-rest/domain/entity"
	"github.com/hobord/golang-poc-rest/usecase"
	"net/http"
)

// FooRestHTTPModule is handle the entity related http request responses
type FooRestHTTPModule struct {
	entityInteractor usecase.FooInteractorInterface
}

// CreateFooRestHTTPModule create a new http handler app to entity
func CreateFooRestHTTPModule(entityInteractor usecase.FooInteractorInterface) *FooRestHTTPModule {
	return &FooRestHTTPModule{
		entityInteractor: entityInteractor,
	}
}

// GetByID return entity by id
func (app *FooRestHTTPModule) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		http.Error(w, "missing id", http.StatusBadRequest)
		return
	}

	entity, err := app.entityInteractor.GetByID(r.Context(), string(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if entity == nil {
		err = errors.New("no resource found")
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	entityDTO := &dto.FooResponse{
		ID:    entity.ID,
		Title: entity.Title,
	}
	js, err := json.Marshal(entityDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// GetAll return all entity
func (app *FooRestHTTPModule) GetAll(w http.ResponseWriter, r *http.Request) {
	res, err := app.entityInteractor.GetAll(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if res == nil || len(res) == 0 {
		err = errors.New("no resource found")
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	entityDTOs := make([]dto.FooResponse, 0)
	for _, entity := range res {
		entityDTO := &dto.FooResponse{
			ID:    entity.ID,
			Title: entity.Title,
		}
		entityDTOs = append(entityDTOs, *entityDTO)
	}

	// Convert to json
	js, err := json.Marshal(entityDTOs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send back to response.
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// Create is update to persistent the entity
func (app *FooRestHTTPModule) Create(w http.ResponseWriter, r *http.Request) {
	// Decode the request DTO.
	decoder := json.NewDecoder(r.Body)
	var createDTO dto.FooCreateRequest
	err := decoder.Decode(&createDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create a new entity.
	entity, err := entity.CreateFooEntity(createDTO.Title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Save the new entity.
	err = app.entityInteractor.Save(r.Context(), entity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create a new response DTO.
	entityDTO := &dto.FooResponse{
		ID:    entity.ID,
		Title: entity.Title,
	}
	// Convert to json
	js, err := json.Marshal(entityDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send back to response.
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// Update is update to persistent the entity.
func (app *FooRestHTTPModule) Update(w http.ResponseWriter, r *http.Request) {
	// Decode the request DTO.
	decoder := json.NewDecoder(r.Body)
	var updateDTO dto.FooUpdateRequest
	err := decoder.Decode(&updateDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if updateDTO.ID == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	// Load the original entity.
	entity, err := app.entityInteractor.GetByID(r.Context(), updateDTO.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if entity == nil {
		err = errors.New("no resource found")
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Update the entity properties.
	entity.Title = updateDTO.Title

	// save the entity
	err = app.entityInteractor.Save(r.Context(), entity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create a response DTO.
	entityDTO := &dto.FooResponse{
		ID:    entity.ID,
		Title: entity.Title,
	}
	// Convert to json
	js, err := json.Marshal(entityDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send back to response.
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// Delete entity
func (app *FooRestHTTPModule) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		http.Error(w, "missing id", http.StatusBadRequest)
		return
	}

	entity, err := app.entityInteractor.GetByID(r.Context(), string(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if entity == nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = app.entityInteractor.Delete(r.Context(), entity.ID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

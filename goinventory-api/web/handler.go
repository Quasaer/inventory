package web

import (
	"encoding/json"
	"net/http"

	"github.com/Quasaer/goinventory-api/goinventory"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
)

func NewHandler(store goinventory.Store) *Handler {
	h := &Handler{
		Mux:   chi.NewMux(),
		store: store,
	}

	h.Use(middleware.Logger)

	h.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ping"))
	})

	h.Route("/inventoryList", func(r chi.Router) {
		r.Get("/{id}", h.GetInventoryList())
		r.Post("/", h.CreateInventoryList())
		r.Put("/{id}", h.UpdateInventoryList())
	})

	return h
}

type Handler struct {
	*chi.Mux
	store goinventory.Store
}

func (h *Handler) GetInventoryList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		idStr := chi.URLParam(r, "id")

		id, err := uuid.Parse(idStr)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		i, err := h.store.GetInventoryListByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusFound)
		json.NewEncoder(w).Encode(i)

	}
}

func (h *Handler) CreateInventoryList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		inventoryList := &goinventory.InventoryList{}
		if err := json.NewDecoder(r.Body).Decode(&inventoryList); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := h.store.ValidateInventoryListOnCreate(inventoryList); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := h.store.CreateInventoryList(inventoryList); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var returnInventoryListCreatedResp = &returnCreatedResponse{
			CreatedAt: inventoryList.CreatedAt,
			ID:        inventoryList.ID,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(returnInventoryListCreatedResp)

	}
}

func (h *Handler) UpdateInventoryList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")

		id, err := uuid.Parse(idStr)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		inventoryList := &goinventory.InventoryList{
			ID: id,
		}
		if err := json.NewDecoder(r.Body).Decode(&inventoryList); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := h.store.ValidateInventoryListOnUpdate(inventoryList); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := h.store.UpdateInventoryList(inventoryList); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var returnInventoryListUpdatedResp = &returnUpdateResponse{
			UpdatedAt: inventoryList.UpdatedAt,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(returnInventoryListUpdatedResp)

	}
}

type returnUpdateResponse struct {
	UpdatedAt int64
}

type returnCreatedResponse struct {
	CreatedAt int64
	ID        uuid.UUID
}

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
	h.Route("/inventoryList", func(r chi.Router) {
		r.Get("/{id}", h.GetInventoryList())
		r.Post("/", h.CreateInventoryList())
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

		if err := h.store.ValidateInventoryList(inventoryList); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := h.store.CreateInventoryList(inventoryList); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(inventoryList)

	}
}

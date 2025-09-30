package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/eif-courses/restaurant/internal/services"
)

type FoodHandler struct {
	foodService *services.FoodService
}

func NewFoodHandler(foodService *services.FoodService) *FoodHandler {
	return &FoodHandler{foodService: foodService}
}

func (f *FoodHandler) GetFoodList(w http.ResponseWriter, r *http.Request) {

	foods, err := f.foodService.GetFoodList(r.Context())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(foods)

}

func (f *FoodHandler) InsertFruit(w http.ResponseWriter, r *http.Request) {

	type InsertFruitRequest struct {
		Name  string  `json:"name"`
		Price float64 `json:"price"`
	}

	var req InsertFruitRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if req.Name == "" {
		http.Error(w, "name is required", http.StatusBadRequest)
		return
	}
	if req.Price == 0 {
		http.Error(w, "price is required", http.StatusBadRequest)
	}

	fruit, err := f.foodService.InsertFood(r.Context(), req.Name, req.Price)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(fruit)
}

func (f *FoodHandler) GetHelloText(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(f.foodService.GetHello())
}

package services

import (
	"context"

	"github.com/eif-courses/restaurant/internal/repository"
)

// Panasu i klase
type FoodService struct {
	queries *repository.Queries
}

// paprasta funkcija (atitikmuo konstruktorius)
func NewFoodService(queries *repository.Queries) *FoodService {
	return &FoodService{queries: queries}
}

func (f *FoodService) GetFoodList(ctx context.Context) ([]repository.Food, error) {
	return f.queries.GetAllFood(ctx)
}
func (f *FoodService) InsertFood(ctx context.Context, name string, price float64) (*repository.Food, error) {

	params := repository.InsertFoodParams{
		Name:  name,
		Price: price,
	}

	fruit, err := f.queries.InsertFood(ctx, params)

	if err != nil {
		return nil, err
	}

	return &fruit, nil
}

func (f *FoodService) GetHello() string {
	return "Hello World"
}

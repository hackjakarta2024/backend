package repository

import (
	"github.com/google/uuid"
	"github.com/hackjakarta2024/backend/internal/model"
	"gorm.io/gorm"
)

type restaurantRepository struct {
	DB *gorm.DB
}

type RestaurantRepository interface {
	GetRestaurantByID(restaurantID uuid.UUID) (model.Restaurant, error)
}

func NewRestaurantRepository(db *gorm.DB) RestaurantRepository {
	return &restaurantRepository{
		DB: db,
	}
}

func (r *restaurantRepository) GetRestaurantByID(restaurantID uuid.UUID) (model.Restaurant, error) {
	var restaurant model.Restaurant
	err := r.DB.Where("id = ?", restaurantID).First(&restaurant).Error
	return restaurant, err
}

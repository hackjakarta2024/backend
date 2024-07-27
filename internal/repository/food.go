package repository

import (
	"github.com/google/uuid"
	"github.com/hackjakarta2024/backend/internal/model"
	"gorm.io/gorm"
)

type foodRepository struct {
	DB *gorm.DB
}

type FoodRepository interface {
	GetFoodByID(foodID uuid.UUID) (model.Food, error)
}

func NewFoodRepository(db *gorm.DB) FoodRepository {
	return &foodRepository{
		DB: db,
	}
}

func (r *foodRepository) GetFoodByID(foodID uuid.UUID) (model.Food, error) {
	var food model.Food
	err := r.DB.Where("id = ?", foodID).First(&food).Error
	return food, err

}

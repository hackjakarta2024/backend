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
	GetHistoryUserFoodByFoodID(foodID uuid.UUID) ([]model.HistoryUserFood, error)
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

func (r *foodRepository) GetHistoryUserFoodByFoodID(foodID uuid.UUID) ([]model.HistoryUserFood, error) {
	var historyUserFood []model.HistoryUserFood
	err := r.DB.Raw("SELECT * FROM histories_user_food WHERE food_id = ?", foodID).Scan(&historyUserFood).Error
	//err := r.DB.Table("histories_user_food").Where("food_id = ?", foodID).Find(&historyUserFood).Error
	return historyUserFood, err
}

package model

import (
	"github.com/google/uuid"
	"time"
)

type Food struct {
	ID           uuid.UUID `json:"id"`
	RestaurantID uuid.UUID `json:"restaurant_id"`
	Name         string    `json:"name"`
	FakePrice    int       `json:"fake_price"`
	RealPrice    int       `json:"real_price"`
	Image        string    `json:"image"`
	RatingTotal  int       `json:"rating_total"`
}

type FoodRecommendation struct {
	FoodID string `json:"food_id" bigquery:"food_id"`
	Desc   string `json:"desc" bigquery:"description"`
}

type HistoryUserFood struct {
	UserID    uuid.UUID `json:"user_id" gorm:"column:user_id"`
	FoodID    uuid.UUID `json:"food_id" gorm:"column:food_id"`
	Rating    int       `json:"rating" gorm:"column:rating"`
	Review    string    `json:"review" gorm:"column:review"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}

package model

import "github.com/google/uuid"

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

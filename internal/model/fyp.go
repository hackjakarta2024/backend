package model

import (
	"github.com/google/uuid"
)

type Fyp struct {
	UserID              string               `json:"user_id" bigquery:"user_id"`
	PromoID             string               `json:"promo_id" bigquery:"promo_id"`
	Period              string               `json:"period" bigquery:"period"`
	FoodRecommendations []FoodRecommendation `json:"food_recommendations" bigquery:"food_recommendations"`
}

type FypResponse struct {
	UserID uuid.UUID      `json:"user_id"`
	Promo  Promo          `json:"promo"`
	Food   []FoodResponse `json:"food"`
}

type FypResponseAI struct {
	UserID              uuid.UUID    `json:"user_id"`
	PromoID             uuid.UUID    `json:"promo_id"`
	FoodRecommendations []FoodRespAI `json:"food_recommendations"`
	Time                string       `json:"time"`
}

type FoodResponse struct {
	ID             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	RestaurantName string    `json:"restaurant_name"`
	Desc           string    `json:"desc"`
	FakePrice      int       `json:"fake_price"`
	RealPrice      int       `json:"real_price"`
	Image          string    `json:"image"`
	RatingTotal    int       `json:"rating_total"`
}

type FoodRespAI struct {
	FoodID uuid.UUID `json:"food_id"`
	Desc   string    `json:"desc"`
}

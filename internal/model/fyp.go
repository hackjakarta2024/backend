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

type SearchResponse struct {
	UserID uuid.UUID      `json:"user_id"`
	Food   []FoodResponse `json:"food"`
}

type SearchResponseAI struct {
	UserID uuid.UUID    `json:"user_id"`
	Food   []FoodRespAI `json:"food_recommendations"`
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
	UserReview     []Review  `json:"user_review"`
}

type Review struct {
	Name   string `json:"name"`
	Review string `json:"review"`
	Rating int    `json:"rating"`
}

type FoodRespAI struct {
	FoodID uuid.UUID `json:"food_id"`
	Desc   string    `json:"desc"`
}

type Response struct {
	Data   SearchResponseAI `json:"data"`
	Status string           `json:"status"`
}

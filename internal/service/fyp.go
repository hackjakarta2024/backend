package service

import (
	"github.com/google/uuid"
	"github.com/hackjakarta2024/backend/internal/model"
	"github.com/hackjakarta2024/backend/internal/repository"
	"go.uber.org/zap"
)

type fypService struct {
	fypRepository        repository.FypRepository
	promoRepository      repository.PromoRepository
	foodRepository       repository.FoodRepository
	restaurantRepository repository.RestaurantRepository
	logger               *zap.Logger
}

type FypService interface {
	GetFyp(userID uuid.UUID) (model.FypResponse, error)
	Search(userID, query string) (model.SearchResponse, error)
}

func NewFypService(
	fypRepository repository.FypRepository,
	promoRepository repository.PromoRepository,
	foodRepository repository.FoodRepository,
	restaurantRepository repository.RestaurantRepository,
	logger *zap.Logger,
) FypService {
	return &fypService{
		fypRepository:        fypRepository,
		promoRepository:      promoRepository,
		foodRepository:       foodRepository,
		restaurantRepository: restaurantRepository,
		logger:               logger,
	}
}

func (s *fypService) GetFyp(userID uuid.UUID) (model.FypResponse, error) {
	fyp, err := s.fypRepository.GetFypByUserID(userID)
	if err != nil {
		s.logger.Error("Error getting FYP by user ID", zap.Error(err))
		return model.FypResponse{}, err
	}

	promo, err := s.promoRepository.GetPromoByID(uuid.MustParse(fyp.PromoID))
	if err != nil {
		s.logger.Error("Error getting promo by ID", zap.Error(err))
		return model.FypResponse{}, err
	}

	var fypResp model.FypResponse
	fypResp.UserID = uuid.MustParse(fyp.UserID)
	fypResp.Promo = promo

	for _, foodRecommendation := range fyp.FoodRecommendations {
		food, err := s.foodRepository.GetFoodByID(uuid.MustParse(foodRecommendation.FoodID))
		if err != nil {
			s.logger.Error("Error getting food by ID", zap.Error(err))
			return model.FypResponse{}, err
		}

		restaurant, err := s.restaurantRepository.GetRestaurantByID(food.RestaurantID)
		if err != nil {
			s.logger.Error("Error getting restaurant by ID", zap.Error(err))
			return model.FypResponse{}, err
		}

		foodResp := model.FoodResponse{
			ID:             food.ID,
			Name:           food.Name,
			RestaurantName: restaurant.Name,
			Desc:           foodRecommendation.Desc,
			FakePrice:      food.FakePrice,
			RealPrice:      food.RealPrice,
			Image:          food.Image,
			RatingTotal:    food.RatingTotal,
		}

		fypResp.Food = append(fypResp.Food, foodResp)
	}

	return fypResp, nil
}

func (s *fypService) Search(userID, query string) (model.SearchResponse, error) {
	searchRespAI := model.SearchResponseAI{
		UserID: uuid.MustParse("290fbc73-84f1-4a09-aa1f-1b09bbc2539e"),
		Food: []model.FoodRespAI{
			{
				FoodID: uuid.MustParse("2083e20f-30f2-4d0d-ac20-e25530154f77"),
				Desc:   "desc",
			},
		},
	}

	var searchResp model.SearchResponse
	searchResp.UserID = searchRespAI.UserID
	for _, sra := range searchRespAI.Food {
		food, err := s.foodRepository.GetFoodByID(sra.FoodID)
		if err != nil {
			s.logger.Error("Error getting food by ID", zap.Error(err))
			return model.SearchResponse{}, err
		}

		restaurant, err := s.restaurantRepository.GetRestaurantByID(food.RestaurantID)
		if err != nil {
			s.logger.Error("Error getting restaurant by ID", zap.Error(err))
			return model.SearchResponse{}, err
		}

		searchResp.Food = append(searchResp.Food, model.FoodResponse{
			ID:             sra.FoodID,
			Name:           food.Name,
			RestaurantName: restaurant.Name,
			Desc:           sra.Desc,
			FakePrice:      food.FakePrice,
			RealPrice:      food.RealPrice,
			Image:          food.Image,
			RatingTotal:    food.RatingTotal,
		})
	}

	return searchResp, nil
}

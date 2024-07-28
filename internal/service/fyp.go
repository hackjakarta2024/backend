package service

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/hackjakarta2024/backend/internal/model"
	"github.com/hackjakarta2024/backend/internal/repository"
	"go.uber.org/zap"
	"io"
	"net/http"
)

type fypService struct {
	fypRepository        repository.FypRepository
	promoRepository      repository.PromoRepository
	foodRepository       repository.FoodRepository
	restaurantRepository repository.RestaurantRepository
	userRepository       repository.UserRepository
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
	userRepository repository.UserRepository,
	logger *zap.Logger,
) FypService {
	return &fypService{
		fypRepository:        fypRepository,
		promoRepository:      promoRepository,
		foodRepository:       foodRepository,
		restaurantRepository: restaurantRepository,
		userRepository:       userRepository,
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

		historyUserFood, err := s.foodRepository.GetHistoryUserFoodByFoodID(uuid.MustParse(foodRecommendation.FoodID))
		if err != nil {
			s.logger.Error("Error getting history user food by food ID", zap.Error(err))
			return model.FypResponse{}, err
		}

		var userReview []model.Review
		for _, huf := range historyUserFood {
			user, err := s.userRepository.GetUserByID(huf.UserID.String())
			if err != nil {
				s.logger.Error("Error getting user by ID", zap.Error(err))
				return model.FypResponse{}, err
			}

			userReview = append(userReview, model.Review{
				Name:   user.Name,
				Review: huf.Review,
				Rating: huf.Rating,
			})
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
			UserReview:     userReview,
		}

		fypResp.Food = append(fypResp.Food, foodResp)
	}

	return fypResp, nil
}

func (s *fypService) Search(userID, query string) (model.SearchResponse, error) {
	url := fmt.Sprintf("https://hackml.timbangkit.cloud/search?user_id=%s&query=%s", userID, query)

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		s.logger.Error("Error creating request", zap.Error(err))
		return model.SearchResponse{}, err
	}

	resp, err := client.Do(req)
	if err != nil {
		s.logger.Error("Error making request", zap.Error(err))
		return model.SearchResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		s.logger.Error("Error reading response body", zap.Error(err))
		return model.SearchResponse{}, err
	}

	var respAI model.Response
	err = json.Unmarshal(body, &respAI)
	if err != nil {
		s.logger.Error("Error unmarshalling response body", zap.Error(err))
		return model.SearchResponse{}, err
	}

	var searchResp model.SearchResponse
	searchResp.UserID = respAI.Data.UserID
	for _, sra := range respAI.Data.Food {
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

		historyUserFood, err := s.foodRepository.GetHistoryUserFoodByFoodID(sra.FoodID)
		if err != nil {
			s.logger.Error("Error getting history user food by food ID", zap.Error(err))
			return model.SearchResponse{}, err
		}

		var userReview []model.Review
		for _, huf := range historyUserFood {
			user, err := s.userRepository.GetUserByID(huf.UserID.String())
			if err != nil {
				s.logger.Error("Error getting user by ID", zap.Error(err))
				return model.SearchResponse{}, err
			}

			userReview = append(userReview, model.Review{
				Name:   user.Name,
				Review: huf.Review,
				Rating: huf.Rating,
			})
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
			UserReview:     userReview,
		})
	}

	return searchResp, nil
}

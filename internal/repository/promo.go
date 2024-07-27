package repository

import (
	"github.com/google/uuid"
	"github.com/hackjakarta2024/backend/internal/model"
	"gorm.io/gorm"
)

type promoRepository struct {
	DB *gorm.DB
}

type PromoRepository interface {
	GetPromoByID(promoID uuid.UUID) (model.Promo, error)
}

func NewPromoRepository(db *gorm.DB) PromoRepository {
	return &promoRepository{
		DB: db,
	}
}

func (r *promoRepository) GetPromoByID(promoID uuid.UUID) (model.Promo, error) {
	var promo model.Promo
	err := r.DB.Where("id = ?", promoID).First(&promo).Error
	return promo, err
}

package repository

import (
	"cloud.google.com/go/bigquery"
	"context"
	"github.com/google/uuid"
	"github.com/hackjakarta2024/backend/internal/model"
	"go.uber.org/zap"
	"google.golang.org/api/iterator"
	"gorm.io/gorm"
)

type fypRepository struct {
	DB     *gorm.DB
	BQ     *bigquery.Client
	logger *zap.Logger
}

type FypRepository interface {
	GetFypByUserID(userID uuid.UUID) (model.Fyp, error)
}

func NewFypRepository(db *gorm.DB, BQ *bigquery.Client, logger *zap.Logger) FypRepository {
	return &fypRepository{
		DB:     db,
		BQ:     BQ,
		logger: logger,
	}
}

func (r *fypRepository) GetFypByUserID(userID uuid.UUID) (model.Fyp, error) {
	var fyp model.Fyp
	query := r.BQ.Query(`
    	SELECT * 
    	FROM ` + "`hack-jakarta.hackjakarta.recommendation`" + `
    	WHERE user_id = @userID
    	LIMIT 1
	`)
	query.Parameters = []bigquery.QueryParameter{
		{
			Name:  "userID",
			Value: userID.String(),
		},
	}
	it, err := query.Read(context.Background())
	if err != nil {
		r.logger.Error("Error reading from BigQuery", zap.Error(err))
		return model.Fyp{}, err
	}
	for {
		err = it.Next(&fyp)
		if err == iterator.Done {
			break
		}
		if err != nil {
			r.logger.Error("Error reading from BigQuery", zap.Error(err))
			return model.Fyp{}, err
		}

	}

	return fyp, nil
}

package config

import (
	"cloud.google.com/go/bigquery"
	"context"
	"os"
)

func NewBigQuery() (*bigquery.Client, error) {
	client, err := bigquery.NewClient(context.Background(), os.Getenv("PROJECT_ID"))
	if err != nil {
		return nil, err
	}

	return client, nil
}

func CloseBigQuery(dbbq *bigquery.Client) {
	dbbq.Close()
}

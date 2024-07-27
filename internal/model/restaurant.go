package model

import "github.com/google/uuid"

type Restaurant struct {
	ID     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	Rating int       `json:"rating"`
}

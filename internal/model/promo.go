package model

import "github.com/google/uuid"

type Promo struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

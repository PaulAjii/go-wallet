package models

import (
	"time"

	"github.com/google/uuid"
)

type Category string
type Status string

type BaseModel struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}

type BaseWithoutUpdatedAt struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}

const (
	Funding    Category = "funding"
	Transfer   Category = "transfer"
	Withdrawal Category = "withdrawal"
)

const (
	StatusPending Status = "pending"
	StatusSuccess Status = "success"
	StatusFailed  Status = "failed"
)

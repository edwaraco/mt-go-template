package domain

import (
	"time"
)

type MTModel struct {
	CreatedBy string
	CreatedAt time.Time
	UpdatedBy string
	UpdatedAt time.Time
	// TODO: Check if this can apply
	DeletedBy string
	DeletedAt *time.Time
}

package entity

import (
	"context"
)

type Repository interface {
	GetByDate(ctx context.Context, date, code string) ([]*Rate, error)
	Create(ctx context.Context, entity *Rate) error
}

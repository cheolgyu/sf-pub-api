package domain

import (
	"context"
)

type Info struct {
}

type InfoUsecase interface {
	GetUpdateTime(ctx context.Context) (string, error)
}

type InfoRepository interface {
	GetUpdateTime(ctx context.Context) (string, error)
}

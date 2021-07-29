package meta

import (
	"context"
	"time"

	"github.com/cheolgyu/stock-read-pub-api/src/domain"
)

type MetaUsecase struct {
	metaRepo domain.MetaRepository
	//authorRepo     domain.AuthorRepository
	contextTimeout time.Duration
}

func NewUsecase(cr domain.MetaRepository, timeout time.Duration) domain.MetaUsecase {
	return &MetaUsecase{
		metaRepo:       cr,
		contextTimeout: timeout,
	}
}

func (obj *MetaUsecase) GetConfig(ctx context.Context) ([]domain.Config, error) {
	return obj.metaRepo.GetConfig(context.TODO())
}

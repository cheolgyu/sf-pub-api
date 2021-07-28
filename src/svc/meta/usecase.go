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

func (obj *MetaUsecase) GetMarketList(ctx context.Context) ([]domain.Config, error) {
	return obj.metaRepo.GetMarketList(context.TODO())
}

func (obj *MetaUsecase) GetColumnName(ctx context.Context, table_name string) ([]string, error) {
	return obj.metaRepo.GetColumnName(context.TODO(), table_name)
}

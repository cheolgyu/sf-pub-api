package info

import (
	"context"
	"time"

	"github.com/cheolgyu/stock-read-pub-api/src/domain"
)

type InfoUsecase struct {
	infoRepo       domain.InfoRepository
	contextTimeout time.Duration
}

func NewUsecase(cr domain.InfoRepository, timeout time.Duration) domain.InfoUsecase {

	return &InfoUsecase{
		infoRepo:       cr,
		contextTimeout: timeout,
	}
}

func (obj *InfoUsecase) GetUpdateTime(ctx context.Context) (string, error) {
	return obj.infoRepo.GetUpdateTime(context.TODO())
}

package company

import (
	"context"
	"time"

	"github.com/cheolgyu/stock-read-pub-api/src/domain"
)

type CompanyUsecase struct {
	companyRepo domain.CompanyRepository
	//authorRepo     domain.AuthorRepository
	contextTimeout time.Duration
}

func NewUsecase(cr domain.CompanyRepository, timeout time.Duration) domain.CompanyUsecase {
	return &CompanyUsecase{
		companyRepo:    cr,
		contextTimeout: timeout,
	}
}

func (obj *CompanyUsecase) GetByCode(ctx context.Context, code string) (string, error) {
	return obj.companyRepo.GetByCode(context.TODO(), code)
}

func (obj *CompanyUsecase) GetGraphByCodeID(ctx context.Context, code string, page int) (string, error) {
	return obj.companyRepo.GetGraphByCodeID(context.TODO(), code, page)
}

func (obj *CompanyUsecase) GetGraphNextLineByCode(ctx context.Context, code string) (string, error) {
	return obj.companyRepo.GetGraphNextLineByCode(context.TODO(), code)
}

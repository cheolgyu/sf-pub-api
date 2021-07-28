package usecase

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

func NewCompanyUsecase(cr domain.CompanyRepository, timeout time.Duration) domain.CompanyUsecase {
	return &CompanyUsecase{
		companyRepo:    cr,
		contextTimeout: timeout,
	}
}

func (obj *CompanyUsecase) GetByCode(ctx context.Context, code string) (domain.Company, error) {

	var res domain.Company
	var err error = nil

	return res, err
}

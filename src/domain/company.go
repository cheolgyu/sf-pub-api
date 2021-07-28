package domain

import "context"

type Company struct {
}

type CompanyUsecase interface {
	GetByCode(ctx context.Context, code string) (Company, error)
}

type CompanyRepository interface {
	GetByCode(ctx context.Context, code string) (Company, error)
	//Fetch(ctx context.Context, cursor string, num int64) (res []Company, nextCursor string, err error)
	//GetByID(ctx context.Context, id int64) (Company, error)
	//GetByTitle(ctx context.Context, title string) (Company, error)
	//Update(ctx context.Context, ar *Company) error
	//Store(ctx context.Context, a *Company) error
	//Delete(ctx context.Context, id int64) error
}

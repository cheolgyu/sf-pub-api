package repo

import (
	"context"

	"github.com/cheolgyu/stock-read-pub-api/src/domain"
	"github.com/jmoiron/sqlx"
)

type CompanyRepository struct {
	conn *sqlx.DB
}

func NewCompnayRepository(Conn *sqlx.DB) domain.CompanyRepository {

	return &CompanyRepository{conn: Conn}
}

func (obj *CompanyRepository) GetByCode(ctx context.Context, code string) (domain.Company, error) {

	var res domain.Company
	var err error = nil

	return res, err
}

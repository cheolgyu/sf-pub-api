package company

import (
	"context"
	"log"
	"time"

	"github.com/cheolgyu/pubapi/src/domain"
)

type CompanyUsecase struct {
	companyRepo domain.CompanyRepository
	//authorRepo     domain.AuthorRepository
	contextTimeout time.Duration

	var_price_type_list []domain.Config
	var_column_name     map[string][]string
}

func NewUsecase(cr domain.CompanyRepository, meta_repo domain.MetaRepository, timeout time.Duration) domain.CompanyUsecase {

	price_type_list := meta_repo.VarPriceTypeList()
	column_name := meta_repo.VarColumnName()

	return &CompanyUsecase{
		companyRepo:         cr,
		contextTimeout:      timeout,
		var_price_type_list: price_type_list,
		var_column_name:     column_name,
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

func (obj *CompanyUsecase) GetReboundByPaging(ctx context.Context, params_string domain.CompanyHisteParamsString) ([]map[string]interface{}, error) {

	paging, err := params_string.Valid(obj.var_price_type_list, obj.var_column_name, "rebound")
	if err != nil {
		log.Println(err)
		panic(err)
	}

	return obj.companyRepo.GetReboundByPaging(context.TODO(), paging)
}

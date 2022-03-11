package project

import (
	"context"
	"log"
	"time"

	"github.com/cheolgyu/sf-pub-api/src/domain"
)

type ProjectUsecase struct {
	projectRepo    domain.ProjectRepository
	contextTimeout time.Duration

	var_market_list []domain.Config
	var_column_name map[string][]string
}

func NewUsecase(cr domain.ProjectRepository, meta_repo domain.MetaRepository, timeout time.Duration) domain.ProjectUsecase {
	market_list := meta_repo.VarMarketList()
	column_name := meta_repo.VarColumnName()

	return &ProjectUsecase{
		projectRepo:     cr,
		contextTimeout:  timeout,
		var_market_list: market_list,
		var_column_name: column_name,
	}
}

func (obj *ProjectUsecase) GetDayTradingByPaging(ctx context.Context, params_string domain.ProjectParamsString) ([]map[string]interface{}, error) {
	log.Println("~~~~~~~~~~~~~GetDayTradingByPaging~~~~~~~~~~~~~~~~~~~~~~~", obj.var_market_list)
	paging, err := params_string.Valid(obj.var_market_list, nil, "")
	if err != nil {
		log.Println(err)
		panic(err)
	}
	return obj.projectRepo.GetDayTradingByPaging(context.TODO(), paging)
}

func (obj *ProjectUsecase) GetMonthlyPeekByPaging(ctx context.Context, params_string domain.ProjectParamsString) ([]map[string]interface{}, error) {

	paging, err := params_string.Valid(obj.var_market_list, obj.var_column_name, "view_monthly_peek")
	if err != nil {
		log.Println(err)
		panic(err)
	}

	return obj.projectRepo.GetMonthlyPeekByPaging(context.TODO(), paging)
}

package price

import (
	"context"
	"log"
	"time"

	"github.com/cheolgyu/pubapi/src/domain"
)

type PriceUsecase struct {
	priceRepo      domain.PriceRepository
	contextTimeout time.Duration

	var_market_list []domain.Config
	var_column_name map[string][]string
}

func NewUsecase(cr domain.PriceRepository, meta_repo domain.MetaRepository, timeout time.Duration) domain.PriceUsecase {
	market_list := meta_repo.VarMarketList()
	column_name := meta_repo.VarColumnName()

	return &PriceUsecase{
		priceRepo:       cr,
		contextTimeout:  timeout,
		var_market_list: market_list,
		var_column_name: column_name,
	}
}

func (obj *PriceUsecase) GetStockByPaging(ctx context.Context, params_string domain.PriceParamsString) ([]map[string]interface{}, error) {

	paging, err := params_string.Valid(obj.var_market_list, obj.var_column_name, "view_stock")
	if err != nil {
		log.Println(err)
		panic(err)
	}
	return obj.priceRepo.GetStockByPaging(context.TODO(), paging)
}

func (obj *PriceUsecase) GetMarketByPaging(ctx context.Context, params_string domain.PriceParamsString) ([]map[string]interface{}, error) {

	paging, err := params_string.Valid(obj.var_market_list, obj.var_column_name, "view_market")
	if err != nil {
		log.Println(err)
		panic(err)
	}

	return obj.priceRepo.GetMarketByPaging(context.TODO(), paging)
}

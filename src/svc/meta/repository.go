package meta

import (
	"context"
	"fmt"
	"log"

	"github.com/cheolgyu/sf-pub-api/src/domain"
	"github.com/jmoiron/sqlx"
)

type MetaRepository struct {
	conn          *sqlx.DB
	ColumnName    map[string][]string
	MarketList    []domain.Config
	PriceTypeList []domain.Config
}

func NewRepository(Conn *sqlx.DB) domain.MetaRepository {
	repo := &MetaRepository{conn: Conn}
	repo.ColumnName = make(map[string][]string)
	repo.init()

	return repo
}
func (obj *MetaRepository) init() (err error) {
	obj.setConfig()
	var sort_table []string
	sort_table = append(sort_table, "view_stock", "view_market", "rebound", "view_monthly_peek", "company_state")
	for _, v := range sort_table {
		if res, err := obj.setColumnName(v); err == nil {
			obj.ColumnName[v] = res
		}
	}

	return err
}

func (obj *MetaRepository) setColumnName(tb_name string) (res []string, err error) {
	res, err = obj.GetColumnName(context.TODO(), tb_name)
	return res, err
}

func (obj *MetaRepository) setConfig() (res []domain.Config, err error) {
	if res, err = obj.GetConfig(context.TODO()); err != nil {
		log.Fatalln(err)
	}
	for i := 0; i < len(res); i++ {
		if res[i].Upper_code == "market_type" {
			obj.MarketList = append(obj.MarketList, res[i])
		} else if res[i].Upper_code == "price_type" {
			obj.PriceTypeList = append(obj.PriceTypeList, res[i])
		}
	}
	return res, err
}

func (obj *MetaRepository) VarMarketList() []domain.Config {
	return obj.MarketList
}

func (obj *MetaRepository) VarPriceTypeList() []domain.Config {
	return obj.PriceTypeList
}

func (obj *MetaRepository) VarColumnName() map[string][]string {
	return obj.ColumnName
}

func (obj *MetaRepository) GetConfig(ctx context.Context) (res []domain.Config, err error) {

	q := `
	SELECT  *  FROM  meta.config where  upper_code = 'market_type' and code in ('KOSPI','KOSDAQ','KONEX')
		union all
	SELECT  *  FROM  meta.config where  upper_code = 'price_type'
		union all
	SELECT  *  FROM  meta.config where  upper_code = 'unit_type'
	`
	rows, err := obj.conn.Queryx(q)

	for rows.Next() {
		item := domain.Config{}

		err = rows.StructScan(&item)
		if err != nil {
			log.Printf("GetMarketList:MapScan::error::::<%s>  \n", err)
			panic(err)
		}
		res = append(res, item)

	}

	return res, err
}

func (obj *MetaRepository) GetColumnName(ctx context.Context, tb_nm string) (res []string, err error) {

	q := `SELECT column_name FROM INFORMATION_SCHEMA.COLUMNS WHERE table_name = '%s'; `
	query := fmt.Sprintf(q, tb_nm)
	rows, err := obj.conn.Queryx(query)

	for rows.Next() {
		item := ""
		err = rows.Scan(&item)
		if err != nil {
			log.Printf("GetMarketList:MapScan::error::::<%s>  \n", err)
			panic(err)
		}
		res = append(res, item)

	}

	return res, err
}

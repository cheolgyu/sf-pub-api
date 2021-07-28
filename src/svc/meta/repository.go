package meta

import (
	"context"
	"fmt"
	"log"

	"github.com/cheolgyu/stock-read-pub-api/src/domain"
	"github.com/jmoiron/sqlx"
)

type MetaRepository struct {
	conn       *sqlx.DB
	ColumnName map[string][]string
	MarketList []domain.Config
}

func NewRepository(Conn *sqlx.DB) domain.MetaRepository {
	repo := &MetaRepository{conn: Conn}
	repo.ColumnName = make(map[string][]string)
	repo.init()

	return repo
}
func (obj *MetaRepository) init() (err error) {
	obj.setMarketList()
	obj.setColumnName("view_stock")
	obj.setColumnName("view_market")

	return err
}

func (obj *MetaRepository) setColumnName(tb_name string) (res []string, err error) {
	res, err = obj.GetColumnName(context.TODO(), tb_name)
	return res, err
}

func (obj *MetaRepository) setMarketList() (res []domain.Config, err error) {
	if res, err = obj.GetMarketList(context.TODO()); err != nil {
		log.Fatalln(err)
	}
	obj.MarketList = res
	return res, err
}

func (obj *MetaRepository) VarMarketList() []domain.Config {
	return obj.MarketList
}

func (obj *MetaRepository) VarColumnName() map[string][]string {
	return obj.ColumnName
}

func (obj *MetaRepository) GetMarketList(ctx context.Context) (res []domain.Config, err error) {

	q := `SELECT  *  FROM  meta.config where  upper_code = 'market_type' and code in ('KOSPI','KOSDAQ','KONEX') `
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
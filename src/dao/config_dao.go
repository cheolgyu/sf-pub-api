package dao

import (
	"log"

	"github.com/cheolgyu/stock-read-pub-api/src/model"
)

var SqlConfigDao ConfigRepoDao
var Tag string

type ConfigRepoDao struct {
	tag string
}

func NewConfigRepo() *ConfigRepoDao {
	return &ConfigRepoDao{}
}

type ConfigRepo interface {
	GetMarketList(req_id string) []model.Config
}

func init() {
	SqlConfigDao = ConfigRepoDao{tag: "SqlConfigDao"}
}

func (obj *ConfigRepoDao) GetMarketList(req_id string) []model.Config {

	q := `  SELECT  *  FROM  meta.config where  upper_code = 'market_type' and code in ('KOSPI','KOSDAQ','KONEX') `

	rows, err := DB.Queryx(q)

	if err != nil {
		log.Printf(obj.tag, "::error::::<%s>  \n", req_id)
		log.Printf(obj.tag, "::error::::<%s> query= \n", q)
		panic(err)
	}

	defer rows.Close()

	var list []model.Config

	for rows.Next() {
		item := model.Config{}

		err = rows.StructScan(&item)
		if err != nil {
			log.Printf(obj.tag, ":MapScan::error::::<%s>  \n", req_id)
			panic(err)
		}
		list = append(list, item)

	}
	return list
}

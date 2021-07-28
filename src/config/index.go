package config

import "github.com/cheolgyu/stock-read-pub-api/src/model"

var MarketList []model.Config

func init() {
	init_config_market_list()
}

func init_config_market_list() {
	// config_repo := dao.NewConfigRepo()
	// MarketList = config_repo.GetMarketList(" init_config_data ")
}

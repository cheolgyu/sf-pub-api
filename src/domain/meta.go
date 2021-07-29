package domain

import "context"

type Meta struct {
}

type MetaUsecase interface {
	GetMarketList(ctx context.Context) ([]Config, error)
}

type MetaRepository interface {
	VarMarketList() []Config
	VarPriceTypeList() []Config
	VarColumnName() map[string][]string

	GetMarketList(ctx context.Context) ([]Config, error)
	GetPriceTypeList(ctx context.Context) ([]Config, error)
	GetColumnName(ctx context.Context, table string) ([]string, error)
}

type Config struct {
	Id         int    `json:"id"`
	Upper_code string `json:"upper_code"`
	Upper_name string `json:"upper_name"`
	Code       string `json:"code"`
	Name       string `json:"name"`
}

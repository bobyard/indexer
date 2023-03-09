package models

import (
	"time"
)

type Lists struct {
	Id            int       `xorm:"SERIAL"`
	ChainId       int64     `xorm:"not null BIGINT"`
	TokenId       string    `xorm:"not null VARCHAR"`
	SellerAddress string    `xorm:"not null TEXT"`
	SallerValue   int64     `xorm:"not null BIGINT"`
	SellerCoinId  int       `xorm:"not null INTEGER"`
	SellerEndTime time.Time `xorm:"not null DATETIME"`
	CreatedAt     time.Time `xorm:"default now() DATETIME"`
	UpdatedAt     time.Time `xorm:"default now() DATETIME"`
}

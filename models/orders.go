package models

import (
	"time"
)

type Orders struct {
	Id            int       `xorm:"not null pk default nextval('orders_id_seq'::regclass) autoincr INTEGER"`
	TokenId       string    `xorm:"not null VARCHAR"`
	SellerAddress string    `xorm:"not null VARCHAR"`
	BuyerAddress  string    `xorm:"not null VARCHAR"`
	Amount        string    `xorm:"not null VARCHAR"`
	CoinId        int       `xorm:"not null INTEGER"`
	ChainId       int       `xorm:"not null INTEGER"`
	Time          time.Time `xorm:"not null DATETIME"`
}

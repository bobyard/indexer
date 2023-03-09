package models

type Coins struct {
	Id           int    `xorm:"SERIAL"`
	ChainId      int64  `xorm:"not null pk BIGINT"`
	Symbol       string `xorm:"not null VARCHAR"`
	DecimalPoint string `xorm:"not null VARCHAR"`
	Logo         string `xorm:"not null VARCHAR"`
}

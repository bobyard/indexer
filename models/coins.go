package models

type Coins struct {
	Id           int    `xorm:"not null pk default nextval('coins_id_seq'::regclass) autoincr INTEGER"`
	ChainId      int64  `xorm:"not null pk BIGINT"`
	Symbol       string `xorm:"not null VARCHAR"`
	DecimalPoint string `xorm:"not null VARCHAR"`
	Logo         string `xorm:"not null VARCHAR"`
}

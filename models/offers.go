package models

type Offers struct {
	Id           int    `xorm:"not null pk default nextval('offers_id_seq'::regclass) autoincr INTEGER"`
	TokenId      string `xorm:"VARCHAR"`
	BuyerAddress string `xorm:"VARCHAR"`
	Item         string `xorm:"VARCHAR"`
	Amount       string `xorm:"VARCHAR"`
	ChainId      int    `xorm:"INTEGER"`
	OfferId      string `xorm:"VARCHAR"`
}

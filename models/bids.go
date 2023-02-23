package models

type Bids struct {
	Id int `xorm:"not null pk default nextval('bids_id_seq'::regclass) autoincr INTEGER"`
}

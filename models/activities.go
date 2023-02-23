package models

import (
	"time"
)

type Activities struct {
	Id                   int       `xorm:"not null pk default nextval('activities_id_seq'::regclass) autoincr INTEGER"`
	ChainId              int64     `xorm:"not null BIGINT"`
	Version              int64     `xorm:"not null pk unique(event_index) index BIGINT"`
	EventAccountAddress  string    `xorm:"not null unique(event_index) TEXT"`
	EventCreationNumber  int64     `xorm:"not null unique(event_index) BIGINT"`
	EventSequenceNumber  int64     `xorm:"not null unique(event_index) BIGINT"`
	CollectionDataIdHash string    `xorm:"not null TEXT"`
	TokenDataIdHash      string    `xorm:"not null index(ta_tdih_pv_index) TEXT"`
	PropertyVersion      string    `xorm:"not null index(ta_addr_coll_name_pv_index) index(ta_tdih_pv_index) NUMERIC"`
	CreatorAddress       string    `xorm:"not null index(ta_addr_coll_name_pv_index) TEXT"`
	CollectionName       string    `xorm:"not null index(ta_addr_coll_name_pv_index) TEXT"`
	Name                 string    `xorm:"not null index(ta_addr_coll_name_pv_index) TEXT"`
	TransferType         string    `xorm:"not null index(ta_from_ttyp_index) index(ta_to_ttyp_index) TEXT"`
	FromAddress          string    `xorm:"index(ta_from_ttyp_index) TEXT"`
	ToAddress            string    `xorm:"index(ta_to_ttyp_index) TEXT"`
	TokenAmount          string    `xorm:"not null NUMERIC"`
	CoinType             string    `xorm:"TEXT"`
	CoinAmount           string    `xorm:"NUMERIC"`
	TransactionTimestamp time.Time `xorm:"not null DATETIME"`
	CreatedAt            time.Time `xorm:"default now() DATETIME"`
	UpdatedAt            time.Time `xorm:"default now() DATETIME"`
}

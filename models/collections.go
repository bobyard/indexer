package models

import (
	"time"
)

type Collections struct {
	Id               int       `xorm:"SERIAL"`
	ChainId          int64     `xorm:"not null BIGINT"`
	Slug             string    `xorm:"TEXT"`
	CollectionId     string    `xorm:"not null unique VARCHAR"`
	CreatorAddress   string    `xorm:"not null VARCHAR"`
	CollectionName   string    `xorm:"not null VARCHAR"`
	Description      string    `xorm:"not null VARCHAR"`
	Supply           int64     `xorm:"not null BIGINT"`
	Version          int64     `xorm:"not null BIGINT"`
	MetadataUri      string    `xorm:"not null VARCHAR"`
	FloorSellId      int       `xorm:"INTEGER"`
	FloorSellValue   string    `xorm:"NUMERIC"`
	FloorSellCoinId  int       `xorm:"INTEGER"`
	BestBidId        int       `xorm:"INTEGER"`
	BestBidValue     string    `xorm:"NUMERIC"`
	BestBidCoinId    int       `xorm:"INTEGER"`
	Verify           bool      `xorm:"not null default false BOOL"`
	LastMetadataSync time.Time `xorm:"DATETIME"`
	CreatedAt        time.Time `xorm:"default now() DATETIME"`
	UpdatedAt        time.Time `xorm:"default now() DATETIME"`
}

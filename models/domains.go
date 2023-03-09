package models

import (
	"time"
)

type Domains struct {
	Id           int       `xorm:"SERIAL"`
	ChainId      int64     `xorm:"not null BIGINT"`
	HashId       string    `xorm:"TEXT"`
	Domain       string    `xorm:"not null unique(ttyp_index) VARCHAR"`
	SubDomain    string    `xorm:"not null unique(ttyp_index) VARCHAR"`
	Description  string    `xorm:"not null VARCHAR"`
	Supply       int64     `xorm:"not null BIGINT"`
	Version      int64     `xorm:"not null pk BIGINT"`
	MetadataUri  string    `xorm:"not null VARCHAR"`
	MetadataJson string    `xorm:"VARCHAR"`
	Image        string    `xorm:"VARCHAR"`
	ExpiredTime  time.Time `xorm:"DATETIME"`
	RegestTime   time.Time `xorm:"DATETIME"`
	OnwerAddress string    `xorm:"VARCHAR"`
	CreatedAt    time.Time `xorm:"default now() DATETIME"`
	UpdatedAt    time.Time `xorm:"default now() DATETIME"`
}

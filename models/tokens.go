package models

import (
	"time"
)

type Tokens struct {
	Id                       int       `xorm:"SERIAL"`
	ChainId                  int64     `xorm:"not null BIGINT"`
	TokenId                  string    `xorm:"not null unique VARCHAR"`
	CollectionId             string    `xorm:"not null VARCHAR"`
	CreatorAddress           string    `xorm:"not null VARCHAR"`
	CollectionName           string    `xorm:"not null VARCHAR"`
	TokenName                string    `xorm:"not null VARCHAR"`
	Attributes               string    `xorm:"TEXT"`
	Supply                   int64     `xorm:"not null BIGINT"`
	Version                  int64     `xorm:"not null BIGINT"`
	PayeeAddress             string    `xorm:"not null VARCHAR"`
	RoyaltyPointsNumerator   int64     `xorm:"not null BIGINT"`
	RoyaltyPointsDenominator int64     `xorm:"not null BIGINT"`
	OnwerAddress             string    `xorm:"VARCHAR"`
	MetadataUri              string    `xorm:"not null VARCHAR"`
	MetadataJson             string    `xorm:"VARCHAR"`
	Image                    string    `xorm:"VARCHAR"`
	CreatedAt                time.Time `xorm:"default now() DATETIME"`
	UpdatedAt                time.Time `xorm:"default now() DATETIME"`
}

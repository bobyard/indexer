package suimodels

import (
	"time"
)

type Addresses struct {
	Id                  int64     `json:"id" xorm:"pk default nextval('addresses_id_seq'::regclass) autoincr BIGINT"`
	AccountAddress      string    `json:"account_address" xorm:"not null index unique VARCHAR(255)"`
	FirstAppearanceTx   string    `json:"first_appearance_tx" xorm:"not null VARCHAR(255)"`
	FirstAppearanceTime time.Time `json:"first_appearance_time" xorm:"DATETIME"`
}

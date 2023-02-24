package suimodels

import (
	"time"
)

type _DieselSchemaMigrations struct {
	Version string    `json:"version" xorm:"not null pk VARCHAR(50)"`
	RunOn   time.Time `json:"run_on" xorm:"not null default CURRENT_TIMESTAMP DATETIME"`
}

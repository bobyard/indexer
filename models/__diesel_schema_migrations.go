package models

import (
	"time"
)

type _DieselSchemaMigrations struct {
	Version string    `xorm:"not null pk VARCHAR(50)"`
	RunOn   time.Time `xorm:"not null default CURRENT_TIMESTAMP DATETIME"`
}

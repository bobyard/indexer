package suimodels

import (
	"time"
)

type ErrorLogs struct {
	Id        int64     `json:"id" xorm:"pk default nextval('error_logs_id_seq'::regclass) autoincr BIGINT"`
	ErrorType string    `json:"error_type" xorm:"not null VARCHAR(63)"`
	Error     string    `json:"error" xorm:"not null TEXT"`
	ErrorTime time.Time `json:"error_time" xorm:"not null DATETIME"`
}

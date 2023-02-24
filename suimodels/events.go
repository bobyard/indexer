package suimodels

import (
	"time"
)

type Events struct {
	Id                          int64     `json:"id" xorm:"pk default nextval('events_id_seq'::regclass) autoincr BIGINT"`
	TransactionDigest           string    `json:"transaction_digest" xorm:"not null index VARCHAR(255)"`
	EventSequence               int64     `json:"event_sequence" xorm:"not null BIGINT"`
	EventTime                   time.Time `json:"event_time" xorm:"index DATETIME"`
	EventType                   string    `json:"event_type" xorm:"not null VARCHAR"`
	EventContent                string    `json:"event_content" xorm:"not null VARCHAR"`
	NextCursorTransactionDigest string    `json:"next_cursor_transaction_digest" xorm:"VARCHAR(255)"`
}

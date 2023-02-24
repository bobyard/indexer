package suimodels

import (
	"time"
)

type PublishEvents struct {
	Id                int64     `json:"id" xorm:"pk default nextval('publish_events_id_seq'::regclass) autoincr BIGINT"`
	TransactionDigest string    `json:"transaction_digest" xorm:"index unique(publish_events_transaction_digest_event_sequence_key) VARCHAR(255)"`
	EventSequence     int64     `json:"event_sequence" xorm:"not null unique(publish_events_transaction_digest_event_sequence_key) BIGINT"`
	EventTime         time.Time `json:"event_time" xorm:"index DATETIME"`
	EventType         string    `json:"event_type" xorm:"not null VARCHAR"`
	EventContent      string    `json:"event_content" xorm:"not null VARCHAR"`
}

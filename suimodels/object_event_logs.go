package suimodels

type ObjectEventLogs struct {
	Id                 int    `json:"id" xorm:"not null pk default nextval('object_event_logs_id_seq'::regclass) autoincr INTEGER"`
	NextCursorTxDig    string `json:"next_cursor_tx_dig" xorm:"TEXT"`
	NextCursorEventSeq int64  `json:"next_cursor_event_seq" xorm:"BIGINT"`
}

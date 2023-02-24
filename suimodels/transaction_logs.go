package suimodels

type TransactionLogs struct {
	Id                 int    `json:"id" xorm:"not null pk default nextval('transaction_logs_id_seq'::regclass) autoincr INTEGER"`
	NextCursorTxDigest string `json:"next_cursor_tx_digest" xorm:"TEXT"`
}

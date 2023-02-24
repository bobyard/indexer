package suimodels

type CheckpointLogs struct {
	NextCursorSequenceNumber int64 `json:"next_cursor_sequence_number" xorm:"not null pk BIGINT"`
}

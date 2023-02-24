package suimodels

type PackageLogs struct {
	LastProcessedId int64 `json:"last_processed_id" xorm:"not null pk BIGINT"`
}

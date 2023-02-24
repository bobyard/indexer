package suimodels

type Checkpoints struct {
	SequenceNumber       int64  `json:"sequence_number" xorm:"not null pk BIGINT"`
	ContentDigest        string `json:"content_digest" xorm:"not null index VARCHAR(255)"`
	Epoch                int64  `json:"epoch" xorm:"not null index BIGINT"`
	TotalGasCost         int64  `json:"total_gas_cost" xorm:"not null BIGINT"`
	TotalComputationCost int64  `json:"total_computation_cost" xorm:"not null BIGINT"`
	TotalStorageCost     int64  `json:"total_storage_cost" xorm:"not null BIGINT"`
	TotalStorageRebate   int64  `json:"total_storage_rebate" xorm:"not null BIGINT"`
	TotalTransactions    int64  `json:"total_transactions" xorm:"not null BIGINT"`
	PreviousDigest       string `json:"previous_digest" xorm:"VARCHAR(255)"`
	NextEpochCommittee   string `json:"next_epoch_committee" xorm:"TEXT"`
	TimestampMs          int64  `json:"timestamp_ms" xorm:"not null BIGINT"`
}

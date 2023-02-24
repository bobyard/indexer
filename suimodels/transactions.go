package suimodels

import (
	"time"
)

type Transactions struct {
	Id                 int64     `json:"id" xorm:"pk default nextval('transactions_id_seq'::regclass) autoincr BIGINT"`
	TransactionDigest  string    `json:"transaction_digest" xorm:"not null index unique VARCHAR(255)"`
	Sender             string    `json:"sender" xorm:"not null index VARCHAR(255)"`
	TransactionTime    time.Time `json:"transaction_time" xorm:"index DATETIME"`
	TransactionKinds   string    `json:"transaction_kinds" xorm:"not null TEXT"`
	Created            string    `json:"created" xorm:"not null TEXT"`
	Mutated            string    `json:"mutated" xorm:"not null TEXT"`
	Deleted            string    `json:"deleted" xorm:"not null TEXT"`
	Unwrapped          string    `json:"unwrapped" xorm:"not null TEXT"`
	Wrapped            string    `json:"wrapped" xorm:"not null TEXT"`
	GasObjectId        string    `json:"gas_object_id" xorm:"not null index VARCHAR(255)"`
	GasObjectSequence  int64     `json:"gas_object_sequence" xorm:"not null BIGINT"`
	GasObjectDigest    string    `json:"gas_object_digest" xorm:"not null VARCHAR(255)"`
	GasBudget          int64     `json:"gas_budget" xorm:"not null BIGINT"`
	TotalGasCost       int64     `json:"total_gas_cost" xorm:"not null BIGINT"`
	ComputationCost    int64     `json:"computation_cost" xorm:"not null BIGINT"`
	StorageCost        int64     `json:"storage_cost" xorm:"not null BIGINT"`
	StorageRebate      int64     `json:"storage_rebate" xorm:"not null BIGINT"`
	GasPrice           int64     `json:"gas_price" xorm:"not null BIGINT"`
	TransactionContent string    `json:"transaction_content" xorm:"not null TEXT"`
}

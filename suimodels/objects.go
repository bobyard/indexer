package suimodels

type Objects struct {
	Id                   int64  `json:"id" xorm:"pk default nextval('objects_id_seq'::regclass) autoincr BIGINT"`
	ObjectId             string `json:"object_id" xorm:"not null index unique VARCHAR(255)"`
	Version              int64  `json:"version" xorm:"not null BIGINT"`
	OwnerType            string `json:"owner_type" xorm:"not null VARCHAR(255)"`
	OwnerAddress         string `json:"owner_address" xorm:"index VARCHAR(255)"`
	InitialSharedVersion int64  `json:"initial_shared_version" xorm:"BIGINT"`
	PackageId            string `json:"package_id" xorm:"not null index TEXT"`
	TransactionModule    string `json:"transaction_module" xorm:"not null TEXT"`
	ObjectType           string `json:"object_type" xorm:"TEXT"`
	ObjectStatus         string `json:"object_status" xorm:"not null VARCHAR(255)"`
}

package suimodels

type Packages struct {
	Id             int64  `json:"id" xorm:"pk default nextval('packages_id_seq'::regclass) autoincr BIGINT"`
	PackageId      string `json:"package_id" xorm:"not null index unique TEXT"`
	Author         string `json:"author" xorm:"not null TEXT"`
	ModuleNames    string `json:"module_names" xorm:"not null TEXT"`
	PackageContent string `json:"package_content" xorm:"not null TEXT"`
}

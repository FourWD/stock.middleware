package orm

import "github.com/FourWD/middleware/orm"

type CompanyBranch struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36); primary_key;"`
	orm.GormModel

	CompanyID string `json:"company_id" gorm:"type:varchar(36);" query:"company_id"`

	Name          string `json:"name" query:"name" gorm:"type:varchar(100);"`
	IsHQ          bool   `json:"is_hq" query:"is_hq" gorm:"type:tinyint(1); column:is_hq;"`
	Address       string `json:"address" query:"address" gorm:"type:varchar(36);"`
	SubdistrictID string `json:"subdistrict_id" query:"subdistrict_id" gorm:"type:varchar(36);"`
	DistrictID    string `json:"district_id" query:"district_id" gorm:"type:varchar(36);"`
	ProvinceID    string `json:"province_id" query:"province_id" gorm:"type:varchar(36);"`
	//address etc.
	// ชื่อ บ. ชื่อ เต้น, ที่อยู่ จดทะเบียน ที่อยู่เต้น
	orm.RowOrder
}

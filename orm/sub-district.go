package orm

import "github.com/FourWD/middleware/orm"

type SubDistrict struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36); primary_key;"`
	orm.GormModel

	DistrictID string `json:"province_id" gorm:"type:varchar(36)"`
	Name       string `json:"name" query:"name" gorm:"type:varchar(100);"`
	Postcode   string `json:"postcode" gorm:"type:varchar(100); column:postcode" query:"postcode"`
	orm.RowOrder
}

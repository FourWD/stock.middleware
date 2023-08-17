package orm

import "github.com/FourWD/middleware/orm"

type District struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36); primary_key;"`
	orm.GormModel

	ProvinceID string `json:"province_id" gorm:"type:varchar(36)"`
	Name       string `json:"name" query:"name" gorm:"type:varchar(100);"`
	orm.RowOrder
}

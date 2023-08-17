package orm

import "github.com/FourWD/middleware/orm"

type PaymentType struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36); primary_key;"`
	orm.GormModel

	Name          string `json:"name" query:"name" gorm:"type:varchar(100);"`
	IsReserveOnly bool   `json:"is_reserve_only" query:"is_reserve_only" gorm:"type:bool(1); " `
	orm.RowOrder
}

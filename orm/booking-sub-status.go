package orm

import "github.com/FourWD/middleware/orm"

type BookingStatusReason struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	BookingStatusID string `json:"booking_status_id" query:"booking_status_id" gorm:"type:varchar(36)" `
	Name            string `json:"name" query:"name" gorm:"type:varchar(100);" `
	orm.RowOrder
}

package orm

import (
	"github.com/FourWD/middleware/orm"
)

type LogUserActivity struct {
	ID string `json:"id" gorm:"type:varchar(36); primary_key;"`
	orm.GormModel

	UserID     string `json:"user_id" query:"user_id" gorm:"type:varchar(50);" `
	UserAction string ` json:"user_action" query:"user_action"  gorm:"type:varchar(100);" `
}

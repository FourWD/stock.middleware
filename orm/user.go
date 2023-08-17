package orm

import (
	"time"

	"github.com/FourWD/middleware/orm"
)

type User struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36); primary_key;"`
	orm.GormModel

	CompanyID string `json:"company_id" query:"company_id" gorm:"type:varchar(36);"`

	Username   string    `json:"username"  query:"username" gorm:"type:varchar(50);"`
	Password   string    `json:"password"  query:"password" gorm:"type:varchar(100);"`
	Firstname  string    `json:"firstname"  query:"firstname" gorm:"type:varchar(50);"`
	Lastname   string    `json:"lastname"  query:"lastname" gorm:"type:varchar(50);"`
	Nickname   string    `json:"nickname"  query:"nickname" gorm:"type:varchar(50);"`
	Birthday   time.Time `json:"birthday" query:"birthday" gorm:"type:datetime;"`
	Avatar     string    `json:"avatar"  query:"avatar" gorm:"type:varchar(100);"`
	Email      string    `json:"email"  query:"email" gorm:"type:varchar(50);"`
	UserTypeID string    `json:"user_type_id"  query:"user_type_id" gorm:"type:varchar(36);"`
}

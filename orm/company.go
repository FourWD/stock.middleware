package orm

import (
	"github.com/FourWD/middleware/orm"
)

type Company struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36); primary_key;"`
	orm.GormModel

	Name             string `json:"name" query:"name" gorm:"type:varchar(100);"`
	RegisterName     string `json:"register_name"  query:"register_name" gorm:"type:varchar(100);"`
	RegisterAddress  string `json:"register_address"  query:"register_address" gorm:"type:varchar(100);"`
	RegisterTaxNo    string `json:"register_tax_no"  query:"register_tax_no" gorm:"type:varchar(100);"`
	RegisterDocument string `json:"register_document"  query:"register_document" gorm:"type:varchar(100);"`
	// register_address , regis_document TaxID com_name
}

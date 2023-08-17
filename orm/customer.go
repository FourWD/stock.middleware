package orm

import (
	"time"

	"github.com/FourWD/middleware/orm"
)

type Customer struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36); primary_key;"`
	orm.GormModel

	GenderID                 string    `json:"gender_id" `
	Firstname                string    `json:"firstname" query:"firstname" gorm:"type:varchar(50);"`
	Lastname                 string    `json:"lastname"  query:"lastname" gorm:"type:varchar(50);"`
	Birthday                 time.Time `json:"birthday" query:"birthday"  gorm:"type:datetime;"`
	CitizenID                string    `json:"citizen_id" query:"citizen_id" gorm:"type:varchar(50);"`
	CareerID                 string    `json:"career_id" query:"career_id" gorm:"type:varchar(50);"`
	CareerSubID              string    `json:"career_sub_id" query:"career_sub_id" gorm:"type:varchar(50);"`
	IncomeID                 string    `json:"income_id" query:"income_id" gorm:"type:varchar(50);"`
	CustomerTypeID           int       `json:"customer_type_id" query:"customer_type_id" gorm:"type:varchar(50);"`
	Tel                      string    `json:"tel" query:"tel" gorm:"type:varchar(50);"`
	TelHome                  string    `json:"tel_home" query:"tel_home" gorm:"type:varchar(50);"`
	Email                    string    `json:"email" query:"email" gorm:"type:varchar(50);"`
	IDCardAddress            string    `json:"id_card_address" query:"id_card_address" gorm:"type:varchar(50);"`
	IDCardSubdistrictID      string    `json:"id_card_subdistrict_id" query:"id_card_subdistrict_id" gorm:"type:varchar(50);"`
	IDCardDistrictID         string    `json:"id_card_district_id" query:"id_card_district_id" gorm:"type:varchar(50);"`
	IDCardProvinceID         string    `json:"id_card_province_id" query:"id_card_province_id" gorm:"type:varchar(50);"`
	IDCardPostcode           string    `json:"id_card_postcode" query:"id_card_postcode" gorm:"type:varchar(50);"`
	IsContactAddressByIDCard int       `json:"is_contact_address_by_id_card" query:"is_contact_address_by_id_card" gorm:"type:varchar(50);"` //ที่อยู่เดียวกันกับบัตรประชาชน
	ContactAddress           string    `json:"contact_address" query:"contact_address" gorm:"type:varchar(50);"`                             //contact address
	ContactSubdistrictID     string    `json:"contact_subdistrict_id" query:"contact_subdistrict_id" gorm:"type:varchar(50);"`
	ContactDistrictID        string    `json:"contact_district_id" query:"contact_district_id" gorm:"type:varchar(50);"`
	ContactProvinceID        string    `json:"contact_province_id" query:"contact_province_id" gorm:"type:varchar(50);"`
	ContactPostcode          string    `json:"contact_postcode" query:"contact_postcode" gorm:"type:varchar(5);"`

	RegisterSource string `json:"register_source" query:"register_source" gorm:"type:varchar(50);"`
	SuggestBy      string `json:"suggest_by" query:"suggest_by" gorm:"type:varchar(50);"`
}

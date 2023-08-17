package orm

import (
	"time"

	"github.com/FourWD/middleware/orm"
)

type Booking struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36); primary_key;"`
	orm.GormModel

	BookingCode string `json:"booking_code" query:"booking_code" gorm:"type:varchar(36)"`

	CompanyID  string `json:"company_id" query:"company_id" gorm:"type:varchar(36)"`
	VehicleID  string `json:"vehicle_id" query:"vehicle_id" gorm:"type:varchar(36)"`
	UserID     string `json:"user_id" query:"user_id" gorm:"type:varchar(36)"`
	CustomerID string `json:"customer_id" query:"customer_id" gorm:"type:varchar(36)"`

	BookingStatusID    string `json:"booking_status" query:"booking_status" gorm:"type:varchar(36)"`
	PoStatusID         string `json:"po_status" query:"po_status" gorm:"type:varchar(36)"`
	BookingSubStatusID string `json:"booking_sub_status_id" query:"booking_sub_status_id" gorm:"type:varchar(36)"`

	Remark string `json:"remark" query:"remark" gorm:"type:varchar(36)"`

	IsDriver   int `json:"is_driver" query:"is_driver" gorm:"type:varchar(36)"` //ผู้ซื้อเป็นคนใช้งานเองหรือไม่
	EvidenceID int `json:"evidence_id" query:"evidence_id" gorm:"type:varchar(36)"`

	PoCareerID      string `json:"po_career_id" query:"po_career_id" gorm:"type:varchar(36)"`
	PoCareerSubID   string `json:"po_career_sub_id" query:"po_career_sub_id" gorm:"type:varchar(36)"`
	PoIncomeID      string `json:"po_income_id" query:"po_income_id" gorm:"type:varchar(36)"`
	PoTel           string `json:"po_tel" query:"po_tel" gorm:"type:varchar(36)"`
	PoTelHome       string `json:"po_tel_home" query:"po_tel_home" gorm:"type:varchar(36)"`
	PoEmail         string `json:"po_email" query:"po_email" gorm:"type:varchar(36)"`
	PoAddress       string `json:"po_address" query:"po_address" gorm:"type:varchar(36)"`
	PoSubdistrictID string `json:"po_subdistrict_id" query:"po_subdistrict_id" gorm:"type:varchar(36)"`
	PoDistrictID    string `json:"po_district_id" query:"po_district_id" gorm:"type:varchar(36)"`
	PoProvinceID    string `json:"po_province_id" query:"po_province_id" gorm:"type:varchar(36)"`
	PoPostcode      string `json:"po_postcode_" query:"po_postcode_" gorm:"type:varchar(36)"`

	// Payment Method
	IsReservePayment     string `json:"is_reserve_payment" query:"is_reserve_payment" gorm:"type:varchar(36)" `
	ReservePaymentID     string `json:"reserve_payment_id" query:"reserve_payment_id" gorm:"type:varchar(36)" `
	ReservePaymentAmount string `json:"reserve_payment_amount" query:"reserve_payment_amount" gorm:"type:varchar(36)" `
	ReservePaymentDate   string `json:"reserve_payment_date" query:"reserve_payment_date" gorm:"type:varchar(36)" `

	DownPaymentID     string    `json:"down_payment_id" query:"down_payment_id" gorm:"type:varchar(36)" `
	DownPaymentAmount float64   `json:"down_payment_amount" query:"down_payment_amount" gorm:"type:float" `
	DownPaymentDate   time.Time `json:"down_payment_date" query:"down_payment_date" gorm:"type:time" `

	FinancePaymentID         string    `json:"finance_payment_id" query:"finance_payment_id" gorm:"type:varchar(36)" `
	FinanceDownPaymentAmount string    `json:"finance_down_payment_amount" query:"finance_down_payment_amount" gorm:"type:varchar(36)" `
	FinanceID                string    `json:"finance_id" query:"finance_id" gorm:"type:varchar(36)" `
	FinanceAcceptDate        string    `json:"finance_accept_date" query:"finance_accept_date" gorm:"type:varchar(36)" `
	FinanceDocumentNo        string    `json:"finance_document_no" query:"finance_document_no" gorm:"type:varchar(36)" `
	FinanceContractor        string    `json:"finance_contractor" query:"finance_contractor" gorm:"type:varchar(36)" ` // คนทำเอกสารกู้
	FinanceContractorTel     string    `json:"finance_contractor_tel" query:"finance_contractor_tel" gorm:"type:varchar(36)" `
	FinanceContractDate      time.Time `json:"finance_contractor_date" query:"finance_contractor_date" gorm:"type:varchar(36)" `
	FinanceLoanTypeID        string    `json:"finance_loan_type_id" query:"finance_loan_type_id" gorm:"type:time" ` //กู้เดียวกู่้ร่วม
	FullDocDate              time.Time `json:"full_doc_date" query:"full_doc_date" gorm:"type:time"`

	PriceFinancing   float64 `json:"price_financing" query:"price_financing" gorm:"type:varchar(36)"`
	FinanceInterrest float64 //  5.8%/48 6.4%/60

	//

	InsuranceCompanyID   string `json:"insurance_name_id" query:"insurance_name_id" gorm:"type:varchar(36)"`
	InsuranceLevelID     string `json:"insurance_level_id" query:"insurance_level_id" gorm:"type:varchar(36)"`
	PriceInsurance       string `json:"price_insurance" query:"price_insurance" gorm:"type:varchar(36)"`
	PriceCommonInsurance string `json:"price_commom_insurance" query:"price_commom_insurance" gorm:"type:varchar(36)"` //ค่าประกันภัยพื้นฐาน
	IsGiftInsurance      int    `json:"is_gift_insurance" query:"is_gift_insurance" gorm:"type:varchar(36)"`

	SubmitInsuranceDate   time.Time `json:"submit_insurance_date" query:"submit_insurance_date" gorm:"type:time"`
	ExpectVehicelDelivery string    `json:"expect_vehicle_delivery" query:"expect_vehicle_delivery" gorm:"type:varchar(50)"`
	ActualVehicelDelivery string    `json:"actual_vehicle_delivery" query:"actual_vehicle_delivery" gorm:"type:varchar(50)"`
}

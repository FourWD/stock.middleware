package orm

import (
	"time"

	"github.com/FourWD/middleware/orm"
)

type Vehicle struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36); primary_key;"`
	orm.GormModel
	VehicleCode string `json:"vehicle_code" query:"vehicle_code" gorm:"type:varchar(36);"`

	VehicleStatusID string `json:"vehicle_status_id" query:"vehicle_status_id" gorm:"type:varchar(36);"`
	BranchID        string `json:"branch_id" query:"branch_id" gorm:"type:varchar(36);"`
	BrandID         string `json:"brand_id" query:"brand_id" gorm:"type:varchar(36);"`
	ModelID         string `json:"model_id" query:"model_id" gorm:"type:varchar(36);"`
	SubmodelID      string `json:"submodel_id" query:"submodel_id" gorm:"type:varchar(36);"`
	VehicleTypeID   string `json:"vehicle_type_id" query:"vehicle_type_id" gorm:"type:varchar(36);"`
	GearTypeID      string `json:"gear_type_id" query:"gear_type_id" gorm:"type:varchar(36);"`
	DriveSystemID   string `json:"drive_system_id" query:"drive_system_id" gorm:"type:varchar(36);"`
	ColorID         string `json:"color_id" query:"color_id" gorm:"type:varchar(36);"`
	FuelTypeID      string `json:"fuel_type_id" query:"fuel_type_id" gorm:"type:varchar(36);"`

	ChassisNumber string `json:"chassis_number" query:"chassis_number" gorm:"type:varchar(50);"`
	EngineNumber  string `json:"engine_number" query:"engine_number" gorm:"type:varchar(50);"`

	YearBu     string `json:"year_bu" query:"year_bu" gorm:"type:varchar(4);"`
	YearReg    string `json:"year_reg" query:"year_reg" gorm:"type:varchar(4);"`
	EngineSize string `json:"engine_size" query:"engine_size" gorm:"type:varchar(50);"`

	Seat                  int       `json:"seat" query:"seat" gorm:"type:varchar(4);"`
	Mileage               int       `json:"mileage" query:"mileage" gorm:"type:varchar(4);"`
	LicensePlate          string    `json:"license_plate" query:"license_plate" gorm:"type:varchar(10);"`
	LicenseProvinceID     string    `json:"license_province_id" query:"license_province_id" gorm:"type:varchar(36);"` // Join Province
	LicenseRegisteredDate time.Time `json:"license_registered_date" query:"license_registered_date" gorm:"type:time;"`
	LicenseExpired        time.Time `json:"license_expired" query:"license_expired" gorm:"type:time;"`
	InsuranceExpired      time.Time `json:"insurance_expired" query:"insurance_expired" gorm:"type:time;"`

	Price    float64 `json:"price" query:"price" gorm:"type:float;"`
	Vat      float64 `json:"vat" query:"vat" gorm:"type:float;"`
	NetPrice float64 `json:"net_price" query:"net_price" gorm:"type:float;"`

	Remark string `json:"remark" query:"remark" gorm:"type:varchar(200);"`

	ImgStrFront             string `json:"img_str_front" query:"img_str_front" gorm:"type:varchar(500);"`
	ImgStrBack              string `json:"img_str_back" query:"img_str_back" gorm:"type:varchar(500);"`
	ImgStrRight             string `json:"img_str_right" query:"img_str_right" gorm:"type:varchar(500);"`
	ImgStrLeft              string `json:"img_str_left" query:"img_str_left" gorm:"type:varchar(500);"`
	ImgFrontLeft45          string `json:"img_front_left_45" query:"img_front_left_45" gorm:"type:varchar(500);"`
	ImgFrontRight45         string `json:"img_front_right_45" query:"img_front_right_45" gorm:"type:varchar(500);"`
	ImgBackLeft45           string `json:"img_back_left_45" query:"img_back_left_45" gorm:"type:varchar(500);"`
	ImgBackRight45          string `json:"img_back_right_45" query:"img_back_right_45" gorm:"type:varchar(500);"`
	ImgInFront              string `json:"img_in_front" query:"img_in_front" gorm:"type:varchar(500);"`
	ImgInBack               string `json:"img_in_back" query:"img_in_back" gorm:"type:varchar(500);"`
	ImgConsole              string `json:"img_console" query:"img_console" gorm:"type:varchar(500);"`
	ImgMileage              string `json:"img_mileage" query:"img_mileage" gorm:"type:varchar(500);"`
	ImgVehTools             string `json:"img_veh_tools" query:"img_veh_tools" gorm:"type:varchar(500);"`
	ImgEngineRoom           string `json:"img_engine_room" query:"img_engine_room" gorm:"type:varchar(500);"`
	ImgGas                  string `json:"img_gas" query:"img_gas" gorm:"type:varchar(500);"`
	ImgOut360               string `json:"img_out_360" query:"img_out_360" gorm:"type:varchar(500);"`
	ImgIn360                string `json:"img_in_360" query:"img_in_360" gorm:"type:varchar(500);"`
	ImgAct                  string `json:"img_act" query:"img_act" gorm:"type:varchar(500);"`
	ImgInsurance            string `json:"img_insurance" query:"img_insurance" gorm:"type:varchar(500);"`
	ImgInspectionFront      string `json:"img_inspection_front" query:"img_inspection_front" gorm:"type:varchar(500);"`
	ImgInspectionBack       string `json:"img_inspection_back" query:"img_inspection_back" gorm:"type:varchar(500);"`
	ImgStrFrontThumb        string `json:"img_str_front_thumb" query:"img_str_front_thumb" gorm:"type:varchar(500);"`
	ImgStrBackThumb         string `json:"img_str_back_thumb" query:"img_str_back_thumb" gorm:"type:varchar(500);"`
	ImgStrRightThumb        string `json:"img_str_right_thumb" query:"img_str_right_thumb" gorm:"type:varchar(500);"`
	ImgStrLeftThumb         string `json:"img_str_left_thumb" query:"img_str_left_thumb" gorm:"type:varchar(500);"`
	ImgFrontLeft45Thumb     string `json:"img_front_left_45_thumb" query:"img_front_left_45_thumb" gorm:"type:varchar(500);"`
	ImgFrontRight45Thumb    string `json:"img_front_right_45_thumb" query:"img_front_right_45_thumb" gorm:"type:varchar(500);"`
	ImgBackLeft45Thumb      string `json:"img_back_left_45_thumb" query:"img_back_left_45_thumb" gorm:"type:varchar(500);"`
	ImgBackRight45Thumb     string `json:"img_back_right_45_thumb" query:"img_back_right_45_thumb" gorm:"type:varchar(500);"`
	ImgInFrontThumb         string `json:"img_in_front_thumb" query:"img_in_front_thumb" gorm:"type:varchar(500);"`
	ImgInBackThumb          string `json:"img_in_back_thumb" query:"img_in_back_thumb" gorm:"type:varchar(500);"`
	ImgConsoleThumb         string `json:"img_console_thumb" query:"img_console_thumb" gorm:"type:varchar(500);"`
	ImgMileageThumb         string `json:"img_mileage_thumb" query:"img_mileage_thumb" gorm:"type:varchar(500);"`
	ImgVehToolsThumb        string `json:"img_veh_tools_thumb" query:"img_veh_tools_thumb" gorm:"type:varchar(500);"`
	ImgSpareWheel           string `json:"img_spare_wheel" query:"img_spare_wheel" gorm:"type:varchar(500);"`
	ImgEngineRoomThumb      string `json:"img_engine_room_thumb" query:"img_engine_room_thumb" gorm:"type:varchar(500);"`
	ImgGasThumb             string `json:"img_gas_thumb" query:"img_gas_thumb" gorm:"type:varchar(500);"`
	ImgOut360Thumb          string `json:"img_out_360_thumb" query:"img_out_360_thumb" gorm:"type:varchar(500);"`
	ImgIn360Thumb           string `json:"img_in_360_thumb" query:"img_in_360_thumb" gorm:"type:varchar(500);"`
	ImgActThumb             string `json:"img_act_thumb" query:"img_act_thumb" gorm:"type:varchar(500);"`
	ImgInsuranceThumb       string `json:"img_insurance_thumb" query:"img_insurance_thumb" gorm:"type:varchar(500);"`
	ImgInspectionFrontThumb string `json:"img_inspection_front_thumb" query:"img_inspection_front_thumb" gorm:"type:varchar(500);"`
	ImgInspectionBackThumb  string `json:"img_inspection_back_thumb" query:"img_inspection_back_thumb" gorm:"type:varchar(500);"`
}

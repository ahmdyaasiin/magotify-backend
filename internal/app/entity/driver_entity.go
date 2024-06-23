package entity

type Driver struct {
	//
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"-"`
	VehicleID   string `json:"vehicle_id"`
	WarehouseID string `json:"warehouse_id"`
	PlateNumber string `json:"plate_number"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

func (e Driver) GetTableName() string {
	return "drivers"
}

package models

type Customers struct {
	ID      uint64     `gorm:"primaryKey" json:"id"`
	Name    string     `json:"name" binding:"required"`
	Nid     uint64     `json:"nid" binding:"required"`
	Gender  string     `json:"gender" binding:"required"`
	Phone   string     `json:"phone_number" binding:"required"`
	Address string     `json:"address" binding:"required"`
	Job     string     `json:"job" binding:"required"`
	Salary  string     `json:"salary" binding:"required"`
	Account []Accounts `gorm:"many2one:account_customers" json:"type_account,omitempty"`
}

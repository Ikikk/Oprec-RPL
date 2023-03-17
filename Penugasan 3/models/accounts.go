package models

type Accounts struct {
	ID         uint64     `gorm:"primaryKey" json:"id"`
	Type       string     `json:"account_type" binding:"required"`
	Balance    uint64     `json:"balance" binding:"required"`
	Atm_Number string     `json:"ATM_number" binding:"required"`
	CVV        int        `json:"CVV_number" binding:"required"`
	Cust_id    uint64     `gorm:"foreignkey" json:"cust_id" binding:"required"`
	Cust       *Customers `gorm:"constraint:OnUpdate:CASCADE;" json:"user,omitempty"`
}

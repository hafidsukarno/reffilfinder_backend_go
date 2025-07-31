package models

import "gorm.io/gorm"

type Transaksi struct {
	gorm.Model
	UserID        uint   `gorm:"not null" json:"user_id"`
	PaymentMethod string `gorm:"not null" json:"payment_method"`
	Address       string `gorm:"not null" json:"address"`
	TransaksiDate string `gorm:"not null" json:"transaksi_date"`

	User               User               `gorm:"foreignKey:UserID"`
	DetailTransaksis   []DetailTransaksi `gorm:"foreignKey:TransaksiID"`
}

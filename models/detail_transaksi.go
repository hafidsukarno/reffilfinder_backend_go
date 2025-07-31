package models

import "gorm.io/gorm"

type DetailTransaksi struct {
	gorm.Model
	TransaksiID uint   `gorm:"not null" json:"transaksi_id"`
	ProductID   uint   `gorm:"not null" json:"product_id"`
	Quantity    int    `gorm:"not null" json:"quantity"`
	Status      string `gorm:"not null" json:"status"`

	Transaksi   Transaksi `gorm:"foreignKey:TransaksiID"`
	Product     Product   `gorm:"foreignKey:ProductID"`
}

package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"not null;unique" json:"username"`
	Email    string `gorm:"not null;unique" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Role     string `gorm:"not null" json:"role"` // "penjual" atau "pembeli"

	Products   []Product   `gorm:"foreignKey:SellerID"`
	Promotions []Promotion `gorm:"foreignKey:SellerID"`
	Carts      []Cart      `gorm:"foreignKey:UserID"`
	Transaksis []Transaksi `gorm:"foreignKey:UserID"`
}

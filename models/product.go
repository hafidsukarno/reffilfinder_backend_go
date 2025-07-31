package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ProductName        string  `gorm:"not null" json:"product_name"`
	ProductDescription string  `gorm:"not null" json:"product_description"`
	ProductPrice       float64 `gorm:"not null" json:"product_price"`
	ProductStock       int     `gorm:"not null" json:"product_stock"`
	ProductImage       string  `json:"product_image"`
	SellerID           uint    `gorm:"not null" json:"seller_id"`

	Seller             User               `gorm:"foreignKey:SellerID"`
	DetailTransaksis   []DetailTransaksi `gorm:"foreignKey:ProductID"`
}

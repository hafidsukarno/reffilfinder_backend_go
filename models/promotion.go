package models

import "gorm.io/gorm"

type Promotion struct {
	gorm.Model
	PromoName  string `gorm:"not null" json:"promo_name"`
	PromoImage string `json:"promo_image"`
	SellerID   uint   `gorm:"not null" json:"seller_id"`

	Seller     User   `gorm:"foreignKey:SellerID"`
}

package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserID    uint    `gorm:"not null" json:"user_id"`
	ProductID uint    `gorm:"not null" json:"product_id"`
	Quantity  int     `gorm:"not null" json:"quantity"`

	User      User    `gorm:"foreignKey:UserID"`
	Product   Product `gorm:"foreignKey:ProductID"`
}

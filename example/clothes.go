package models

import "gorm.io/gorm"

// PackingClothes
// @Gorm
type PackingClothes struct {
	gorm.Model
	PackingID int  `gorm:"type:integer;not null;uniqueIndex:packing_clothes_pid_cid_uidx,priority:1"`
	ClothesID int  `gorm:"type:integer;not null;uniqueIndex:packing_clothes_pid_cid_uidx,priority:2"`
	IsChecked bool `gorm:"default:false;not null"`

	Clothes Clothes `gorm:"foreignKey:ClothesID;references:ID"`
}

// Clothes
// @Gorm
type Clothes struct {
	gorm.Model
	Name        string `gorm:"type:varchar(100);not null"`
	Description string `gorm:"type:text"`
}

package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Product struct {
	ID        	uint   `gorm:"primaryKey"`
	UUID        string `gorm:"not null"`
	Name     		string `gorm:"not null" json:"name" form:"name" valid:"required~Your product name is required"`
	ImageUrl    string `gorm:"not null"`
	AdminId 		uint 
	Admin     	*Admin
	Variant   	[]Variant `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"variants"`
	CreatedAt 	*time.Time
	UpdatedAt 	*time.Time
}


func (p Product) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (p Product) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Variant struct {
	ID        	uint   		`gorm:"primaryKey"`
	UUID        string   	`gorm:"not null"`
	VariantName	string 		`gorm:"not null;unique" json:"variant_name" form:"title" valid:"required~Name of varient is required"`
	Quantity    uint			`gorm:"not null" json:"quantity" form:"quantity" valid:"required~Quantity of variant is required, numeric~Invalid quantity format"`
	ProductId		uint 			
	Product     *Product	
	CreatedAt 	*time.Time 
	UpdatedAt 	*time.Time 
}

func (v *Variant) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(v)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (v *Variant) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(v)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

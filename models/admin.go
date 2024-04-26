package models

import (
	"time"
	"trade-api/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Admin struct {
	ID        	uint   `gorm:"primaryKey"`
	UUID        string `gorm:"not null"` 
	Name     		string `gorm:"not null;type:varchar(191)" json:"name" form:"name" valid:"required~Name is required"`
	Email     	string `gorm:"not null;unique;type:varchar(191)" json:"email" form:"email" valid:"required~Email is required"`
	Password 		string `gorm:"not null;type:varchar(191)" json:"password" form:"password" valid:"required~Password is required"`
	Product   	[]Product  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"products"`
	CreatedAt 	time.Time 
	UpdatedAt 	time.Time 
}


func (a *Admin) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(a)

	if errCreate != nil {
		err = errCreate
		return
	}

	a.Password = helpers.HashPass(a.Password)

	err = nil
	return
}

func (a *Admin) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(a)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

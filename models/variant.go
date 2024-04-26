package models

import "time"

type Variant struct {
	ID        	uint   		`gorm:"primaryKey"`
	UUID        string   	`gorm:"not null"`
	VariantName	string 		`gorm:"not null;unique" json:"variant_name" form:"title" valid:"required~Name of varient is required"`
	Quantity    uint			`gorm:"not null" json:"quantity" form:"quantity" valid:"required~Quantity of variant is required, numeric~Invalid quantity format"`
	ProductId		uint 			`gorm:"not null"`
	Product     *Product
	CreatedAt 	*time.Time 
	UpdatedAt 	*time.Time 
}

// func (p *Variant) BeforeCreate(tx *gorm.DB) error {
// 	p.UUID = uuid.New()
// 	return nil
// }
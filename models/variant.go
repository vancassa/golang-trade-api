package models

import "time"

type Variant struct {
	ID        	uint   		`gorm:"primaryKey"`
	UUID        string   
	VariantName	string 		`gorm:"not null;unique;type:varchar(191)"`
	Quantity    uint			`gorm:"not null"`
	ProductId		uint 			`gorm:"not null"`
	CreatedAt 	time.Time 
	UpdatedAt 	time.Time 
}

// func (p *Variant) BeforeCreate(tx *gorm.DB) error {
// 	p.UUID = uuid.New()
// 	return nil
// }
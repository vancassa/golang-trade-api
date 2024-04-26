package models

import "time"

type Product struct {
	ID        	uint   `gorm:"primaryKey"`
	UUID        string   
	Name     		string `gorm:"not null;unique;type:varchar(191)"`
	File    		string `gorm:"not null;type:varchar(191)"`
	AdminId 		uint 
	CreatedAt 	time.Time 
	UpdatedAt 	time.Time 
}

// func (p *Product) BeforeCreate(tx *gorm.DB) error {
// 	p.UUID = uuid.New()
// 	return nil
// }
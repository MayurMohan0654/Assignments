package models

import "time"

type Orders struct {
	ID           string     `json:"order_id" gorm:"primaryKey" binding:"required"`
	FacilityCode string     `json:"facility_code" binding:"required"`
	Status       string     `json:"status" binding:"required"`
	CreatedAt    time.Time  `json:"created_at" gorm:"autoCreateTime"`
	Facility     Facilities `gorm:"foreignKey:FacilityCode;references:Code" binding:"required"`
}

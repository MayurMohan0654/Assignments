package models

import "time"

type Orders struct {
	ID           string     `json:"order_id" gorm:"primaryKey"`
	FacilityCode string     `json:"facility_code"`
	Status       string     `json:"status"`
	CreatedAt    time.Time  `json:"created_at" gorm:"autoCreateTime"`
	Facility     Facilities `gorm:"foreignKey:FacilityCode;references:Code"`
}

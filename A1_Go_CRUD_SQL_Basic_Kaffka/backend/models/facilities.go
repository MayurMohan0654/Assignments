package models

import (
	"time"
)

type Facilities struct {
	Code      string    `json:"code" gorm:"primaryKey; type:varchar(10)" binding:"required"`
	Name      string    `json:"name" binding:"required"`
	Address   string    `json:"address" binding:"required"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}

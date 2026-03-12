package models

import (
	"time"
)

type Facilities struct {
	Code      string    `json:"code" gorm:"primaryKey; type:varchar(10)"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}

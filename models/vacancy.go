package models

import (
	"time"

	"gorm.io/gorm"
)

type Vacancy struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	PlaceName   string         `json:"place_name"`
	Description string         `json:"description"`
	Payment     float32        `json:"payment"`
	Author      string         `json:"author"`
	ImageURL    string         `json:"img_url"`
	CreatedAt   time.Time      `json:"created"`
	UpdatedAt   time.Time      `json:"updated"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted"`
}

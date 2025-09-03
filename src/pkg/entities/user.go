package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Email    string    `json:"email" gorm:"unique;not null"`
	Password string    `json:"password" gorm:"not null"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
	WeightUnit string         `json:"weightUnit" gorm:"type:varchar(5);not null;check:weight_unit IN ('KG','LBS')"`
	HeightUnit string         `json:"heightUnit" gorm:"type:varchar(10);not null;check:height_unit IN ('CM','INCH')"`
	Weight     float64        `json:"weight" gorm:"not null;check:weight >= 10 AND weight <= 1000"`
	Height     float64        `json:"height" gorm:"not null;check:height >= 3 AND height <= 250"`
	Name       string         `json:"name" gorm:"type:varchar(60)"`
	ImageURI   string         `json:"imageUri" gorm:"type:text"`
	
}

type UserDeleteRequest struct {
	ID uuid.UUID `json:"id"`
}

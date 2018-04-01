package model

import "time"

type Meigen struct {
	ID          int
	Message     string
	CharacterID int `gorm:"index"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

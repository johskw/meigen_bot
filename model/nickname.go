package model

import "time"

type Nickname struct {
	ID          int
	Nickname    string `gorm:"index"`
	CharacterID int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

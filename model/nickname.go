package model

import "time"

type Nickname struct {
	ID          int
	Nickname    string `form:"nickname" gorm:"index"`
	CharacterID int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (nickname Nickname) Create() (Nickname, error) {
	err := db.Create(&nickname).Error
	return nickname, err
}

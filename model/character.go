package model

import (
	"time"
)

type Character struct {
	ID        int
	Name      string
	Nicknames []Nickname
	Meigens   []Meigen
	CreatedAt time.Time
	UpdatedAt time.Time
}

func GetAllCharacters() (characters []Character, err error) {
	err = db.Find(&characters).Error
	return
}

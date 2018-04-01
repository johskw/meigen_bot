package model

import (
	"time"
)

type Character struct {
	ID        int
	Name      string `form:"name"`
	Nicknames []Nickname
	Meigens   []Meigen
	CreatedAt time.Time
	UpdatedAt time.Time
}

func GetAllCharacters() (characters []Character, err error) {
	err = db.Find(&characters).Error
	return
}

func (character Character) Create() (Character, error) {
	err := db.Create(&character).Error
	return character, err
}

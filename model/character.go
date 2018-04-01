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

func GetCharacter(id int) (character Character, err error) {
	err = db.First(&character, id).Related(&character.Nicknames).Related(&character.Meigens).Error
	return
}

func (character Character) Create() (Character, error) {
	err := db.Create(&character).Error
	return character, err
}

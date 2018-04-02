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

func (character Character) Update() (Character, error) {
	err := db.Save(&character).Error
	return character, err
}

func DeleteCharacter(id int) (err error) {
	var character Character
	err = db.First(&character, id).Related(&character.Nicknames).Related(&character.Meigens).Error
	if err != nil {
		return
	}
	tx := db.Begin()
	err = tx.Delete(&character).Error
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Delete(Nickname{}, "character_id = ?", character.ID).Error
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Delete(Meigen{}, "character_id = ?", character.ID).Error
	if err != nil {
		tx.Rollback()
		return
	}
	return tx.Commit().Error
}

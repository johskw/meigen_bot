package model

import "time"

type Meigen struct {
	ID          int
	Text        string `form:"text"`
	CharacterID int    `gorm:"index"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (meigen Meigen) Create() (Meigen, error) {
	err := db.Create(&meigen).Error
	return meigen, err
}

func DeteleMeigen(id int) (err error) {
	var meigen Meigen
	err = db.First(&meigen, id).Error
	if err != nil {
		return
	}
	err = db.Delete(&meigen).Error
	return
}

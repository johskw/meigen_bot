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

func (nickname Nickname) Update() (Nickname, error) {
	err := db.Save(&nickname).Error
	return nickname, err
}

func DeteleNickname(id int) (err error) {
	var nickname Nickname
	err = db.First(&nickname, id).Error
	if err != nil {
		return
	}
	err = db.Delete(&nickname).Error
	return
}

func GetNickname(text string) (nickname Nickname, err error) {
	err = db.Where("nickname = ?", text).First(&nickname).Error
	return
}

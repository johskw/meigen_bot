package model

import (
	"math/rand"
	"time"
)

func GetMeigenFromNickname(text string) (meigenText string, err error) {
	nickname, err := GetNickname(text)
	if err != nil {
		meigenText = "そのキャラクターにはまだ対応してないんだ...ゴメンね(> <)"
		return
	}
	meigens, err := GetMeigens(nickname.CharacterID)
	if err != nil {
		return
	} else if len(meigens) == 0 {
		meigenText = "そのキャラクターにはまだ対応してないんだ...ゴメンね(> <)"
		return
	}
	rand.Seed(time.Now().UnixNano())
	meigenText = meigens[rand.Intn(len(meigens))].Text
	return
}

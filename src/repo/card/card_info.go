package card

import "sunny.ksw.kr/repo"

type Card_Info struct {
	repo.MongoBase `bson:",inline"`

	CardAdId int `json:"cardAdId" bson:"cardAdId"`

	CardName string `json:"cardName" bson:"cardName"`

	Benefit []Card_Benefit `json:"benefit" bson:"benefit"`
}

type Card_Benefit struct {
	Title string `json:"title" bson:"title"`

	Content string `json:"content" bson:"content"`
}

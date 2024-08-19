package card_test

import (
	"strings"
	"testing"

	"sunny.ksw.kr/inits"
	"sunny.ksw.kr/repo/card"
)

func TestCard(t *testing.T) {
	// .
	mongo_uri := "mongodb://localhost:27017/?maxPoolSize=20&w=majority"
	mongo_db := "local"
	// mongo_uri := "mongodb://localhost:27017"
	// mongo_db := "test" //

	inits.MongoInitDebug(mongo_uri, mongo_db)

	search := card.SearchCard{}
	search.Finds()

	for i, card := range search.Cards {
		for j, discount := range card.Max_discount {
			search.Cards[i].Max_discount[j].Amount = strings.TrimSpace(discount.Amount)
			search.Cards[i].Max_discount[j].Price = strings.TrimSpace(discount.Price)

			card.Update()
		}
	}
}

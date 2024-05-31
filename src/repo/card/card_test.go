package card_test

import (
	"testing"
	"time"

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

	for _, v := range search.Cards {
		v.CreatedTime = time.Now()
		v.Update()
	}
}

package comn

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/**

  search 사용법

  search := repopkg.ModelSearch{}

  search.Filter()

  search.SetLimitPageOffset(10, 0)

  search.SelectField()


  search.Find()



*/

type Search struct {

	// 페이징 처리 여부를 간단하게 체크한다.
	IsPage bool
	//.
	PageOffset int `json:"page_offset" default:"0"`

	Page int `json:"page" default:"1"`

	Limit int `json:"limit" default:"15"`

	//Able bool `default:"true"`

	Total int64 `json:"total"`

	SortField     string `json:"sort_field"`
	SortDirection int    `json:"sort_direction"`

	// .
	Fields *options.FindOptions `json:"-"`

	// .
	Filter bson.M `json:"-"`

	//
	OffLimit *options.FindOptions `json:"-"`

	CollectionName string `json:"-"`
}

// .
func (search *Search) SetLimitPageOffset(limit int, pageoff int) {

	if limit > 0 && pageoff > -1 {

		search.Limit = limit
		search.PageOffset = pageoff

		search.OffLimit = options.Find().SetSkip(int64(search.Limit) * int64(search.PageOffset)).SetLimit(int64(search.Limit))
		search.IsPage = true

	} else {

		search.IsPage = false
	}

}

func (search *Search) SetFilter(filter bson.M) {

	filter["able"] = true

	search.Filter = filter

}

// .
func (search *Search) SetFields(fields bson.D) {

	search.Fields = options.Find().SetProjection(fields)

}

package card

import (
	"context"
	"log"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sunny.ksw.kr/co"
	"sunny.ksw.kr/comn"
	"sunny.ksw.kr/inits"
	"sunny.ksw.kr/repo"
)

type Card_Info struct {
	repo.MongoBase `bson:",inline"`

	CardAdId string `json:"cardAdId" bson:"cardAdId"`

	CardName string `json:"cardName" bson:"cardName"`

	Benefit []Card_Benefit `json:"benefit" bson:"benefit"`

	Info_detail []Info `json:"info_detail" bson:"info_detail"`
}

type Card_Benefit struct {
	Title string `json:"title" bson:"title"`

	Content string `json:"content" bson:"content"`
}

type Info struct {
	Title string `json:"title" bson:"title"`
	Info  string `json:"info" bson:"info"`
}

func Card_InfoCollectionName() string {
	return "card_info"
}

func (model *Card_Info) CollectionName() string {
	//
	return Card_InfoCollectionName()
}

/* ****************************************************************************
  Basic CRUD

***************************************************************************** */

// .
func (model *Card_Info) GetById(id string) (errEx co.MsgEx) {

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return co.ErrorPass(err.Error())
	}

	err = inits.MongoDb.Collection(model.CollectionName()).FindOne(context.TODO(), bson.D{{Key: "_id", Value: objectId}}).Decode(&model)
	if err != nil {
		return co.ErrorPass(err.Error())
	}

	return co.SuccessPass("")

}

// .
func (model *Card_Info) Create() (errEx co.MsgEx) {

	model.Able = true
	model.CreatedTime = time.Now()

	_, err := inits.MongoDb.Collection(model.CollectionName()).InsertOne(context.TODO(), model)
	if err != nil {
		log.Print("  >> err : ", err.Error())
		return co.ErrorPass(err.Error())
	}

	return co.SuccessPass("")
}

// .
func (model *Card_Info) Delete() (errEx co.MsgEx) {

	model.Able = false

	_, err := inits.MongoDb.Collection(model.CollectionName()).UpdateOne(context.TODO(), bson.D{{Key: "_id", Value: model.ID}}, bson.M{"$set": model})
	if err != nil {
		return co.ErrorPass(err.Error())
	}
	return co.SuccessPass("")
}

// .
func (model *Card_Info) Update() (errEx co.MsgEx) {

	model.Able = true
	//update := bson.D{}
	model.UpdatedTime = time.Now()

	_, err := inits.MongoDb.Collection(model.CollectionName()).UpdateOne(context.TODO(), bson.D{{Key: "_id", Value: model.ID}}, bson.M{"$set": model})
	if err != nil {
		return co.ErrorPass(err.Error())
	}

	return co.SuccessPass("")
}

func (model *Card_Info) GetList(page string, limit string) (result []*Card_Info, errEx co.MsgEx) {
	pageInt, _ := strconv.Atoi(page)
	limitInt, _ := strconv.Atoi(limit)

	skip := 0
	if pageInt > 0 {
		skip = (pageInt - 1) * limitInt
	}
	cursor, err := inits.MongoDb.Collection(model.CollectionName()).Find(context.TODO(), bson.D{{Key: "able", Value: true}}, options.Find().SetSkip(int64(skip)).SetLimit(int64(limitInt)).SetSort(bson.M{"createdtime": -1}))
	if err != nil {
		return nil, co.ErrorPass(err.Error())
	} else {
		cursor.All(context.TODO(), &result)
	}
	return result, co.SuccessPass("")
}

/* ****************************************************************************
  Find
***************************************************************************** */
// .
func FindCard_InfoById(id string) (model Card_Info, errMsg co.MsgEx) {

	//err = mgm.Coll(&L01201{}).FindByID(id, &model)

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model, co.ErrorPass(err.Error())
	}
	err = inits.MongoDb.Collection(Card_InfoCollectionName()).FindOne(context.TODO(), bson.M{"_id": objectId}).Decode(&model)
	if err != nil {
		return model, co.ErrorPass(err.Error())
	}

	return model, co.SuccessPass("")
}

func FindCard_InfoByCardID(card_id string) (model Card_Info, errMsg co.MsgEx) {

	//err = mgm.Coll(&L01201{}).FindByID(id, &model)

	err := inits.MongoDb.Collection(Card_InfoCollectionName()).FindOne(context.TODO(), bson.M{"cardAdId": card_id}).Decode(&model)
	if err != nil {
		return model, co.ErrorPass(err.Error())
	}

	return model, co.SuccessPass("")
}

// .

/* ***********************************************************************
* search
 * *********************************************************************** */

// .

type SearchCard_Info struct {
	//
	comn.Search

	Card_Infos []*Card_Info `json:"card_infos" `
}

func (search *SearchCard_Info) CollectionName() string {
	//
	return Card_InfoCollectionName()
}
func (search *SearchCard_Info) condition() bson.M {

	filter := bson.M{}

	// filter["able"] = true

	return filter

}

// .
func (search *SearchCard_Info) Finds() (errEx co.MsgEx) {
	sort := bson.M{}
	if co.NotEmptyString(search.SortField) {
		if search.SortDirection != 1 {
			search.SortDirection = -1
		} else {
			search.SortDirection = 1
		}
		sort[search.SortField] = search.SortDirection
	} else {
		sort["createdtime"] = -1
	}

	if search.Limit > 0 && search.PageOffset > -1 {

		cursor, err := inits.MongoDb.Collection(search.CollectionName()).Find(context.TODO(), search.condition(),
			options.Find().SetSkip(int64(search.Limit)*int64(search.PageOffset)).SetLimit(int64(search.Limit)).SetSort(sort))
		if err != nil {
			return co.ErrorPass(err.Error())
		}

		if err = cursor.All(context.TODO(), &search.Card_Infos); err != nil {
			return co.ErrorPass(err.Error())
		}

	} else {

		cursor, err := inits.MongoDb.Collection(search.CollectionName()).Find(context.TODO(), search.condition(), options.Find().SetSort(sort))
		if err != nil {
			return co.ErrorPass(err.Error())
		}

		if err = cursor.All(context.TODO(), &search.Card_Infos); err != nil {
			return co.ErrorPass(err.Error())
		}
	}

	total, err := inits.MongoDb.Collection(search.CollectionName()).CountDocuments(context.TODO(), search.condition())
	if err != nil {
		return co.ErrorPass(err.Error())
	}

	search.Total = total

	return errEx

}

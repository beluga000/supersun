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

type Card struct {
	repo.MongoBase `bson:",inline"`

	CardAdId               string    `json:"cardAdId" bson:"cardAdId"`
	CardName               string    `json:"cardName" bson:"cardName"`
	CompanyCode            string    `json:"companyCode" bson:"companyCode"`
	TitleDescription       string    `json:"titleDescription" bson:"titleDescription"`
	CardImage              string    `json:"cardImage" bson:"cardImage"`
	CardImageUrl           string    `json:"cardImageUrl" bson:"cardImageUrl"`
	RegisterUrl            string    `json:"registerUrl" bson:"registerUrl"`
	RegisterUrlForNoCharge string    `json:"registerUrlForNoCharge" bson:"registerUrlForNoCharge"`
	DomesticAnnualFee      int       `json:"domesticAnnualFee" bson:"domesticAnnualFee"`
	ForeignAnnualFee       int       `json:"foreignAnnualFee" bson:"foreignAnnualFee"`
	EnableNpayMO           bool      `json:"enableNpayMO" bson:"enableNpayMO"`
	EnableNpayPC           bool      `json:"enableNpayPC" bson:"enableNpayPC"`
	ImpBeacon              string    `json:"impBeacon" bson:"impBeacon"`
	Benefits               []Benefit `json:"benefits" bson:"benefits"`
	ReleaseAt              string    `json:"releaseAt" bson:"releaseAt"`
	BizType                string    `json:"bizType" bson:"bizType"`
	IsMinCPC               bool      `json:"isMinCPC" bson:"isMinCPC"`
}

type Benefit struct {
	Order                     int    `json:"order" bson:"order"`
	RootBenefitCategoryIdName string `json:"rootBenefitCategoryIdName" bson:"rootBenefitCategoryIdName"`
	IconFileName              string `json:"iconFileName" bson:"iconFileName"`
	IconFileNameUrl           string `json:"iconFileNameUrl" bson:"iconFileNameUrl"`
}

func CardCollectionName() string {
	return "card"
}

func (model *Card) CollectionName() string {
	//
	return CardCollectionName()
}

/* ****************************************************************************
  Basic CRUD

***************************************************************************** */

// .
func (model *Card) GetById(id string) (errEx co.MsgEx) {

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
func (model *Card) Create() (errEx co.MsgEx) {

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
func (model *Card) Delete() (errEx co.MsgEx) {

	model.Able = false

	_, err := inits.MongoDb.Collection(model.CollectionName()).UpdateOne(context.TODO(), bson.D{{Key: "_id", Value: model.ID}}, bson.M{"$set": model})
	if err != nil {
		return co.ErrorPass(err.Error())
	}
	return co.SuccessPass("")
}

// .
func (model *Card) Update() (errEx co.MsgEx) {

	model.Able = true
	//update := bson.D{}
	model.UpdatedTime = time.Now()

	_, err := inits.MongoDb.Collection(model.CollectionName()).UpdateOne(context.TODO(), bson.D{{Key: "_id", Value: model.ID}}, bson.M{"$set": model})
	if err != nil {
		return co.ErrorPass(err.Error())
	}

	return co.SuccessPass("")
}

func (model *Card) GetList(page string, limit string) (result []*Card, errEx co.MsgEx) {
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
func FindCardById(id string) (model Card, errMsg co.MsgEx) {

	//err = mgm.Coll(&L01201{}).FindByID(id, &model)

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model, co.ErrorPass(err.Error())
	}
	err = inits.MongoDb.Collection(CardCollectionName()).FindOne(context.TODO(), bson.M{"_id": objectId}).Decode(&model)
	if err != nil {
		return model, co.ErrorPass(err.Error())
	}

	return model, co.SuccessPass("")
}

func FindDataRequestByIdPlusCount(id string) (model Card, errMsg co.MsgEx) {

	//err = mgm.Coll(&L01201{}).FindByID(id, &model)

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model, co.ErrorPass(err.Error())
	}
	err = inits.MongoDb.Collection(CardCollectionName()).FindOne(context.TODO(), bson.M{"_id": objectId}).Decode(&model)
	if err != nil {
		return model, co.ErrorPass(err.Error())
	}

	// model.ViewCount = model.ViewCount + 1
	model.Update()

	return model, co.SuccessPass("")
}

// .

/* ***********************************************************************
* search
 * *********************************************************************** */

// .
type SearchCard struct {
	//
	comn.Search

	//

	Cards []*Card
}

func (search *SearchCard) CollectionName() string {
	//
	return CardCollectionName()
}

// .
func (search *SearchCard) condition() bson.M {

	filter := bson.M{}

	return filter

}

// .
func (search *SearchCard) Finds() (errEx co.MsgEx) {
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

		if err = cursor.All(context.TODO(), &search.Cards); err != nil {
			return co.ErrorPass(err.Error())
		}

	} else {

		cursor, err := inits.MongoDb.Collection(search.CollectionName()).Find(context.TODO(), search.condition(), options.Find().SetSort(sort))
		if err != nil {
			return co.ErrorPass(err.Error())
		}

		if err = cursor.All(context.TODO(), &search.Cards); err != nil {
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

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
	// 카드ID
	CardAdId string `json:"cardAdId" bson:"cardAdId"`
	// 카드명
	CardName string `json:"cardName" bson:"cardName"`
	// 카드사 코드
	CompanyCode string `json:"companyCode" bson:"companyCode"`
	// 카드사 이름
	CompanyName string `json:"companyName" bson:"companyName"`
	// 한줄 광고
	TitleDescription string `json:"titleDescription" bson:"titleDescription"`
	// 카드 이미지
	CardImage string `json:"cardImage" bson:"cardImage"`
	// 카드 이미지 URL
	CardImageUrl string `json:"cardImageUrl" bson:"cardImageUrl"`
	// 카드 신청 URL
	RegisterUrl string `json:"registerUrl" bson:"registerUrl"`
	// 카드 신청 URL (무료)
	RegisterUrlForNoCharge string `json:"registerUrlForNoCharge" bson:"registerUrlForNoCharge"`
	// 연회비 국내
	DomesticAnnualFee int `json:"domesticAnnualFee" bson:"domesticAnnualFee"`
	// 연회비 해외
	ForeignAnnualFee int `json:"foreignAnnualFee" bson:"foreignAnnualFee"`
	// 네이버 페이 모바일 여부
	EnableNpayMO bool `json:"enableNpayMO" bson:"enableNpayMO"`
	// 네이버 페이 PC 여부
	EnableNpayPC bool `json:"enableNpayPC" bson:"enableNpayPC"`
	//
	ImpBeacon string `json:"impBeacon" bson:"impBeacon"`
	// 혜택
	Benefits []Benefit `json:"benefits" bson:"benefits"`
	// 카드 출시일
	ReleaseAt string `json:"releaseAt" bson:"releaseAt"`
	// 전월실적
	Basement int    `json:"basement" bson:"basement"`
	BizType  string `json:"bizType" bson:"bizType"`
	IsMinCPC bool   `json:"isMinCPC" bson:"isMinCPC"`
}

type Benefit struct {
	// 순서
	Order int `json:"order" bson:"order"`
	// 혜택명
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

func FindCardByCardID(card_id string) (model Card, errMsg co.MsgEx) {

	//err = mgm.Coll(&L01201{}).FindByID(id, &model)

	err := inits.MongoDb.Collection(CardCollectionName()).FindOne(context.TODO(), bson.M{"cardAdId": card_id}).Decode(&model)
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

type SearchCard struct {
	//
	comn.Search

	Code string `json:"code" `

	MaxAnnualFee int `json:"maxAnnualFee" `

	Benefits []string `json:"benefits" `

	Basement int `json:"basement" `

	Cards []*Card `json:"cards" `
}

func (search *SearchCard) CollectionName() string {
	//
	return CardCollectionName()
}
func (search *SearchCard) condition() []bson.M {
	matchStage := bson.M{
		"$match": bson.M{},
	}

	// 카드사 검색조건
	if co.NotEmptyString(search.Code) {
		matchStage["$match"].(bson.M)["companyCode"] = search.Code
	}

	// 카드 혜택 검색조건
	if len(search.Benefits) > 0 {
		matchStage["$match"].(bson.M)["benefits.rootBenefitCategoryIdName"] = bson.M{"$all": search.Benefits}
	}

	// 연회비 검색조건
	if search.MaxAnnualFee != 0 {
		matchStage["$match"].(bson.M)["domesticAnnualFee"] = bson.M{"$lte": search.MaxAnnualFee}
	}

	// if search.MaxAnnualFee > 0 {
	// 	matchStage["$match"].(bson.M)["domesticAnnualFee"] = bson.M{"$lte": search.MaxAnnualFee}
	// }

	// 전월실적 검색조건
	if search.Basement != 0 {
		matchStage["$match"].(bson.M)["basement"] = bson.M{"$lte": search.Basement}
	}

	return []bson.M{matchStage}
}

// .
func (search *SearchCard) Finds() (errEx co.MsgEx) {
	pipeline := search.condition()

	sort := bson.M{"createdtime": -1}
	if co.NotEmptyString(search.SortField) {
		if search.SortDirection != 1 {
			search.SortDirection = -1
		} else {
			search.SortDirection = 1
		}
		sort = bson.M{search.SortField: search.SortDirection}
	}

	pipeline = append(pipeline, bson.M{"$sort": sort})

	if search.Limit > 0 && search.PageOffset > -1 {
		pipeline = append(pipeline, bson.M{"$skip": int64(search.Limit) * int64(search.PageOffset)})
		pipeline = append(pipeline, bson.M{"$limit": int64(search.Limit)})
	}

	cursor, err := inits.MongoDb.Collection(search.CollectionName()).Aggregate(
		context.TODO(),
		pipeline,
	)
	if err != nil {
		return co.ErrorPass(err.Error())
	}

	if err = cursor.All(context.TODO(), &search.Cards); err != nil {
		return co.ErrorPass(err.Error())
	}

	// 전체 데이터 갯수
	total, err := inits.MongoDb.Collection(search.CollectionName()).CountDocuments(context.TODO(), pipeline[0]["$match"].(bson.M))
	if err != nil {
		return co.ErrorPass(err.Error())
	}

	search.Total = total

	return errEx
}

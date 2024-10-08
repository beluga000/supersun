package bank

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

// 예금
type Instalment_Savings struct {
	repo.MongoBase `bson:",inline"`

	TypeCode          string   `json:"typeCode" bson:"typeCode"`
	Code              string   `json:"code" bson:"code"`
	Name              string   `json:"name" bson:"name"`
	CompanyCode       string   `json:"companyCode" bson:"companyCode"`
	CompanyName       string   `json:"companyName" bson:"companyName"`
	IsBrokerage       bool     `json:"isBrokerage" bson:"isBrokerage"`
	InterestRate      float64  `json:"interestRate" bson:"interestRate"`
	PrimeInterestRate float64  `json:"primeInterestRate" bson:"primeInterestRate"`
	ProductCategories []string `json:"productCategories" bson:"productCategories"`
	CompanyLogoURL    string   `json:"companyLogoURL" bson:"companyLogoURL"`
	Product_period    int      `json:"product_period" bson:"product_period"`
}

func Instalment_SavingsCollectionName() string {
	return "instalment_savings"
}

func (model *Instalment_Savings) CollectionName() string {
	//
	return Instalment_SavingsCollectionName()
}

/* ****************************************************************************
  Basic CRUD

***************************************************************************** */

// .
func (model *Instalment_Savings) GetById(id string) (errEx co.MsgEx) {

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
func (model *Instalment_Savings) Create() (errEx co.MsgEx) {

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
func (model *Instalment_Savings) Delete() (errEx co.MsgEx) {

	model.Able = false

	_, err := inits.MongoDb.Collection(model.CollectionName()).UpdateOne(context.TODO(), bson.D{{Key: "_id", Value: model.ID}}, bson.M{"$set": model})
	if err != nil {
		return co.ErrorPass(err.Error())
	}
	return co.SuccessPass("")
}

// .
func (model *Instalment_Savings) Update() (errEx co.MsgEx) {

	model.Able = true
	//update := bson.D{}
	model.UpdatedTime = time.Now()

	_, err := inits.MongoDb.Collection(model.CollectionName()).UpdateOne(context.TODO(), bson.D{{Key: "_id", Value: model.ID}}, bson.M{"$set": model})
	if err != nil {
		return co.ErrorPass(err.Error())
	}

	return co.SuccessPass("")
}

func (model *Instalment_Savings) GetList(page string, limit string) (result []*Instalment_Savings, errEx co.MsgEx) {
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
func FindInstalment_SavingsById(id string) (model Instalment_Savings, errMsg co.MsgEx) {

	//err = mgm.Coll(&L01201{}).FindByID(id, &model)

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model, co.ErrorPass(err.Error())
	}
	err = inits.MongoDb.Collection(Instalment_SavingsCollectionName()).FindOne(context.TODO(), bson.M{"_id": objectId}).Decode(&model)
	if err != nil {
		return model, co.ErrorPass(err.Error())
	}

	return model, co.SuccessPass("")
}

func FindInstalment_SavingsByCode(code string) (model Instalment_Savings, errMsg co.MsgEx) {

	err := inits.MongoDb.Collection(Instalment_SavingsCollectionName()).FindOne(context.TODO(), bson.M{"code": code}).Decode(&model)
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
type SearchInstalment_Savings struct {
	//
	comn.Search

	Bank_Name string `json:"bank_name"`

	Period string `json:"period"`

	Categories []string `json:"categories" `

	Basic_Rate_Sort string `json:"basic_rate_sort"`

	Max_Rate_Sort string `json:"max_rate_sort"`

	//

	Instalment_Savingss []*Instalment_Savings `json:"Instalment_Savingss"`
}

func (search *SearchInstalment_Savings) CollectionName() string {
	//
	return Instalment_SavingsCollectionName()
}

// .
func (search *SearchInstalment_Savings) condition() []bson.M {

	matchStage := bson.M{
		"$match": bson.M{}}

	if co.NotEmptyString(search.Bank_Name) {
		matchStage["$match"].(bson.M)["companyName"] = search.Bank_Name
	}

	if co.NotEmptyString(search.Period) && search.Period != "전체" {
		period, _ := strconv.Atoi(search.Period)
		matchStage["$match"].(bson.M)["product_period"] = bson.M{"$lte": period}
	}

	if len(search.Categories) > 0 {
		matchStage["$match"].(bson.M)["productCategories"] = bson.M{"$all": search.Categories}
	}

	return []bson.M{matchStage}

}

// .
func (search *SearchInstalment_Savings) Finds() (errEx co.MsgEx) {
	pipeline := search.condition()

	sort := bson.M{"createdtime": -1}

	if search.Basic_Rate_Sort == "asc" {
		sort = bson.M{"interestRate": 1}
	} else if search.Basic_Rate_Sort == "desc" {
		sort = bson.M{"interestRate": -1}
	}

	if search.Max_Rate_Sort == "asc" {
		sort = bson.M{"primeInterestRate": 1}
	} else if search.Max_Rate_Sort == "desc" {
		sort = bson.M{"primeInterestRate": -1}
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

	if err = cursor.All(context.TODO(), &search.Instalment_Savingss); err != nil {
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

func (search *SearchInstalment_Savings) Finds_Top3() (errEx co.MsgEx) {
	pipeline := search.condition()

	sort := bson.M{"primeInterestRate": -1}

	pipeline = append(pipeline, bson.M{"$sort": sort})
	pipeline = append(pipeline, bson.M{"$limit": 3})

	cursor, err := inits.MongoDb.Collection(search.CollectionName()).Aggregate(
		context.TODO(),
		pipeline,
	)
	if err != nil {
		return co.ErrorPass(err.Error())
	}

	if err = cursor.All(context.TODO(), &search.Instalment_Savingss); err != nil {
		return co.ErrorPass(err.Error())
	}

	total, err := inits.MongoDb.Collection(search.CollectionName()).CountDocuments(context.TODO(), pipeline[0]["$match"].(bson.M))
	if err != nil {
		return co.ErrorPass(err.Error())
	}

	search.Total = total

	return errEx
}

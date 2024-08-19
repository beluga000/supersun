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

type Deposit_Detail struct {
	repo.MongoBase `bson:",inline"`

	Code             string   `json:"code" bson:"code"`
	Product_name     string   `json:"product_name" bson:"product_name"`
	Company_code     string   `json:"company_code" bson:"company_code"`
	Bank_name        string   `json:"bank_name" bson:"bank_name"`
	Product_category []string `json:"product_category" bson:"product_category"`
	// 기본이율
	Basic_rate float64 `json:"basic_rate" bson:"basic_rate"`
	// 최고이율
	Max_rate float64 `json:"max_rate" bson:"max_rate"`
	// 최대 가입기간
	Product_period                    int      `json:"product_period" bson:"product_period"`
	SpecialOffer_Summary              string   `json:"specialOffer_Summary" bson:"specialOffer_Summary"`
	SpecialOffer_Period               string   `json:"specialOffer_Period" bson:"specialOffer_Period"`
	Join_period                       string   `json:"join_period" bson:"join_period"`
	Join_amount                       string   `json:"join_amount" bson:"join_amount"`
	Join_target                       string   `json:"join_target" bson:"join_target"`
	Join_channel                      string   `json:"join_channel" bson:"join_channel"`
	Join_payment                      string   `json:"join_payment" bson:"join_payment"`
	Join_note                         string   `json:"join_note" bson:"join_note"`
	Join_protection                   string   `json:"join_protection" bson:"join_protection"`
	Join_deliberationNumber           string   `json:"join_deliberationNumber" bson:"join_deliberationNumber"`
	Join_deliberationNumber_period    string   `json:"join_deliberationNumber_period" bson:"join_deliberationNumber_period"`
	Company_tel                       string   `json:"company_tel" bson:"company_tel"`
	Company_pcLink                    string   `json:"company_pcLink" bson:"company_pcLink"`
	Company_mobileLink                string   `json:"company_mobileLink" bson:"company_mobileLink"`
	Rate_kind                         string   `json:"rate_kind" bson:"rate_kind"`
	RateTable_period                  string   `json:"rateTable_period" bson:"rateTable_period"`
	SpecialCondition_description      string   `json:"specialCondition_description" bson:"specialCondition_description"`
	SpecialCondition_description_info []string `json:"specialCondition_description_info" bson:"specialCondition_description_info"`
	// 월 최대납입 금액
	Amount_min int `json:"amount_min" bson:"amount_min"`
	// 월 최소납입 금액
	Amount_max int `json:"amount_max" bson:"amount_max"`
	// 누구나 가입 가능조건
	Is_everyone bool `json:"is_everyone" bson:"is_everyone"`
	// 청년전용
	Is_young bool `json:"is_young" bson:"is_young"`
	// 사업자전용
	Is_business bool `json:"is_business" bson:"is_business"`
	// 자녀전용
	Is_children bool `json:"is_children" bson:"is_children"`
	// 취약계층전용
	Is_vulnerable_social_group bool `json:"is_vulnerable_social_group" bson:"is_vulnerable_social_group"`
	// 군인전용
	Is_soldier bool `json:"is_soldier" bson:"is_soldier"`
	// 노인전용
	Is_old bool `json:"is_old" bson:"is_old"`
}

func Deposit_DetailCollectionName() string {
	return "deposit_detail"
}

func (model *Deposit_Detail) CollectionName() string {
	//
	return Deposit_DetailCollectionName()
}

/* ****************************************************************************
  Basic CRUD

***************************************************************************** */

// .
func (model *Deposit_Detail) GetById(id string) (errEx co.MsgEx) {

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
func (model *Deposit_Detail) Create() (errEx co.MsgEx) {

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
func (model *Deposit_Detail) Delete() (errEx co.MsgEx) {

	model.Able = false

	_, err := inits.MongoDb.Collection(model.CollectionName()).UpdateOne(context.TODO(), bson.D{{Key: "_id", Value: model.ID}}, bson.M{"$set": model})
	if err != nil {
		return co.ErrorPass(err.Error())
	}
	return co.SuccessPass("")
}

// .
func (model *Deposit_Detail) Update() (errEx co.MsgEx) {

	model.Able = true
	//update := bson.D{}
	model.UpdatedTime = time.Now()

	_, err := inits.MongoDb.Collection(model.CollectionName()).UpdateOne(context.TODO(), bson.D{{Key: "_id", Value: model.ID}}, bson.M{"$set": model})
	if err != nil {
		return co.ErrorPass(err.Error())
	}

	return co.SuccessPass("")
}

func (model *Deposit_Detail) GetList(page string, limit string) (result []*Deposit_Detail, errEx co.MsgEx) {
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
func FindDeposit_DetailById(id string) (model Deposit_Detail, errMsg co.MsgEx) {

	//err = mgm.Coll(&L01201{}).FindByID(id, &model)

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model, co.ErrorPass(err.Error())
	}
	err = inits.MongoDb.Collection(Deposit_DetailCollectionName()).FindOne(context.TODO(), bson.M{"_id": objectId}).Decode(&model)
	if err != nil {
		return model, co.ErrorPass(err.Error())
	}

	return model, co.SuccessPass("")
}

func FindDeposit_DetailByCode(code string) (model Deposit_Detail, errMsg co.MsgEx) {

	err := inits.MongoDb.Collection(Deposit_DetailCollectionName()).FindOne(context.TODO(), bson.M{"code": code}).Decode(&model)
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
type SearchDeposit_Detail struct {
	//
	comn.Search

	//

	Bank_name string `json:"bank_name" bson:"bank_name"`

	Period string `json:"period" bson:"period"`

	// Youth string `json:"youth" bson:"youth"`

	Max_Rate_Sort string `json:"max_rate_sort"`

	Business string `json:"business" bson:"business"`

	Children string `json:"children" bson:"children"`

	Vulnerable_social_group string `json:"vulnerable_social_group" bson:"vulnerable_social_group"`

	Young string `json:"young" bson:"young"`

	Soldier string `json:"soldier" bson:"soldier"`

	Old string `json:"old" bson:"old"`

	Deposit_Details []*Deposit_Detail
}

func (search *SearchDeposit_Detail) CollectionName() string {
	//
	return Deposit_DetailCollectionName()
}
func addMatchCondition(matchStage bson.M, field string, value string) {
	if value == "N" {
		matchStage[field] = false
	} else if value == "Y" {
		if _, ok := matchStage["$or"]; !ok {
			matchStage["$or"] = []bson.M{}
		}
		matchStage["$or"] = append(matchStage["$or"].([]bson.M), bson.M{field: true}, bson.M{field: false}, bson.M{field: bson.M{"$exists": false}})
	}
}

// .
func (search *SearchDeposit_Detail) condition() []bson.M {
	matchStage := bson.M{"$match": bson.M{}}

	if co.NotEmptyString(search.Bank_name) {
		matchStage["$match"].(bson.M)["bank_name"] = search.Bank_name
	}

	if co.NotEmptyString(search.Period) && search.Period != "전체" {
		period, _ := strconv.Atoi(search.Period)
		matchStage["$match"].(bson.M)["product_period"] = bson.M{"$lte": period}
	}

	if co.NotEmptyString(search.Children) {
		addMatchCondition(matchStage["$match"].(bson.M), "is_children", search.Children)
	}

	if co.NotEmptyString(search.Business) {
		addMatchCondition(matchStage["$match"].(bson.M), "is_business", search.Business)
	}

	if co.NotEmptyString(search.Vulnerable_social_group) {
		addMatchCondition(matchStage["$match"].(bson.M), "is_vulnerable_social_group", search.Vulnerable_social_group)
	}

	if co.NotEmptyString(search.Young) {
		addMatchCondition(matchStage["$match"].(bson.M), "is_young", search.Young)
	}

	if co.NotEmptyString(search.Soldier) {
		addMatchCondition(matchStage["$match"].(bson.M), "is_soldier", search.Soldier)
	}

	if co.NotEmptyString(search.Old) {
		addMatchCondition(matchStage["$match"].(bson.M), "is_old", search.Old)
	}

	return []bson.M{matchStage}
}

// .
func (search *SearchDeposit_Detail) Finds() (errEx co.MsgEx) {

	pipeline := search.condition()

	sort := bson.M{"createdtime": -1}
	if search.Max_Rate_Sort == "asc" {
		sort = bson.M{"max_rate": 1}
	} else if search.Max_Rate_Sort == "desc" {
		sort = bson.M{"max_rate": -1}
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

	if err = cursor.All(context.TODO(), &search.Deposit_Details); err != nil {
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

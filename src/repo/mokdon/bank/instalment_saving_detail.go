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

type Instalment_Savings_Detail struct {
	repo.MongoBase `bson:",inline"`

	Code                              string   `json:"code" bson:"code"`
	Product_name                      string   `json:"product_name" bson:"product_name"`
	Company_code                      string   `json:"company_code" bson:"company_code"`
	Bank_name                         string   `json:"bank_name" bson:"bank_name"`
	Product_category                  []string `json:"product_category" bson:"product_category"`
	Basic_rate                        string   `json:"basic_rate" bson:"basic_rate"`
	Max_rate                          string   `json:"max_rate" bson:"max_rate"`
	Product_period                    string   `json:"product_period" bson:"product_period"`
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
	RateTable_head                    []string `json:"rateTable_head" bson:"rateTable_head"`
	RateTable_rate                    []string `json:"rateTable_rate" bson:"rateTable_rate"`
	SpecialCondition_description      string   `json:"specialCondition_description" bson:"specialCondition_description"`
	SpecialCondition_description_info []string `json:"specialCondition_description_info" bson:"specialCondition_description_info"`
}

func Instalment_Savings_DetailCollectionName() string {
	return "instalment_savings_detail"
}

func (model *Instalment_Savings_Detail) CollectionName() string {
	//
	return Instalment_Savings_DetailCollectionName()
}

/* ****************************************************************************
  Basic CRUD

***************************************************************************** */

// .
func (model *Instalment_Savings_Detail) GetById(id string) (errEx co.MsgEx) {

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
func (model *Instalment_Savings_Detail) Create() (errEx co.MsgEx) {

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
func (model *Instalment_Savings_Detail) Delete() (errEx co.MsgEx) {

	model.Able = false

	_, err := inits.MongoDb.Collection(model.CollectionName()).UpdateOne(context.TODO(), bson.D{{Key: "_id", Value: model.ID}}, bson.M{"$set": model})
	if err != nil {
		return co.ErrorPass(err.Error())
	}
	return co.SuccessPass("")
}

// .
func (model *Instalment_Savings_Detail) Update() (errEx co.MsgEx) {

	model.Able = true
	//update := bson.D{}
	model.UpdatedTime = time.Now()

	_, err := inits.MongoDb.Collection(model.CollectionName()).UpdateOne(context.TODO(), bson.D{{Key: "_id", Value: model.ID}}, bson.M{"$set": model})
	if err != nil {
		return co.ErrorPass(err.Error())
	}

	return co.SuccessPass("")
}

func (model *Instalment_Savings_Detail) GetList(page string, limit string) (result []*Instalment_Savings_Detail, errEx co.MsgEx) {
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
func FindInstalment_Savings_DetailById(id string) (model Instalment_Savings_Detail, errMsg co.MsgEx) {

	//err = mgm.Coll(&L01201{}).FindByID(id, &model)

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model, co.ErrorPass(err.Error())
	}
	err = inits.MongoDb.Collection(Instalment_Savings_DetailCollectionName()).FindOne(context.TODO(), bson.M{"_id": objectId}).Decode(&model)
	if err != nil {
		return model, co.ErrorPass(err.Error())
	}

	return model, co.SuccessPass("")
}

func FindInstalment_Savings_DetailByCode(code string) (model Instalment_Savings_Detail, errMsg co.MsgEx) {

	err := inits.MongoDb.Collection(Instalment_Savings_DetailCollectionName()).FindOne(context.TODO(), bson.M{"code": code}).Decode(&model)
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
type SearchInstalment_Savings_Detail struct {
	//
	comn.Search

	//

	Instalment_Savings_Details []*Instalment_Savings_Detail
}

func (search *SearchInstalment_Savings_Detail) CollectionName() string {
	//
	return Instalment_Savings_DetailCollectionName()
}

// .
func (search *SearchInstalment_Savings_Detail) condition() bson.M {

	filter := bson.M{}

	return filter

}

// .
func (search *SearchInstalment_Savings_Detail) Finds() (errEx co.MsgEx) {
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

		if err = cursor.All(context.TODO(), &search.Instalment_Savings_Details); err != nil {
			return co.ErrorPass(err.Error())
		}

	} else {

		cursor, err := inits.MongoDb.Collection(search.CollectionName()).Find(context.TODO(), search.condition(), options.Find().SetSort(sort))
		if err != nil {
			return co.ErrorPass(err.Error())
		}

		if err = cursor.All(context.TODO(), &search.Instalment_Savings_Details); err != nil {
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

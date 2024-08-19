package encryption

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

type User struct {
	repo.MongoBase `bson:",inline"`

	Id string `json:"id" bson:"id"`

	Pwd string `json:"pwd" bson:"pwd"`

	Name string `json:"name" bson:"name"`

	Email string `json:"email" bson:"email"`

	Age int `json:"age" bson:"age"`
}

func UserCollectionName() string {
	return "user"
}

func (model *User) CollectionName() string {
	//
	return UserCollectionName()
}

/* ****************************************************************************
  Basic CRUD

***************************************************************************** */

// .
func (model *User) GetById(id string) (errEx co.MsgEx) {

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
func (model *User) Create() (errEx co.MsgEx) {

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
func (model *User) Delete() (errEx co.MsgEx) {

	model.Able = false

	_, err := inits.MongoDb.Collection(model.CollectionName()).UpdateOne(context.TODO(), bson.D{{Key: "_id", Value: model.ID}}, bson.M{"$set": model})
	if err != nil {
		return co.ErrorPass(err.Error())
	}
	return co.SuccessPass("")
}

// .
func (model *User) Update() (errEx co.MsgEx) {

	model.Able = true
	//update := bson.D{}
	model.UpdatedTime = time.Now()

	_, err := inits.MongoDb.Collection(model.CollectionName()).UpdateOne(context.TODO(), bson.D{{Key: "_id", Value: model.ID}}, bson.M{"$set": model})
	if err != nil {
		return co.ErrorPass(err.Error())
	}

	return co.SuccessPass("")
}

func (model *User) GetList(page string, limit string) (result []*User, errEx co.MsgEx) {
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
func FindUserById(id string) (model User, errMsg co.MsgEx) {

	//err = mgm.Coll(&L01201{}).FindByID(id, &model)

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model, co.ErrorPass(err.Error())
	}
	err = inits.MongoDb.Collection(UserCollectionName()).FindOne(context.TODO(), bson.M{"_id": objectId}).Decode(&model)
	if err != nil {
		return model, co.ErrorPass(err.Error())
	}

	return model, co.SuccessPass("")
}

func FindUserData(id string) (model User, errMsg co.MsgEx) {

	//err = mgm.Coll(&L01201{}).FindByID(id, &model)

	err := inits.MongoDb.Collection(UserCollectionName()).FindOne(context.TODO(), bson.M{"id": id}).Decode(&model)
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
type SearchUser struct {
	//
	comn.Search

	//

	Users []*User
}

func (search *SearchUser) CollectionName() string {
	//
	return UserCollectionName()
}

// .
func (search *SearchUser) condition() bson.M {

	filter := bson.M{}

	filter["able"] = true

	return filter

}

// .
func (search *SearchUser) Finds() (errEx co.MsgEx) {
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

		if err = cursor.All(context.TODO(), &search.Users); err != nil {
			return co.ErrorPass(err.Error())
		}

	} else {

		cursor, err := inits.MongoDb.Collection(search.CollectionName()).Find(context.TODO(), search.condition(), options.Find().SetSort(sort))
		if err != nil {
			return co.ErrorPass(err.Error())
		}

		if err = cursor.All(context.TODO(), &search.Users); err != nil {
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

package local

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

type Member struct {

	// .
	repo.MongoBase `bson:",inline"`

	Name string `bson:"name"   json:"name"`

	Email string `bson:"email"   json:"email"`

	Phone string `bson:"phone"   json:"phone"`
}

// .

func MemberCollectionName() string {
	return "member"
}

func (model *Member) CollectionName() string {
	//
	return MemberCollectionName()
}

/* ****************************************************************************
  validate

***************************************************************************** */

// .
func (model *Member) validate() bool {

	// nomalis = true

	// if len(model.QuestionTitle) < 2 {
	// 	nomalis = false
	// 	erris = append(erris, []string{"question_title", "* 제목을 입력해주세요."})
	// }
	// if len(model.QuestionContent) < 2 {
	// 	nomalis = false
	// 	erris = append(erris, []string{"question_content", "* 문의 내용을 입력해주세요."})
	// }

	// log.Print(" erris : > ", erris)
	// // check final
	// if !nomalis {
	// 	return nomalis, erris
	// } else {
	// 	return true, nil

	// }
	return true
}

/* ****************************************************************************
  Basic CRUD

***************************************************************************** */

// .
func (model *Member) GetById(id string) (errEx co.MsgEx) {

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
func (model *Member) Create() (errEx co.MsgEx) {

	// nomalis, erris := model.validate()

	// log.Print(" nomalis : ", nomalis)
	// log.Print(" erris : ", erris)

	// if !nomalis {
	// 	errEx.Failure = true
	// 	errEx.Success = false
	// 	errEx.Message = "입력값이 잘못되었습니다."
	// 	errEx.ValidateErr = erris

	// 	return errEx
	// }

	model.Able = true
	model.CreatedTime = time.Now()

	if !model.validate() {
		return errEx
	}

	_, err := inits.MongoDb.Collection(model.CollectionName()).InsertOne(context.TODO(), model)
	if err != nil {
		log.Print("  >> err : ", err.Error())
		return co.ErrorPass(err.Error())
	}

	return co.SuccessPass("")
}

// .
func (model *Member) Delete() (errEx co.MsgEx) {

	model.Able = false

	_, err := inits.MongoDb.Collection(model.CollectionName()).UpdateOne(context.TODO(), bson.D{{Key: "_id", Value: model.ID}}, bson.M{"$set": model})
	if err != nil {
		return co.ErrorPass(err.Error())
	}
	return co.SuccessPass("")
}

// .
func (model *Member) Update() (errEx co.MsgEx) {

	model.Able = true
	//update := bson.D{}
	model.UpdatedTime = time.Now()

	_, err := inits.MongoDb.Collection(model.CollectionName()).UpdateOne(context.TODO(), bson.D{{Key: "_id", Value: model.ID}}, bson.M{"$set": model})
	if err != nil {
		return co.ErrorPass(err.Error())
	}

	return co.SuccessPass("")
}

// .
func (model *Member) UpdateWorkStatusId() (errEx co.MsgEx) {

	//update := bson.D{}
	model.UpdatedTime = time.Now()

	// model.WorkStatusId = "12312"

	return co.SuccessPass("")
}

func (model *Member) GetList(page string, limit string) (result []*Member, errEx co.MsgEx) {
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
func FindMemberById(id string) (model Member, errMsg co.MsgEx) {

	//err = mgm.Coll(&L01201{}).FindByID(id, &model)

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model, co.ErrorPass(err.Error())
	}
	err = inits.MongoDb.Collection(MemberCollectionName()).FindOne(context.TODO(), bson.M{"_id": objectId}).Decode(&model)
	if err != nil {
		return model, co.ErrorPass(err.Error())
	}

	return model, co.SuccessPass("")
}

func FindDataRequestByIdPlusCount(id string) (model Member, errMsg co.MsgEx) {

	//err = mgm.Coll(&L01201{}).FindByID(id, &model)

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model, co.ErrorPass(err.Error())
	}
	err = inits.MongoDb.Collection(MemberCollectionName()).FindOne(context.TODO(), bson.M{"_id": objectId}).Decode(&model)
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
type SearchMember struct {
	//
	comn.Search

	//

	Members []*Member
}

func (search *SearchMember) CollectionName() string {
	//
	return MemberCollectionName()
}

// .
func (search *SearchMember) condition() bson.M {

	filter := bson.M{}

	filter["able"] = true

	// cateUrlDecode, _ := url.PathUnescape(search.Cat)

	return filter

}

// .
func (search *SearchMember) Finds() (errEx co.MsgEx) {
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

		if err = cursor.All(context.TODO(), &search.Members); err != nil {
			return co.ErrorPass(err.Error())
		}

	} else {

		cursor, err := inits.MongoDb.Collection(search.CollectionName()).Find(context.TODO(), search.condition(), options.Find().SetSort(sort))
		if err != nil {
			return co.ErrorPass(err.Error())
		}

		if err = cursor.All(context.TODO(), &search.Members); err != nil {
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

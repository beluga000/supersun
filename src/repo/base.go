package repo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
	"sunny.ksw.kr/co"
	"sunny.ksw.kr/inits"
)

//.

type MongoBase struct {

	//.
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`

	CreatedTime time.Time `json:"createdtime" bson:"createdtime,omitempty"`

	Able bool `json:"-" bson:"able"`

	DeleteAt time.Time `json:"-" bson:"deleteat,omitempty"`

	UpdatedTime time.Time `json:"-" bson:"updatedtime,omitempty"`

	Writer string `json:"writer" bson:"writer,omitempty"`
}

/*
PrepareID(id interface{}) (interface{}, error)
GetID() interface{}
SetID(id interface{})
*/
func (model *MongoBase) SetID(id interface{}) {
	//func (model *MongoBase) SetID(st primitive.ObjectID) {

	model.ID = id.(primitive.ObjectID)
	// return "ins_file_info"
}

func (model *MongoBase) GetID() interface{} {

	//model.ID = st
	// return "ins_file_info"
	return "interface{}"
}

func (model *MongoBase) PrepareID(id interface{}) (interface{}, error) {

	return "", nil
}

type MariaBase struct {

	//.
	ID          string `json:"id"  gorm:"id" `
	Active      uint   `json:"-"  gorm:"active" `
	CreatedTime int64  `json:"-"  gorm:"created_time" `
	UpdatedTime int64  `json:"-"  gorm:"updated_time" `

	//.
	AuditEvent    string `json:"audit_event"  gorm:"audit_event" `
	AuditUserId   string `json:"audit_user_id"  gorm:"audit_user_id" `
	AuditUserName string `json:"audit_user_name"  gorm:"audit_user_name" `
}

// .
func GetCollectionNames() ([]string, co.MsgEx) {

	names, err := inits.MongoDb.ListCollectionNames(context.TODO(), bson.M{})
	if err != nil {
		return []string{}, co.ErrorPass(err.Error())

	}

	return names, co.SuccessPass("")
}

// .
func GetCollectIndexs(collection string) ([]string, co.MsgEx) {

	rtnarray := []string{}
	cursor, err := inits.MongoDb.Collection(collection).Indexes().List(context.TODO()) //
	if err != nil {
		//log.Print(err)
		return []string{}, co.ErrorPass(err.Error())
	}

	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		// log.Fatal(err)
		return []string{}, co.ErrorPass(err.Error())
	}

	for _, v := range results {

		for k2, _ := range v["key"].(bson.M) {
			rtnarray = append(rtnarray, k2)
		}

	}
	// log.Print("   ...  ", rtnarray)

	return rtnarray, co.SuccessPass("")

}

func GetCollectCount(collection string) (int64, co.MsgEx) {

	count, err := inits.MongoDb.Collection(collection).CountDocuments(context.TODO(), bson.M{}) //
	if err != nil {
		//log.Print(err)
		return 0, co.ErrorPass(err.Error())
	}

	return count, co.SuccessPass("")
}

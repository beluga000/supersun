package inits

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sunny.ksw.kr/co"

	"github.com/kamva/mgm/v3"
)

// .
var MongoDb *mongo.Database

// .
var MongoDbQueen *mongo.Database

var MongoDbLog *mongo.Database

var MongoDbLocal *mongo.Database

// .
func MongoInit(uri string, db string) (errEx co.MsgEx) {

	// Connection URI
	//const uri = "mongodb://user:pass@sample.host:27017/?maxPoolSize=20&w=majority"

	//uri := "mongodb://localhost:27017/?maxPoolSize=20&w=majority"
	//uri := os.Getenv("MONGODB_URI")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		//
		fmt.Print("MongoInit Error: " + err.Error())
		return co.ErrorPass(err.Error())

	}
	// defer func() {
	// 	if err := client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()

	// MongoDb = client.Database("test")
	MongoDb = client.Database(db)

	return co.SuccessPass("")

}

// .
func MongoInitDebug(uri string, db string) (errEx co.MsgEx) {

	cmdMonitor := &event.CommandMonitor{
		Started: func(_ context.Context, evt *event.CommandStartedEvent) {
			fmt.Print(evt.Command)
		},
	}

	// Connection URI
	//const uri = "mongodb://user:pass@sample.host:27017/?maxPoolSize=20&w=majority"

	//uri := "mongodb://localhost:27017/?maxPoolSize=20&w=majority"
	//uri := os.Getenv("MONGODB_URI")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri).SetMonitor(cmdMonitor))
	if err != nil {
		//
		return co.ErrorPass(err.Error())

	}
	// defer func() {
	// 	if err := client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()

	// MongoDb = client.Database("test")
	MongoDb = client.Database(db)

	return co.SuccessPass("")

}

// .
func MongoQueenInit(uri string, db string) (errEx co.MsgEx) {

	// Connection URI
	//const uri = "mongodb://user:pass@sample.host:27017/?maxPoolSize=20&w=majority"

	//uri := "mongodb://localhost:27017/?maxPoolSize=20&w=majority"
	//uri := os.Getenv("MONGODB_URI")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return co.ErrorPass(err.Error())
	}
	// defer func() {
	// 	if err := client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()

	// MongoDb = client.Database("test")
	MongoDbQueen = client.Database(db)

	return co.SuccessPass("")

}

// .
func MongoQueenInitDebug(uri string, db string) (errEx co.MsgEx) {

	cmdMonitor := &event.CommandMonitor{
		Started: func(_ context.Context, evt *event.CommandStartedEvent) {
			fmt.Print(evt.Command)
		},
	}

	// Connection URI
	//const uri = "mongodb://user:pass@sample.host:27017/?maxPoolSize=20&w=majority"

	//uri := "mongodb://localhost:27017/?maxPoolSize=20&w=majority"
	//uri := os.Getenv("MONGODB_URI")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri).SetMonitor(cmdMonitor))
	if err != nil {
		return co.ErrorPass(err.Error())
	}
	// defer func() {
	// 	if err := client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()

	// MongoDb = client.Database("test")
	MongoDbQueen = client.Database(db)

	return co.SuccessPass("")

}

// .
func MongoLocalInit(uri string, db string) (errEx co.MsgEx) {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return co.ErrorPass("MongoLocalInit : " + err.Error())
	}

	MongoDbLocal = client.Database(db)

	return co.SuccessPass("")

}

// .
func MongoLocalInitDebug(uri string, db string) (errEx co.MsgEx) {

	cmdMonitor := &event.CommandMonitor{
		Started: func(_ context.Context, evt *event.CommandStartedEvent) {
			fmt.Print(evt.Command)
		},
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri).SetMonitor(cmdMonitor))
	if err != nil {
		return co.ErrorPass("MongoLocalInitDebug : " + err.Error())
	}

	MongoDbQueen = client.Database(db)

	return co.SuccessPass("")

}

type SoftDelete struct {
	Able bool `json:"able" bson:"able"`

	DeleteAt time.Time `json:"delete_at" bson:"delete_at"`
}

/*
service_time_uuid
*/
type RequestWorkId string

type RequestWork struct {
	ServiceSpec string `json:"service_spec" bson:"service_spec"`

	WorkAt time.Time `json:"work_at" bson:"work_at"`

	ServiceId string
}

/* gwork CreateIndex */
func CreateIndex(CollectionName string, field string, unique bool) bool {

	// 1. Lets define the keys for the index we want to create
	mod := mongo.IndexModel{
		Keys:    bson.M{field: 1}, // index in ascending order or -1 for descending order
		Options: options.Index().SetUnique(unique),
	}
	//_, err := mgm.CollectionByName(CollectionName).Indexes().CreateOne(mgm.Ctx(), mod)

	_, err := MongoDb.Collection(CollectionName).Indexes().CreateOne(context.TODO(), mod)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	// 6. All went well, we return true
	return true
}

/* gframe CreateIndex */
// func CreateIndex(CollectionName string, field string, unique bool) bool {

// 	// 1. Lets define the keys for the index we want to create
// 	mod := mongo.IndexModel{
// 		Keys:    bson.M{field: 1}, // index in ascending order or -1 for descending order
// 		Options: options.Index().SetUnique(unique),
// 	}
// 	//_, err := mgm.CollectionByName(CollectionName).Indexes().CreateOne(mgm.Ctx(), mod)
// 	//mgm.Ctx()
// 	_, err := MongoDb.Collection(CollectionName).Indexes().CreateOne(context.TODO(), mod)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return false
// 	}
// 	// 6. All went well, we return true
// 	return true
// }

// .
func MongoLogInit(uri string, db string) (errEx co.MsgEx) {

	// Connection URI
	//const uri = "mongodb://user:pass@sample.host:27017/?maxPoolSize=20&w=majority"

	//uri := "mongodb://localhost:27017/?maxPoolSize=20&w=majority"
	//uri := os.Getenv("MONGODB_URI")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		//
		return co.ErrorPass(err.Error())

	}
	// defer func() {
	// 	if err := client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()

	// MongoDb = client.Database("test")
	MongoDbLog = client.Database(db)

	return co.SuccessPass("")

}

func MongoInit2(dbname string, connurl string) {
	// Setup mgm default config

	// err := mgm.SetDefaultConfig(&mgm.Config{CtxTimeout: 20 * time.Second}, dbname, options.Client().ApplyURI(connurl))
	err := mgm.SetDefaultConfig(&mgm.Config{CtxTimeout: 60 * time.Second}, dbname, options.Client().ApplyURI(connurl))

	if err != nil {
		panic("")
	} else {
		fmt.Print("!@MongoDB 연결성공!!")
	}

}

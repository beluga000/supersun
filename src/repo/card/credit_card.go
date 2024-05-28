package card

// type Credit_Card struct {
// 	repo.MongoBase `bson:",inline"`

// 	// 카드명
// 	Name string `bson:"name"   json:"name"`
// 	// 카드사
// 	Company string `bson:"company"   json:"company"`
// 	// 유형
// 	Type string `bson:"type"   json:"type"`
// 	// 혜택
// 	Benefits []string `bson:"benefits"   json:"benefits"`
// 	// 주요 혜택
// 	Main_Benefit []string `bson:"main_benefit"   json:"main_benefit"`
// 	// 전월실적
// 	Previous_Performance string `bson:"previous_performance"   json:"previous_performance"`
// 	// 연회비
// 	Annual_Fee []Annual_Fee `bson:"annual_fee"   json:"annual_fee"`
// 	// 한달 상세 최대 혜택
// 	Max_Benefit_Per_Month_Detail []Max_Benefit_Per_Month `bson:"max_benefit_per_month_detail"   json:"max_benefit_per_month_detail"`
// 	// 제공 서비스
// 	Service []Service `bson:"service"   json:"service"`
// 	// 카드 이미지 경로
// 	Image_Path string `bson:"image_path"   json:"image_path"`
// }

// type Service struct {
// 	// 서비스명
// 	Name string `bson:"name"   json:"name"`
// 	// 서비스 내용
// 	Content string `bson:"content"   json:"content"`
// }

// type Max_Benefit_Per_Month struct {

// 	// 금액
// 	Amount string `bson:"amount"   json:"amount"`

// 	// 최대할인한도
// 	Max_Discount_Limit int `bson:"max_discount_limit"   json:"max_discount_limit"`

// 	// 혜택
// 	Benefits []Benefit `bson:"benefits"   json:"benefits"`
// }

// type Benefit struct {

// 	// 혜택명
// 	Name string `bson:"name"   json:"name"`

// 	// 금액
// 	Amount string `bson:"amount"   json:"amount"`
// }

// type Annual_Fee struct {
// 	Type string `bson:"type"   json:"type"`

// 	Fee string `bson:"fee"   json:"fee"`
// }

// func Credit_CardCollectionName() string {
// 	return "credit_card"
// }

// func (model *Credit_Card) CollectionName() string {
// 	//
// 	return Credit_CardCollectionName()
// }

// /* ****************************************************************************
//   Basic CRUD

// ***************************************************************************** */

// // .
// func (model *Credit_Card) GetById(id string) (errEx co.MsgEx) {

// 	objectId, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return co.ErrorPass(err.Error())
// 	}

// 	err = inits.MongoDb.Collection(model.CollectionName()).FindOne(context.TODO(), bson.D{{Key: "_id", Value: objectId}}).Decode(&model)
// 	if err != nil {
// 		return co.ErrorPass(err.Error())
// 	}

// 	return co.SuccessPass("")

// }

// // .
// func (model *Credit_Card) Create() (errEx co.MsgEx) {

// 	model.Able = true
// 	model.CreatedTime = time.Now()

// 	_, err := inits.MongoDb.Collection(model.CollectionName()).InsertOne(context.TODO(), model)
// 	if err != nil {
// 		log.Print("  >> err : ", err.Error())
// 		return co.ErrorPass(err.Error())
// 	}

// 	return co.SuccessPass("")
// }

// // .
// func (model *Credit_Card) Delete() (errEx co.MsgEx) {

// 	model.Able = false

// 	_, err := inits.MongoDb.Collection(model.CollectionName()).UpdateOne(context.TODO(), bson.D{{Key: "_id", Value: model.ID}}, bson.M{"$set": model})
// 	if err != nil {
// 		return co.ErrorPass(err.Error())
// 	}
// 	return co.SuccessPass("")
// }

// // .
// func (model *Credit_Card) Update() (errEx co.MsgEx) {

// 	model.Able = true
// 	//update := bson.D{}
// 	model.UpdatedTime = time.Now()

// 	_, err := inits.MongoDb.Collection(model.CollectionName()).UpdateOne(context.TODO(), bson.D{{Key: "_id", Value: model.ID}}, bson.M{"$set": model})
// 	if err != nil {
// 		return co.ErrorPass(err.Error())
// 	}

// 	return co.SuccessPass("")
// }

// func (model *Credit_Card) GetList(page string, limit string) (result []*Credit_Card, errEx co.MsgEx) {
// 	pageInt, _ := strconv.Atoi(page)
// 	limitInt, _ := strconv.Atoi(limit)

// 	skip := 0
// 	if pageInt > 0 {
// 		skip = (pageInt - 1) * limitInt
// 	}
// 	cursor, err := inits.MongoDb.Collection(model.CollectionName()).Find(context.TODO(), bson.D{{Key: "able", Value: true}}, options.Find().SetSkip(int64(skip)).SetLimit(int64(limitInt)).SetSort(bson.M{"createdtime": -1}))
// 	if err != nil {
// 		return nil, co.ErrorPass(err.Error())
// 	} else {
// 		cursor.All(context.TODO(), &result)
// 	}
// 	return result, co.SuccessPass("")
// }

// /* ****************************************************************************
//   Find
// ***************************************************************************** */
// // .
// func FindCredit_CardById(id string) (model Credit_Card, errMsg co.MsgEx) {

// 	//err = mgm.Coll(&L01201{}).FindByID(id, &model)

// 	objectId, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return model, co.ErrorPass(err.Error())
// 	}
// 	err = inits.MongoDb.Collection(Credit_CardCollectionName()).FindOne(context.TODO(), bson.M{"_id": objectId}).Decode(&model)
// 	if err != nil {
// 		return model, co.ErrorPass(err.Error())
// 	}

// 	return model, co.SuccessPass("")
// }

// func FindDataRequestByIdPlusCount(id string) (model Credit_Card, errMsg co.MsgEx) {

// 	//err = mgm.Coll(&L01201{}).FindByID(id, &model)

// 	objectId, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return model, co.ErrorPass(err.Error())
// 	}
// 	err = inits.MongoDb.Collection(Credit_CardCollectionName()).FindOne(context.TODO(), bson.M{"_id": objectId}).Decode(&model)
// 	if err != nil {
// 		return model, co.ErrorPass(err.Error())
// 	}

// 	// model.ViewCount = model.ViewCount + 1
// 	model.Update()

// 	return model, co.SuccessPass("")
// }

// // .

// /* ***********************************************************************
// * search
//  * *********************************************************************** */

// // .
// type SearchCredit_Card struct {
// 	//
// 	comn.Search

// 	//

// 	Credit_Cards []*Credit_Card
// }

// func (search *SearchCredit_Card) CollectionName() string {
// 	//
// 	return Credit_CardCollectionName()
// }

// // .
// func (search *SearchCredit_Card) condition() bson.M {

// 	filter := bson.M{}

// 	filter["able"] = true

// 	return filter

// }

// // .
// func (search *SearchCredit_Card) Finds() (errEx co.MsgEx) {
// 	sort := bson.M{}
// 	if co.NotEmptyString(search.SortField) {
// 		if search.SortDirection != 1 {
// 			search.SortDirection = -1
// 		} else {
// 			search.SortDirection = 1
// 		}
// 		sort[search.SortField] = search.SortDirection
// 	} else {
// 		sort["createdtime"] = -1
// 	}

// 	if search.Limit > 0 && search.PageOffset > -1 {

// 		cursor, err := inits.MongoDb.Collection(search.CollectionName()).Find(context.TODO(), search.condition(),
// 			options.Find().SetSkip(int64(search.Limit)*int64(search.PageOffset)).SetLimit(int64(search.Limit)).SetSort(sort))
// 		if err != nil {
// 			return co.ErrorPass(err.Error())
// 		}

// 		if err = cursor.All(context.TODO(), &search.Credit_Cards); err != nil {
// 			return co.ErrorPass(err.Error())
// 		}

// 	} else {

// 		cursor, err := inits.MongoDb.Collection(search.CollectionName()).Find(context.TODO(), search.condition(), options.Find().SetSort(sort))
// 		if err != nil {
// 			return co.ErrorPass(err.Error())
// 		}

// 		if err = cursor.All(context.TODO(), &search.Credit_Cards); err != nil {
// 			return co.ErrorPass(err.Error())
// 		}
// 	}

// 	total, err := inits.MongoDb.Collection(search.CollectionName()).CountDocuments(context.TODO(), search.condition())
// 	if err != nil {
// 		return co.ErrorPass(err.Error())
// 	}

// 	search.Total = total

// 	return errEx

// }

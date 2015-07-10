package mongo

import (
	"github.com/kahoona77/emerald/app/models"
	"github.com/revel/revel"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

var mongoSession *mgo.Session

func InitDB() {

	//creating db
	var err error
	mongoSession, err = mgo.Dial("localhost") //mgo.Dial("192.168.56.101") //mgo.Dial("localhost")
	if err != nil {
		revel.INFO.Println("DB Error", err)
	} else {
		revel.INFO.Println("DB Connected")
	}
}

func getCollection(collectionName string) *mgo.Collection {
	return mongoSession.DB("xtv").C(collectionName)
}

func All(collection string, results interface{}) error {
	return getCollection(collection).Find(nil).All(results)
}

func CountAll(collection string) (int, error) {
	return getCollection(collection).Find(nil).Count()
}

func FindWithQuery(collection string, query *bson.M, results interface{}) error {
	return getCollection(collection).Find(query).All(results)
}

func FindById(collection string, docId string, result models.MongoModel) error {
	return getCollection(collection).FindId(docId).One(result)
}

func FindFirst(collection string, result models.MongoModel) error {
	return getCollection(collection).Find(nil).One(result)
}

func Remove(collection string, docId string) error {
	return getCollection(collection).RemoveId(docId)
}

func RemoveAll(collection string, query *bson.M) (info *mgo.ChangeInfo, err error) {
	return getCollection(collection).RemoveAll(query)
}

func Save(collection string, docId string, doc models.MongoModel) (info *mgo.ChangeInfo, err error) {
	if docId == "" {
		docId = bson.NewObjectId().Hex()
		doc.SetId(docId)
	}
	return getCollection(collection).UpsertId(docId, doc)
}

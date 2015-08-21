package mongo

import (
	"log"

	"github.com/kahoona77/emerald/models"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

//MongoService connects with MongoDB
type MongoService struct {
	session *mgo.Session
}

//NewService creates a new MongoService
func NewService(conf *models.AppConfig) *MongoService {
	m := new(MongoService)

	//creating db
	session, err := mgo.Dial(conf.Mongodb) //mgo.Dial("192.168.56.101") //mgo.Dial("localhost")
	if err != nil {
		log.Print("DB Error", err)
	} else {
		m.session = session
		log.Print("DB Connected")
	}
	return m
}

func (m *MongoService) getCollection(collectionName string) *mgo.Collection {
	return m.session.DB("xtv").C(collectionName)
}

func (m *MongoService) All(collection string, results interface{}) error {
	return m.getCollection(collection).Find(nil).All(results)
}

func (m *MongoService) CountAll(collection string) (int, error) {
	return m.getCollection(collection).Find(nil).Count()
}

func (m *MongoService) FindWithQuery(collection string, query *bson.M, results interface{}) error {
	return m.getCollection(collection).Find(query).All(results)
}

func (m *MongoService) FindById(collection string, docId string, result models.MongoModel) error {
	return m.getCollection(collection).FindId(docId).One(result)
}

func (m *MongoService) FindFirst(collection string, result models.MongoModel) error {
	return m.getCollection(collection).Find(nil).One(result)
}

func (m *MongoService) Remove(collection string, docId string) error {
	return m.getCollection(collection).RemoveId(docId)
}

func (m *MongoService) RemoveAll(collection string, query *bson.M) (info *mgo.ChangeInfo, err error) {
	return m.getCollection(collection).RemoveAll(query)
}

func (m *MongoService) Save(collection string, docId string, doc models.MongoModel) (info *mgo.ChangeInfo, err error) {
	if docId == "" {
		docId = bson.NewObjectId().Hex()
		doc.SetId(docId)
	}
	return m.getCollection(collection).UpsertId(docId, doc)
}

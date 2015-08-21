package db

import "github.com/HouzuoGuo/tiedot/db"

type DataBaseService struct {
	db *db.DB
}

func NewService() *DataBaseService {
	service := new(DataBaseService)

	// (Create if not exist) open a database
	emeraldDB, err := db.OpenDB("db")
	if err != nil {
		panic(err)
	}

	if err := emeraldDB.Create("servers"); err != nil {
		panic(err)
	}
	if err := emeraldDB.Create("settings"); err != nil {
		panic(err)
	}

	service.db = emeraldDB

	return service
}

func (dbs *DataBaseService) getCollection(collectionName string) *db.Col {
	return dbs.db.Use(collectionName)
}

func (dbs *DataBaseService) readQueryResults(col *db.Col, queryResult map[int]struct{}, results interface{}) error {
	// Query result are document IDs
	for id := range queryResult {
		// To get query result document, simply read it
		readBack, err := col.Read(id)
		if err != nil {
			panic(err)
		}
	}
	return nil
}

func (dbs *DataBaseService) All(collection string) error {
	col := dbs.getCollection(collection)
	queryResult := make(map[int]struct{})
	if err := db.EvalAllIDs(col, &queryResult); err != nil {
		return err
	}

	return nil
}

// func (dbs *DataBaseService) CountAll(collection string) (int, error) {
// 	return dbs.getCollection(collection).Find(nil).Count()
// }
//
// func (dbs *DataBaseService) FindWithQuery(collection string, query *bson.M, results interface{}) error {
// 	return dbs.getCollection(collection).Find(query).All(results)
// }
//
// func (dbs *DataBaseService) FindById(collection string, docId string, result models.MongoModel) error {
// 	return dbs.getCollection(collection).FindId(docId).One(result)
// }
//
// func (dbs *DataBaseService) FindFirst(collection string, result models.MongoModel) error {
// 	return dbs.getCollection(collection).Find(nil).One(result)
// }
//
// func (dbs *DataBaseService) Remove(collection string, docId string) error {
// 	return dbs.getCollection(collection).RemoveId(docId)
// }
//
// func (dbs *DataBaseService) RemoveAll(collection string, query *bson.M) (info *mgo.ChangeInfo, err error) {
// 	return dbs.getCollection(collection).RemoveAll(query)
// }
//
// func (dbs *DataBaseService) Save(collection string, docId string, doc models.MongoModel) (info *mgo.ChangeInfo, err error) {
// 	if docId == "" {
// 		docId = bson.NewObjectId().Hex()
// 		doc.SetId(docId)
// 	}
// 	return dbs.getCollection(collection).UpsertId(docId, doc)
// }

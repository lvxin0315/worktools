package driver

import (
	"gopkg.in/mgo.v2"
	"fmt"
)

func GetMongoDBSession(url string) * mgo.Session {
	session, err := mgo.Dial(url) //连接数据库
	if err != nil {
		panic(err)
		return nil
	}
	session.SetMode(mgo.Monotonic, true)
	return session
}

func ChangeDB(session * mgo.Session, dbName string) * mgo.Database{
	db := session.DB(dbName) //数据库名称
	return db
}

func UserCollection(db * mgo.Database,collectionName string) * mgo.Collection {
	collection := db.C(collectionName) //如果该集合已经存在的话，则直接返回
	return collection
}

func Connect(url string,dbName string,collectionName string) (*mgo.Session,* mgo.Collection){
	session := GetMongoDBSession(url)
	if nil == session {
		fmt.Println("session is error")
	}
	db := ChangeDB(session,dbName)
	collection := UserCollection(db,collectionName)
	if nil == collection{
		fmt.Println("collection is error")
	}
	return session,collection
}
package db

import (
	mgo "gopkg.in/mgo.v2"
)

func InitMongoDB() *mgo.Session {
	session, err := mgo.Dial("localhost")
    if err!= nil {
        panic(err)
    }
    return session
}
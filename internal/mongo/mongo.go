package mongo

import (
	"log"

	mgo "gopkg.in/mgo.v2"
)

var session *mgo.Session

func ConnectOrDie(addr string) {
	var err error

	session, err = mgo.Dial(addr)
	if err != nil {
		log.Fatal(err)
	}
}

func Session() *mgo.Session {
	return session
}

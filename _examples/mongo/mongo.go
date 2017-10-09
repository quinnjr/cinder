package main

import (
	"fmt"
	"os"
	"time"

	"github.com/quinnjr/cinder"
	"github.com/quinnjr/cinder/handlers/mongo"
	mgo "gopkg.in/mgo.v2"
)

func main() {

	info, err := mgo.ParseURL("mongodb://127.0.0.1:27017")
	if err != nil {
		panic(err)
	}

	session, err := mgo.DialWithInfo(info)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	logCol := session.DB("cinder").C("logs")

	logger := cinder.New(cinder.DebugLevel, mongo.New(logCol))

	time := time.Now()

	logger.WithField("time", time).Info("mongodb connection info")

	var res []interface{}

	err = logCol.Find(nil).All(&res)
	if err != nil {
		panic(err)
	}

	fmt.Print(res)

	os.Exit(0)
}

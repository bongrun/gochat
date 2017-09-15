// This sample is about using long polling and WebSocket to build a web-based chat room based on beego.
package main

import (
	"github.com/astaxie/beego"
	_ "gochat/back/routers"
	"gochat/back/conf"
)

const (
	APP_VER = "0.1.1.0227"
)

type Person struct {
	Name string
	Phone string
}

func main() {
	conf.Init()
	//
	//c := conf.DI.DB.C("people")
	//err := c.Insert(&Person{"Ale", "+55 53 8116 9639"},
	//	&Person{"Cla", "+55 53 8402 8510"})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//result := Person{}
	//err = c.Find(bson.M{"name": "Ale"}).One(&result)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println("Phone:", result.Phone)

	beego.Info(beego.BConfig.AppName, APP_VER)
	beego.Run()
	//fmt.Println("|" + beego.AppConfig.String("mongo_host") + "|")
	//session, err := mgo.Dial(beego.AppConfig.String("mongo_host"))
	//if err != nil {
	//	panic(err)
	//}
	//defer session.Close()
	//
	//// Optional. Switch the session to a monotonic behavior.
	//session.SetMode(mgo.Monotonic, true)
	//
	//c := session.DB("test").C("people")
	//err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
	//	&Person{"Cla", "+55 53 8402 8510"})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//result := Person{}
	//err = c.Find(bson.M{"name": "Ale"}).One(&result)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println("Phone:", result.Phone)
}
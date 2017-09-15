package conf

import (
	"gopkg.in/mgo.v2"
)

type di struct {
	DB *mgo.Database
}

var DI = new(di)

func Init()  {
	DI.DB = func() *mgo.Database {
		session, err := mgo.Dial("192.168.99.100")
		if err != nil {
			panic(err)
		}
		//defer session.Close()

		// Optional. Switch the session to a monotonic behavior.
		session.SetMode(mgo.Monotonic, true)

		return session.DB("gochat")
	}()
}

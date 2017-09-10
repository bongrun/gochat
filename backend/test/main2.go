package test

import (
	"net/http";
	"io"
)
import (
	"encoding/json"
	"strconv"
	"log"
	"fmt"
)

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func hello(res http.ResponseWriter, req *http.Request) {
	id, _ := strconv.ParseInt(req.FormValue("id"), 10, 64)
	mapD := map[string]int64{"apple": 5, "lettuce": 7, "id": id}
	mapB, _ := json.Marshal(mapD)
	res.Header().Set(
		"Content-Type",
		"application/json",
	)
	io.WriteString(
		res,
		string(mapB),
	)
}

func main() {
	mongo()
	//http.HandleFunc("/hello", hello)
	//http.ListenAndServe(":9002", nil)
}

type Person struct {
	Name  string
	Phone string
}

func mongo() {
	session, err := mgo.Dial("gochat_mongo")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer func() {
		session.Close()
		fmt.Println("Close")
	}()


	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("people")
	err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
		&Person{"Cla", "+55 53 8402 8510"})
	if err != nil {
		log.Fatal(err)
	}

	result := Person{}
	err = c.Find(bson.M{"name": "Ale"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Phone:", result.Phone)
}

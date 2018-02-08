package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Baby struct {
	Id        bson.ObjectId `bson:"_id"`
	Name      string        `bson:"name"`
	Age       int           `bson:"age"`
	Address   string        `bson:"address"`
	CreatedAt time.Time     `bson:"created_at"`
	Hobbies   []string      `bson:"hobbies"`
}

func main() {
	session, err := mgo.Dial("192.168.209.128:27017")
	if err != nil {
		fmt.Println(err)
	}
	defer session.Close()

	db := session.DB("test").C("baby")

	db.Insert(
		&Baby{bson.NewObjectId(), "aaa", 11, "上海", time.Now(), []string{"玩", "吃"}},
		&Baby{bson.NewObjectId(), "bbb", 12, "北京", time.Now(), []string{"玩", "吃"}},
		&Baby{bson.NewObjectId(), "ccc", 13, "南京", time.Now(), []string{"玩", "吃"}},
	)

	var result []Baby
	db.Find(nil).All(&result)
	for _, row := range result {
		fmt.Println(row.Id, row.Name, row.Age, row.Address, row.CreatedAt, row.Hobbies)
	}

	var row Baby
	db.Find(bson.M{"name": "aaa"}).One(&row)
	fmt.Println(row.Id, row.Name, row.Age, row.Address, row.CreatedAt, row.Hobbies)
}

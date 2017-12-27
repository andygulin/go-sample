package main

import (
	"encoding/json"
	"time"
	"fmt"
)

//字段名必须首字母大写才能转，醉了
type User struct {
	Id        int      `json:"id"`
	Name      string   `json:"name"`
	Age       int      `json:"age"`
	Address   string   `json:"address"`
	CreatedAt JSONTime `json:"created_at"`
}

//重写序列化，格式化时间字段
type Marshaler interface {
	MarshalJSON() ([]byte, error)
}

type JSONTime time.Time

func (t JSONTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}

func main() {
	jsonUser()
	jsonUserArray()
	jsonMap()
}

func jsonUser() {
	user := &User{
		Id:        1,
		Name:      "aaa",
		Age:       11,
		Address:   "上海",
		CreatedAt: JSONTime(time.Now()),
	}
	b, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
}

func jsonUserArray() {
	users := []User{
		{Id: 1, Name: "aaa", Age: 11, Address: "上海", CreatedAt: JSONTime(time.Now())},
		{Id: 2, Name: "bbb", Age: 11, Address: "北京", CreatedAt: JSONTime(time.Now())},
		{Id: 3, Name: "ccc", Age: 11, Address: "广州", CreatedAt: JSONTime(time.Now())},
		{Id: 4, Name: "ddd", Age: 11, Address: "杭州", CreatedAt: JSONTime(time.Now())},
		{Id: 5, Name: "eee", Age: 11, Address: "南京", CreatedAt: JSONTime(time.Now())},
	}
	b, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
}

func jsonMap() {
	row := make(map[string]interface{})
	row["id"] = 1
	row["name"] = "aaa"
	row["age"] = 11
	row["address"] = "上海"
	row["created_at"] = JSONTime(time.Now())

	b, err := json.Marshal(row)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
}

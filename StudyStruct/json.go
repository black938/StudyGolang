package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	ID     int `json:"idididid"`
	Gender string
	Name   string
}

type Class struct {
	Title    string
	Students []*Student
}

func tryJSON() {
	c := Class{
		Title:    "101",
		Students: make([]*Student, 0, 200),
	}
	for i := 0; i < 10; i++ {
		stu := &Student{
			ID:     i,
			Gender: "男",
			Name:   fmt.Sprintf("171144%02d", i),
		}
		c.Students = append(c.Students, stu)
	}
	// JSON序列化 struct -> json string
	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println("json marshal failed")
		return
	}
	fmt.Printf("json string : %s \n", data)
	// JSON反序列化 json string -> struct
	sss := `{"Id":1,"Name":"wwy"}`
	c1 := &struct {
		Id   int
		Name string
	}{}
	err = json.Unmarshal([]byte((sss)), c1)
	if err != nil {
		panic(err)
		fmt.Println("json unmarshal failed")
		return
	}
	fmt.Printf("%#v \n", c1)

	// 如果不知道JSON的格式 可以用一个空接口来接收
	// 其它golang中json的tricks https://www.liwenzhou.com/posts/Go/json_tricks_in_go/
}

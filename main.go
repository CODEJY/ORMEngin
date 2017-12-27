package main

import (
	"fmt"
	"time"

	"github.com/CODEJY/ORMEngine/dao"

	"github.com/CODEJY/ORMEngine/entity"
)

var user1 entity.UserInfo
var user2 entity.UserInfo
var engine *dao.ORMEngine

func init() {
	engine = dao.NewEngine("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=true")
	t := time.Now()
	user1 = entity.UserInfo{
		UserName:   "wujy",
		DepartName: "Software Design",
		CreateAt:   &t,
	}
	t = time.Now()
	user2 = entity.UserInfo{
		UserName:   "gary",
		DepartName: "Software Development",
		CreateAt:   &t,
	}
}
func main() {
	result, err := engine.Insert(user1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d row(s) result after inserting %s\n", result, user1.UserName)
	result, err = engine.Insert(user2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d row(s) result after inserting %s\n", result, user2.UserName)

	allUsers := make([]*entity.UserInfo, 0)
	err = engine.Find(&allUsers)
	if err != nil {
		panic(err)
	}
	fmt.Println("all users: ")
	for i := 0; i < len(allUsers); i++ {
		fmt.Println(*allUsers[i])
	}
}

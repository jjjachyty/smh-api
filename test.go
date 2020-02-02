package main

import (
	"encoding/json"
	"fmt"
	"smh-api/models"
)

func main() {
	// db.InitDB()
	// mv := models.Movie{Name: "1987：黎明到来的那一天",
	// 	Actor:    "金允石 河正宇 金泰梨",
	// 	Director: "张俊焕",
	// 	Genre:    "剧情片",
	// 	Region:   "韩国",
	// 	Years:    "2018",
	// 	ScoreDB:  90,
	// 	Cover:    "https://www.imgdouban.com/diaosidao/uploads/allimg/180209/1bcd0ad678c4b6b4.jpg",
	// }
	// mv.ID = hex.EncodeToString([]byte(mv.Name))
	// mv.Insert()
	// res := models.Resources{MovieID: "e99d9ee5b8b8e5908ce4bc99", Name: "HD1280", URL: "https://videos.kkyun-iqiyi.com/ppvod/D26C9E72A904B7090C4B63680122B669.m3u8"}
	// res.Insert()

	user := models.User{IP: "12121"}
	user1 := models.User{}
	byts, err := json.Marshal(user)
	fmt.Println(err)
	fmt.Println(json.Unmarshal(byts, &user1))
}

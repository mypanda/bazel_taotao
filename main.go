package main

import (
	"fmt"
	"taotao/handler"
	jsoniter "github.com/json-iterator/go"
)


type AppInfo struct{
	Name string
	Age int
}

func main() {
	info := AppInfo{
		Name:"TAOTAO",
		Age:10,
	}
	jsonString,_ := jsoniter.Marshal(&info)
	fmt.Println(string(jsonString))
	handler.ShowInfo()
}
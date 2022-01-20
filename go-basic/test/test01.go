package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// 记忆考验 2022年01月16日 13:58:21
	// 1.基本数据类型及转换
	// 2.init函数 main函数
	// 3.array
	// 4.slice
	// 5.map
	// 6.ptr
	// 7.自定义类型和类型别名
	// 8.struct

	var initUint8 uint8 = 255
	fmt.Println(initUint8)

	var initArr []int = []int{1, 1, 1}
	initArr2 := []string{"你好", "世界"}
	fmt.Println(initArr, initArr2)

	slice := initArr2[:]
	fmt.Println("当前切片:", slice)

	var initMap = make(map[string]string, 2)
	fmt.Println("初始化的map信息:", initMap, len(initMap))
	initMap["username"] = "董健"
	initMap["nickname"] = "猜"
	initMap["shortname"] = "猜2"
	fmt.Println("现在的map信息:", initMap, len(initMap))

	var str string = "hello world"
	var p = &str
	fmt.Printf("当前申明的p的类型是:%T\n", p)
	*p = "hello2 world2"
	fmt.Println(str)

	type newInt int
	var initInt newInt = 123
	fmt.Printf("initInt的值和类型是:%v,%T \n", initInt, initInt)

	type User struct {
		ID    int32 `json:"c_id"`
		Name  string
		pwd   string
		Token string
	}

	var user1 User
	// user1.ID = 1
	user1.Name = "董健"
	// user1.pwd = "QWEGSFSFSAF"
	// user1.Token = "Token-xxxxxxxx"
	user2 := User{
		ID:    2,
		Name:  "董健2",
		pwd:   "fjslajfl",
		Token: "Token-xxxxx",
	}

	user1Json, err := json.Marshal(user1)
	if err != nil {
		fmt.Println("struct转json失败!")
	} else {
		fmt.Println("json格式的字符串是:", string(user1Json))
	}
	user2Json, err := json.Marshal(user2)
	if err != nil {
		fmt.Println("struct转json失败!")
	} else {
		fmt.Println("json格式的字符串是:", string(user2Json))
	}

}

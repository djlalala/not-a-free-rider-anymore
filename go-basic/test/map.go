// author 董健 <joeyana@aliyun.com>
package main

import "fmt"

/***
	map是一种无序的基于key-value的数据结构，Go语言中的map是引用类型，必须初始化才能使用。
	map的定义语法:map[KeyType]ValueType
	    | KeyType:表示键的类型。
		| ValueType:表示键对应的值的类型。
	map类型的变量默认初始值为nil，需要使用make()函数来分配内存。
		| 语法: make(map[KeyType]ValueType, [cap])
***/
func main() {

	// 基本使用1
	fmt.Println("------------------------------------------------------------")
	fmt.Println("map初始化2:")
	scoreMap := make(map[string]int, 8)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	scoreMap["小贱贱"] = 101
	fmt.Println(scoreMap)
	fmt.Println(scoreMap["小明"])
	fmt.Printf("scoreMap的类型是: %T\n", scoreMap)

	// 基本使用2
	// map也支持在声明的时候填充元素
	fmt.Println("------------------------------------------------------------")
	fmt.Println("map初始化2:")
	userinfo := map[string]string{
		"username": "d.joeyana",
		"password": "pwd",
	}
	fmt.Println("userinfo:", userinfo)

	// 判断某个键是否存在
	// value, ok := map[key]
	fmt.Println("------------------------------------------------------------")
	fmt.Println("判断map的某个键是否存在:")
	v, ok := scoreMap["张三"]
	if ok {
		fmt.Println("他的socre:", v)
	} else {
		fmt.Println("查无此人")
	}

	// map的遍历
	fmt.Println("------------------------------------------------------------")
	fmt.Println("map的遍历:")
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}

	// 使用delete()函数删除键值对
	fmt.Println("------------------------------------------------------------")
	fmt.Println("map删除键值对:")
	fmt.Println("删除之前的map:", scoreMap)
	delete(scoreMap, "小贱贱")
	fmt.Println("删除之后的map:", scoreMap)

	// 元素为map类型的切片
	fmt.Println("------------------------------------------------------------")
	fmt.Println("元素为map类型的切片:")
	var mapSlice = make([]map[string]string, 3)
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
	fmt.Println("after init")
	// 对切片中的map元素进行初始化
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "王五"
	mapSlice[0]["password"] = "123456"
	mapSlice[1] = make(map[string]string, 10)
	mapSlice[1]["address"] = "红旗大街"
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}

	// 值为切片类型的map
	fmt.Println("------------------------------------------------------------")
	fmt.Println("值为切片类型的map:")
	var sliceMap = make(map[string][]string, 3)
	fmt.Println(sliceMap)
	fmt.Println("after init")
	key := "中国"
	value, ok := sliceMap[key]
	if !ok {
		value = make([]string, 0, 2)
	}
	value = append(value, "北京", "上海")
	sliceMap[key] = value
	fmt.Println(sliceMap)
}

// author 董健 <joeyana@aliyun.com>
package main

import (
	"encoding/json"
	"fmt"
	"sort"
)

//Student 学生
type Student struct {
	ID     int    `json:"json_id"` // 大写字母开头 通过指定tag实现json序列化该字段时的key
	Gender string // 大写字母开头 json序列化是默认使用字段名作为key
	name   string // 小写字母开头 私有不能被json包访问
}

/***
	1.结构体标签（Tag）
	2.map的有序输出
***/
func main() {
	// Tag是结构体的元信息，可以在运行的时候通过反射的机制读取出来。
	// Tag在结构体字段的后方定义，由一对反引号包裹起来，具体的格式如下：`key1:"value1" key2:"value2"`
	fmt.Println("------------------------------------------------------------")
	fmt.Println("用Tag定义结构体做json序列化时使用的结构体:")
	s1 := Student{
		ID:     1,
		Gender: "女",
		name:   "pprof",
	}
	data, err := json.Marshal(s1)
	if err != nil {
		fmt.Println("json marshal failed!")
		return
	}
	fmt.Printf("jsonMarshal的值类型和值:%T,%s\n", data, data) //[]uint8,{"json_id":1,"Gender":"女"}

	// 实现map有序输出
	fmt.Println("------------------------------------------------------------")
	fmt.Println("实现map有序输出:")
	map1 := make(map[int]string, 5)
	map1[1] = "张三"
	map1[2] = "小贱贱"
	map1[5] = "小月月"
	map1[9] = "阿爸阿爸"
	map1[4] = "想静静"
	kSlice := []int{}
	// step1: 先循环一遍map，把map1中的key单独放到kSlice中
	for k, _ := range map1 {
		kSlice = append(kSlice, k)
	}
	// step2: Ints sorts a slice of ints in increasing order.
	fmt.Println("old kSlice:", kSlice)
	sort.Ints(kSlice)
	fmt.Println("new kSlice:", kSlice)
	// step3:根据kSlice的key排好的顺序，输出map的对应value
	for i := 0; i < len(map1); i++ {
		fmt.Println(i, " : ", map1[kSlice[i]])
	}
}

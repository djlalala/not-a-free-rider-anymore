// author 董健 <joeyana@aliyun.com>
package main

import "fmt"

/***
	1.类型别名和自定义类型 的区别
	2.结构体
***/
func main() {
	// 1.类型别名和自定义类型
	// a.自定义类型 type NewInt int
	// b.类型别名 type MyInt = int
	// ab 区别:
	// 结果显示a的类型是main.NewInt，表示main包下定义的NewInt类型。
	// b的类型是int。MyInt类型只会在代码中存在，编译完成时并不会有MyInt类型。
	fmt.Println("------------------------------------------------------------")
	fmt.Println("类型别名和自定义类型的区别:")
	type NewInt int
	type MyInt = int
	type byte = uint8
	type rune = int32
	var a NewInt
	var b MyInt
	var str string = "你们好"
	bytes := []byte(str)
	runes := []rune(str)
	fmt.Printf("当前a的类型是: %T\n", a)
	fmt.Printf("当前b的类型是: %T\n", b)
	fmt.Printf("当前bytes的类型是: %T\n", bytes)
	fmt.Printf("当前runes的类型是: %T\n", runes)

	// 使用type和struct关键字来定义结构体
	fmt.Println("------------------------------------------------------------")
	fmt.Println("结构体的定义和实例化:")
	type person struct {
		name string
		city string
		age  int8
	}
	var p1 person
	p1.name = "pprof.cn"
	p1.city = "北京"
	p1.age = 18
	fmt.Printf("p1=%v\n", p1)  //p1={pprof.cn 北京 18}
	fmt.Printf("p1=%#v\n", p1) //p1=main.person{name:"pprof.cn", city:"北京", age:18}

}

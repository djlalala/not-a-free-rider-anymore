package main

import (
	"fmt"
	tools "go-basic/tools"
	"os"
	// tools "go-basic/tools"
)

func init() {
	tools.Hello()
	fmt.Println("------------------------------------------------------------")
	dir, _ := os.Getwd()
	fmt.Println("当前路径:", dir)
	fmt.Printf("当前位置: %v() \n", tools.GetFunctionName(main))
	fmt.Println("Slice底层实现:")
}

/***
	Slice底层实现
		| 切片和数组
		| 切片的数据结构
		| 创建切片
		| 切片扩容
***/
func main() {

}

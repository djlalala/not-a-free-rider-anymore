// author 董健 <joeyana@aliyun.com>
package main

import (
	"fmt"
	tools "go-basic/tools"
)

/***
	1.Init函数和main函数
	2.先执行init方法，再去执行main
	3.如果当前包中有import的包，则优先执行被导入的包中的init方法
***/
func init() {
	fmt.Println("init 函数")
}

func main() {
	tools.Hello()
	fmt.Println("main 函数")
}

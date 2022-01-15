// author 董健 <joeyana@aliyun.com>
package main

import "fmt"

/***
	指针小练习
***/
func main() {
	fmt.Println("------------------------------------------------------------")
	fmt.Println("指针小练习:")
	var a int
	fmt.Println(a, &a)
	var p *int
	p = &a
	*p = 20
	fmt.Println(*p, p)
	fmt.Println(*p, &p)
	fmt.Println(a, &a)
}

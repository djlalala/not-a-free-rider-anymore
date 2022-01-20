package main

import "fmt"

func printValue(args ...string) { // 可以接受任意个数string
	for _, v := range args {
		fmt.Println(v)
	}
}

/****
	0. 三个点 '...' 的用法
	1. if else
	2. switch case
	3. select case	(每个case必须是一个通信操作,要么是发送要么是接收)
	4. for
	5. for range (复制????)
	6. goto / break / continue
****/
func main() {
	// ...的用法
	var s = []string{"hello", "wolrd", "do", "not"}
	printValue(s...) // 切片被打散传入

	// if, else if, else,
	if x := 0; x > 0 {
		fmt.Println("x>0")
	} else if x == 0 {
		fmt.Println("x==0")
		fmt.Printf("%T\n", x)
	} else {
		fmt.Println("x>0")
	}

	// switch case
	var grade string = "B"
	switch {
	case grade == "A":
		fmt.Printf("优秀!\n")
	case grade == "B", grade == "C":
		fmt.Printf("良好\n")
	case grade == "D":
		fmt.Printf("及格\n")
	case grade == "F":
		fmt.Printf("不及格\n")
	default:
		fmt.Printf("差\n")
	}
	fmt.Printf("你的等级是 %s\n", grade)

	ss := "abcd"
	for i, n := 0, length(ss); i < n; i++ { // 避免多次调用 length 函数。
		println(i, s[i])
	}
}

func length(s string) int {
	println("call length.")
	return len(s)
}

package main

import (
	"fmt"
	"time"
)

/****
	主协程退出了，其他任务还执行吗?
	---程序启动时，Go程序就会为main()函数创建一个默认的goroutine。
	---当main()函数返回的时候该goroutine就结束了，所有在main()函数中启动的goroutine会一同结束，
****/
func main() {
	// 合起来写
	go func() {
		i := 0
		for {
			i++
			fmt.Printf("new goroutine: i = %d\n", i)
			time.Sleep(time.Second)
		}
	}()
	i := 0
	for {
		i++
		fmt.Printf("main goroutine: i = %d\n", i)
		time.Sleep(time.Second)
		if i == 2 {
			break
		}
	}
}

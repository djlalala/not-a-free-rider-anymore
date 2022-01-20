package main

import (
	"fmt"
	"runtime"
	"time"
)

/***
	2.runtime.Goexit()
	---退出当前协程,且在当前协程中goexit后面的代码行都不会执行
***/
func main() {
	go func() {
		defer fmt.Println("A.defer")
		func() {
			go func() {
				for {
					fmt.Println("goroutine的第 1 匿名函数,定时打印...")
					time.Sleep(time.Second)
				}
			}()
			defer fmt.Println("B.defer")
			// 结束当前协程，goexit后面的都不会执行
			fmt.Println("准备结束当前协程")
			runtime.Goexit()
			defer fmt.Println("C.defer")
			fmt.Println("B")
		}()

		func() {
			go func() {
				for {
					fmt.Println("goroutine的第 2 匿名函数,定时打印...")
					time.Sleep(time.Second)
				}
			}()
		}()
		fmt.Println("A")
	}()
	for { // 添加死循环，确保主协程一直不结束
		fmt.Println("主协程定时打印...")
		time.Sleep(time.Second)
	}
}

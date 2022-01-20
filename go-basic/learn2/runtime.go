package main

import (
	"fmt"
	"runtime"
)

/***
	1.runtime.Gosched()
	---让出CPU时间片，重新等待安排任务
***/
func main() {
	go func(s string) {
		for i := 0; i < 2; i++ {
			fmt.Println(s)
		}
	}("world")
	// 主协程
	for i := 0; i < 2; i++ {
		// 切一下，再次分配任务，不切的话，主协程就跑完了
		runtime.Gosched()
		fmt.Println("hello")
	}
}

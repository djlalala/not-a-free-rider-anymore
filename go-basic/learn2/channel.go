package main

import (
	"fmt"
)

func main() {
	defer func() {
		fmt.Println("程序结束.")
	}()
	// 使用ch := make(chan int)创建的是无缓冲的通道，
	// 无缓冲的通道只有在有人接收值的时候才能发送值。就像你住的小区没有快递柜和代收点，
	// 快递员给你打电话必须要把这个物品送到你的手中，简单来说就是无缓冲的通道必须有接收才能发送。
	// 上面的代码会阻塞在ch <- 10这一行代码形成死锁，那如何解决这个问题呢？
	// ch := make(chan int) // 无缓冲的通道
	// ch <- 10
	// fmt.Println("发送成功")

	ch := make(chan int, 1) // 创建一个容量为1的有缓冲区通道
	ch <- 10
	fmt.Println("发送成功", ch)

	/***
		可以通过内置的close()函数关闭channel（如果你的管道不往里存值或者取值的时候一定记得关闭管道）
	***/
	fmt.Println("------------------------------------------------------------")
	fmt.Println("goroutinew往channel里面循环发送数据和取数据:")
	c := make(chan int) // 无缓冲区通道
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
			fmt.Println("channel接受数据:", i, c)
		}
		close(c)
	}()
	for {
		if data, ok := <-c; ok {
			fmt.Println("channel发送数据:", data, c)
		} else {
			break
		}
	}
}

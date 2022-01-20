package main

import "fmt"

/***
	没有结构化异常，使用 panic 抛出错误，recover 捕获错误。
	异常的使用场景简单描述：Go中可以抛出一个panic的异常，然后在defer中通过recover捕获这个异常，然后正常处理。
	panic:
		1、内置函数
		2、假如函数F中书写了panic语句，会终止其后要执行的代码，在panic所在函数F内如果存在要执行的defer函数列表，按照defer的逆序执行
		3、返回函数F的调用者G，在G中，调用函数F语句之后的代码不会执行，假如函数G中存在要执行的defer函数列表，按照defer的逆序执行
		4、直到goroutine整个退出，并报告错误
	recover:
		1、内置函数
		2、用来控制一个goroutine的panicking行为，捕获panic，从而影响应用的行为
		3、一般的调用建议
			a). 在defer函数中，通过recever来终止一个goroutine的panicking过程，从而恢复正常代码的执行
			b). 可以获取通过panic传递的error
***/

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("这里捕获的错误是:", err)
		}
	}()

	var ch chan int = make(chan int, 10)
	close(ch)
	ch <- 1

	// 发现:并没有继续执行
	fmt.Println("程序上方已经抛出错误，这段代码依然执行!")

	panic("test panic")
}
